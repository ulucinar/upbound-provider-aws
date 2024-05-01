// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"

	"github.com/upbound/provider-aws/config"
)

type parsedState struct {
	externalNames   map[string]string
	outputVariables map[string]any
}

func readInstanceStateFromFile(tfStateFilePath string) (*terraform.State, error) {
	tfStateFile, err := os.Open(tfStateFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening file")
	}
	defer tfStateFile.Close()

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

func extractExternalNames(ctx context.Context, tfStatePath string) (*parsedState, error) {
	parsed := &parsedState{}
	provider, err := config.GetProvider(ctx, true)
	if err != nil {
		return nil, err
	}

	tfState, err := readInstanceStateFromFile(tfStatePath)
	if err != nil {
		return nil, err
	}

	if tfState.Empty() {
		return parsed, nil
	}

	if len(tfState.Modules) == 0 {
		return parsed, nil
	}

	parsed.externalNames = make(map[string]string, len(tfState.Modules[0].Resources))
	for resEntryKey, tfRes := range tfState.Modules[0].Resources {
		resType := tfRes.Type
		ujResourse, ok := provider.Resources[resType]
		if !ok {
			continue
		}

		tfStateJsonMap, err := fromInstanceStateToJSONMap(ujResourse.TerraformResource, tfRes.Primary)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting instance state to JSON for %s", resEntryKey)
		}

		externalName, err := ujResourse.ExternalName.GetExternalNameFn(tfStateJsonMap)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot get external name for %s", resEntryKey)
		}
		parsed.externalNames[resEntryKey] = externalName
	}

	parsed.outputVariables = make(map[string]any, len(tfState.Modules[0].Outputs))
	for k, s := range tfState.Modules[0].Outputs {
		parsed.outputVariables[k] = s.Value
	}
	return parsed, nil
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
