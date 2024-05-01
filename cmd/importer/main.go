// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/examples"
	"github.com/crossplane/upjet/pkg/registry"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/upbound/provider-aws/config"
)

const (
	blockResource = "resource"
	tfStateFile   = "terraform.tfstate"
	tfExt         = ".tf"
)

func main() {
	var (
		app       = kingpin.New("importer", "Imports a set of resources from the specified Terraform workspace.").DefaultEnvars()
		workspace = app.Flag("workspace", "The path of the Terraform workspace.").Short('w').Default(".").ExistingDir()
		output    = app.Flag("output", "Output directory where the generated MR manifests will be dumped.").Short('o').String()
		region    = app.Flag("region", "Region of the generated MR manifests.").Short('r').Required().String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(false))
	logr := logging.NewLogrLogger(zl.WithName("importer-aws"))

	ctx := context.Background()
	kingpin.FatalIfError(filepath.WalkDir(*workspace, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			if path == *workspace {
				return nil
			}
			// only configurations in the root directory are considered
			return fs.SkipDir
		}
		if filepath.Ext(path) != tfExt {
			return nil
		}
		kingpin.FatalIfError(convertHCLConfiguration(ctx, logr, path, *output, *region), "Failed to convert the HCL configuration at path %s", path)
		return nil
	}), "Failed to discover the HCL configurations in the Terraform workspace %s", *workspace)
}

func convertHCLConfiguration(ctx context.Context, logr logging.Logger, hclPath, output, region string) error {
	workspace := filepath.Dir(hclPath)
	pc, err := config.GetProvider(ctx, true)
	if err != nil {
		return errors.Wrap(err, "failed to get the provider's configuration")
	}

	buff, err := os.ReadFile(hclPath)
	if err != nil {
		return errors.Wrapf(err, "failed to read the HCL configuration")
	}

	parser := hclparse.NewParser()
	f, diag := parser.ParseHCL(buff, hclPath)
	if diag != nil && diag.HasErrors() {
		return errors.Wrap(diag, "failed to parse the Terraform configuration")
	}

	body, ok := f.Body.(*hclsyntax.Body)
	if !ok {
		return errors.New("not an HCL Body")
	}
	resourceName := ""
	processedTypes := make(map[string]struct{})
	for _, b := range body.Blocks {
		if b.Type == blockResource {
			resourceName = b.Labels[0]
			if _, ok := processedTypes[resourceName]; ok {
				continue
			}
			processedTypes[resourceName] = struct{}{}
			if err := convertType(ctx, logr, pc, workspace, hclPath, resourceName, output, region, f, getResourceBlocks(body.Blocks, resourceName)); err != nil {
				return err
			}
		}
	}

	if resourceName == "" {
		// HCL configuration does not contain a resource block
		logr.Info("Skipping non-resource configuration file", "file", hclPath)
	}
	return nil
}

func getResourceBlocks(blocks hclsyntax.Blocks, resourceName string) hclsyntax.Blocks {
	trimmed := make(hclsyntax.Blocks, 0, len(blocks))
	for _, b := range blocks {
		if b.Type == blockResource && b.Labels[0] == resourceName {
			trimmed = append(trimmed, b)
		}
	}
	return trimmed
}

