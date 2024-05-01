// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/config"
)

type parsedState struct {
	// these two maps are maps from the discovered path of Terraform
	// state files to external names & output variables
	externalNames map[string]map[string]string
	references    map[string]map[string]any
}

func (ps *parsedState) getExtName(tfStatePath, k string) string {
	for tfState, m := range ps.externalNames {
		if tfState != tfStatePath {
			continue
		}
		return m[k]
	}
	return ""
}

func (ps *parsedState) resolveRef(tfStatePath, ref string) any {
	var resolved any
	// first try to resolve the reference locally using the HCL configuration's
	// own workspace.
	for tfState, m := range ps.references {
		if tfState != tfStatePath {
			continue
		}
		resolved = m[ref]
		break
	}
	if resolved != nil {
		return resolved
	}
	for tfState, m := range ps.references {
		if tfState == tfStatePath {
			continue
		}
		resolved = m[ref]
		if resolved != nil {
			return resolved
		}
	}
	return ref
}

func getStateFilePath(hclPath string) (string, error) {
	abs, err := filepath.Abs(hclPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get the absolute path for the HCL configuration file %s", hclPath)
	}
	return filepath.Join(filepath.Dir(abs), tfStateFile), nil
}

func (ps *parsedState) addStateFrom(ctx context.Context, tfStatePath string) error {
	provider, err := config.GetProvider(ctx, true)
	if err != nil {
		return errors.Wrap(err, "failed to get the provider configuration")
	}

	tfState, err := readInstanceStateFromFile(tfStatePath)
	if err != nil {
		return err
	}

	if tfState.Empty() || len(tfState.Modules) == 0 {
		return nil
	}

	absPath, err := filepath.Abs(tfStatePath)
	if err != nil {
		return errors.Wrapf(err, "failed to get the absolute path for the Terraform state file %s", tfStatePath)
	}

	refs := make(map[string]any, len(tfState.Modules[0].Outputs))
	for k, s := range tfState.Modules[0].Outputs {
		refs[k] = s.Value
		refs[fmt.Sprintf("data.terraform_remote_state.%s.outputs.%s", filepath.Base(filepath.Dir(absPath)), k)] = s.Value
	}
	if ps.references == nil {
		ps.references = make(map[string]map[string]any)
	}
	ps.references[absPath] = refs

	externalNames := make(map[string]string, len(tfState.Modules[0].Resources))
	for resEntryKey, tfRes := range tfState.Modules[0].Resources {
		// add attributes from the state to the reference resolver
		if tfRes.Primary != nil {
			for attr, v := range tfRes.Primary.Attributes {
				refs[fmt.Sprintf("%s.%s", resEntryKey, attr)] = v
			}
		}

		resType := tfRes.Type
		ujResourse, ok := provider.Resources[resType]
		if !ok {
			continue
		}

		tfStateJsonMap, err := fromInstanceStateToJSONMap(ujResourse.TerraformResource, tfRes.Primary)
		if err != nil {
			return errors.Wrapf(err, "error converting instance state to JSON for %s", resEntryKey)
		}

		externalName, err := ujResourse.ExternalName.GetExternalNameFn(tfStateJsonMap)
		if err != nil {
			return errors.Wrapf(err, "cannot get external name for %s", resEntryKey)
		}
		externalNames[resEntryKey] = externalName
	}

	if ps.externalNames == nil {
		ps.externalNames = make(map[string]map[string]string)
	}
	ps.externalNames[absPath] = externalNames
	return nil
}

func readInstanceStateFromFile(tfStateFilePath string) (*terraform.State, error) {
	tfStateFile, err := os.Open(tfStateFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening file")
	}
	defer func() { _ = tfStateFile.Close() }()

	tfStateBytes, err := io.ReadAll(tfStateFile)
	if err != nil {
		return nil, errors.Wrap(err, "error reading file")
	}

	var tfState terraform.State

	// Decode the JSON data into the map
	if err := json.Unmarshal(tfStateBytes, &tfState); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal TF state")
	}

	return &tfState, nil
}

func fromInstanceStateToJSONMap(tfRes *schema.Resource, newState *terraform.InstanceState) (map[string]interface{}, error) {
	impliedType := tfRes.CoreConfigSchema().ImpliedType()
	attrsAsCtyValue, err := newState.AttrsAsObjectValue(impliedType)
	if err != nil {
		return nil, errors.Wrap(err, "could not convert attrs to cty value")
	}
	stateValueMap, err := schema.StateValueToJSONMap(attrsAsCtyValue, impliedType)
	if err != nil {
		return nil, errors.Wrap(err, "could not convert instance state value to JSON")
	}
	return stateValueMap, nil
}
