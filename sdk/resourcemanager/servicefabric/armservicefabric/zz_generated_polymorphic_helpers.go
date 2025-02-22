//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

import "encoding/json"

func unmarshalPartitionSchemeDescriptionClassification(rawMsg json.RawMessage) (PartitionSchemeDescriptionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PartitionSchemeDescriptionClassification
	switch m["partitionScheme"] {
	case string(PartitionSchemeNamed):
		b = &NamedPartitionSchemeDescription{}
	case string(PartitionSchemeSingleton):
		b = &SingletonPartitionSchemeDescription{}
	case string(PartitionSchemeUniformInt64Range):
		b = &UniformInt64RangePartitionSchemeDescription{}
	default:
		b = &PartitionSchemeDescription{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalServicePlacementPolicyDescriptionClassification(rawMsg json.RawMessage) (ServicePlacementPolicyDescriptionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServicePlacementPolicyDescriptionClassification
	switch m["type"] {
	default:
		b = &ServicePlacementPolicyDescription{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalServicePlacementPolicyDescriptionClassificationArray(rawMsg json.RawMessage) ([]ServicePlacementPolicyDescriptionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ServicePlacementPolicyDescriptionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalServicePlacementPolicyDescriptionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalServiceResourcePropertiesClassification(rawMsg json.RawMessage) (ServiceResourcePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServiceResourcePropertiesClassification
	switch m["serviceKind"] {
	case string(ServiceKindStateful):
		b = &StatefulServiceProperties{}
	case string(ServiceKindStateless):
		b = &StatelessServiceProperties{}
	default:
		b = &ServiceResourceProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalServiceResourceUpdatePropertiesClassification(rawMsg json.RawMessage) (ServiceResourceUpdatePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServiceResourceUpdatePropertiesClassification
	switch m["serviceKind"] {
	case string(ServiceKindStateful):
		b = &StatefulServiceUpdateProperties{}
	case string(ServiceKindStateless):
		b = &StatelessServiceUpdateProperties{}
	default:
		b = &ServiceResourceUpdateProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