func convertType(ctx context.Context, logr logging.Logger, pc *ujconfig.Provider, workspace, hclPath, resourceName, output, region string, f *hcl.File, blocks hclsyntax.Blocks) error {
	cfg := pc.Resources[resourceName]
	if cfg == nil {
		logr.Info("Skipping resource because no configuration was found for it", "file", hclPath, "resource", resourceName)
		return nil
	}

	group := fmt.Sprintf("%s.%s", cfg.ShortGroup, pc.RootGroup)
	version := cfg.Version

	cfg.MetaResource.Examples = nil
	if err := cfg.MetaResource.FindExampleBlock(f, blocks, &resourceName, true); err != nil {
		return errors.Wrapf(err, "failed to find a resource configuration block for resource %q", resourceName)
	}
	cfg.MetaResource.Name = resourceName

	exArr := make([]registry.ResourceExample, 0, len(cfg.MetaResource.Examples))
	for _, re := range cfg.MetaResource.Examples {
		if err := re.Paved.UnmarshalJSON([]byte(re.Manifest)); err != nil {
			return errors.Wrap(err, "failed to unmarshal the JSON manifest")
		}

		schemaBlock := cfg.TerraformResource.CoreConfigSchema()
		rawConfig, err := schema.JSONMapToStateValue(re.Paved.UnstructuredContent(), schemaBlock)
		if err != nil {
			return errors.Wrap(err, "failed to convert the JSON map into a Terraform state value")
		}
		converted, err := schema.StateValueToJSONMap(rawConfig, schemaBlock.ImpliedType())
		if err != nil {
			return errors.Wrap(err, "failed to convert the Terraform state value to JSON map")
		}

		removeNullValues(converted)
		re.Paved = *fieldpath.Pave(converted)
		exArr = append(exArr, re)
	}

	tfState := filepath.Join(workspace, tfStateFile)
	extNames, err := ExtractExternalNames(ctx, tfState)
	if err != nil {
		return errors.Wrapf(err, "failed to extract external names from the Terraform state %s", tfState)
	}

	for _, re := range exArr {
		cfg.MetaResource.Examples[0] = re
		manifestPath := filepath.Join(output, filepath.Base(hclPath)+".yaml")
		gen := examples.NewGenerator(output, pc.ModulePath, pc.ShortName, pc.Resources,
			examples.WithAddExampleMetadata(false), examples.WithGenerateReferences(false), examples.WithAppendFile(true), examples.WithManifestPath(manifestPath))
		if err := gen.Generate(group, version, cfg); err != nil {
			return errors.Wrap(err, "failed to convert the HCL configuration to an MR manifest")
		}

		pm := gen.GetPavedWithManifest(fmt.Sprintf("%s.*", resourceName))
		if err := finalizeManifest(pm, extNames[fmt.Sprintf("%s.%s", resourceName, re.Name)], cfg.ShortGroup, region); err != nil {
			return err
		}

		if err := gen.StoreExamples(); err != nil {
			return errors.Wrapf(err, "failed to store the converted MR manifest to path %s", manifestPath)
		}
	}
	return nil
}

func finalizeManifest(pm *reference.PavedWithManifest, extName, shortGroup, region string) error {
	if err := pm.Paved.SetValue("metadata.annotations", map[string]string{
		"crossplane.io/external-name": extName,
	}); err != nil {
		return errors.Wrap(err, "failed to set the external-name annotation on the MR")
	}

	if err := pm.Paved.SetValue("spec.managementPolicies", []string{"Observe"}); err != nil {
		return errors.Wrap(err, `failed to set the spec.managementPolicies to "Observe" on the MR`)
	}

	if err := pm.Paved.SetValue("spec.deletionPolicy", "Orphan"); err != nil {
		return errors.Wrap(err, `failed to set the spec.deletionPolicy to "Orphan" on the MR`)
	}

	if config.HasRegion(shortGroup) {
		if err := pm.Paved.SetValue("spec.forProvider.region", region); err != nil {
			return errors.Wrapf(err, "failed to set the spec.forProvider.region to %q on the MR", region)
		}
	}
	return nil
}

func removeNullValues(m map[string]any) {
	for k, v := range m {
		switch val := v.(type) {
		case nil:
			// Remove nil values directly
			delete(m, k)
		case map[string]any:
			// Recursive call if the value is another map
			removeNullValues(val)
			if len(val) == 0 { // Check if nested map is now empty
				delete(m, k)
			}
		case []interface{}:
			// Handle slices of interfaces, which might contain maps or other slices
			updatedSlice := removeNilFromSlice(val)
			if len(updatedSlice) == 0 { // Check if slice is now empty
				delete(m, k)
			} else {
				m[k] = updatedSlice // Update the map with the modified slice
			}
		}
	}
}

func removeNilFromSlice(slice []interface{}) []interface{} {
	result := make([]interface{}, 0, len(slice))
	for _, elem := range slice {
		switch e := elem.(type) {
		case nil:
			// Skip nil elements
			continue
		case map[string]any:
			// Recursive removal for maps within the slice
			removeNullValues(e)
			if len(e) != 0 {
				result = append(result, e)
			}
		case []interface{}:
			// Recursively handle nested slices
			nestedSlice := removeNilFromSlice(e)
			if len(nestedSlice) != 0 {
				result = append(result, nestedSlice)
			}
		default:
			// Append non-nil, non-map elements directly
			result = append(result, elem)
		}
	}
	return result
}
