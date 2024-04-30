// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/upjet/pkg/examples"
	"github.com/crossplane/upjet/pkg/registry"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/upbound/provider-aws/config"
)

const (
	blockResource = "resource"
	tfStateFile   = "terraform.tfstate"
)

func main() {
	var (
		app       = kingpin.New("importer", "Imports a set of resources from the specified Terraform workspace.").DefaultEnvars()
		workspace = app.Flag("workspace", "The path of the Terraform workspace.").Short('w').Default(".").ExistingDir()
		output    = app.Flag("output", "Output directory where the generated MR manifests will be dumped.").Short('o').String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	hclPath := filepath.Join(*workspace, "vpc.tf")
	tfState := filepath.Join(*workspace, tfStateFile)
	ctx := context.Background()
	pc, err := config.GetProvider(ctx, true)
	if err != nil {
		panic(err)
	}

	buff, err := os.ReadFile(hclPath)
	if err != nil {
		panic(err)
	}

	parser := hclparse.NewParser()
	f, diag := parser.ParseHCL(buff, hclPath)
	if diag != nil && diag.HasErrors() {
		panic(errors.Wrapf(diag, "failed to parse the Terraform configuration %s. Configuration:\n%s", hclPath, string(buff)))
	}

	body, ok := f.Body.(*hclsyntax.Body)
	if !ok {
		panic(errors.Errorf("not an HCL Body: %s", string(buff)))
	}
	trimmed := make(hclsyntax.Blocks, 0, len(body.Blocks))
	resourceName := ""
	for _, b := range body.Blocks {
		if b.Type == blockResource {
			trimmed = append(trimmed, b)
			// TODO: we need to handle multiple resource types in a single HCL
			resourceName = b.Labels[0]
		}
	}

	group := fmt.Sprintf("%s.%s", pc.Resources[resourceName].ShortGroup, pc.RootGroup)
	version := pc.Resources[resourceName].Version

	pc.Resources[resourceName].MetaResource.Examples = nil
	body.Blocks = trimmed
	if err := pc.Resources[resourceName].MetaResource.FindExampleBlock(f, body.Blocks, &resourceName, true); err != nil {
		panic(err)
	}
	pc.Resources[resourceName].MetaResource.Name = resourceName

	exArr := make([]registry.ResourceExample, 0, len(pc.Resources[resourceName].MetaResource.Examples))

	for _, re := range pc.Resources[resourceName].MetaResource.Examples {
		if err := re.Paved.UnmarshalJSON([]byte(re.Manifest)); err != nil {
			panic(err)
		}

		schemaBlock := pc.Resources[resourceName].TerraformResource.CoreConfigSchema()
		rawConfig, err := schema.JSONMapToStateValue(re.Paved.UnstructuredContent(), schemaBlock)
		if err != nil {
			panic(err)
		}
		converted, err := schema.StateValueToJSONMap(rawConfig, schemaBlock.ImpliedType())
		if err != nil {
			panic(err)
		}

		deleted := make([]string, 0)
		for k, v := range converted {
			if v == nil {
				deleted = append(deleted, k)
			}
		}
		for _, k := range deleted {
			delete(converted, k)
		}
		re.Paved = *fieldpath.Pave(converted)
		exArr = append(exArr, re)
	}

	extNames, err := ExtractExternalNames(ctx, tfState)
	if err != nil {
		panic(err)
	}

	for i, re := range exArr {
		pc.Resources[resourceName].MetaResource.Examples[0] = re
		gen := examples.NewGenerator(filepath.Join(*output, strconv.FormatInt(int64(i), 10)), pc.ModulePath, pc.ShortName, pc.Resources)
		if err := gen.Generate(group, version, pc.Resources[resourceName]); err != nil {
			panic(err)
		}

		pm := gen.GetPavedWithManifest(fmt.Sprintf("%s.*", resourceName))

		if err := pm.Paved.SetValue("metadata.annotations", map[string]string{
			"crossplane.io/external-name": extNames[re.Name],
		}); err != nil {
			panic(err)
		}

		if err := pm.Paved.SetValue("spec.managementPolicies", []string{"Observe"}); err != nil {
			panic(err)
		}

		if err := pm.Paved.SetValue("spec.deletionPolicy", "Orphan"); err != nil {
			panic(err)
		}

		if err := gen.StoreExamples(); err != nil {
			panic(err)
		}
	}
}
