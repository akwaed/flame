// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DesignSchema - Schema to define the roles and their connections
type DesignSchema struct {
	Revision int32 `json:"revision,omitempty"`

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Roles []Role `json:"roles"`

	Channels []Channel `json:"channels"`

	Connectors []Connector `json:"connectors,omitempty"`
}

// AssertDesignSchemaRequired checks if the required fields are not zero-ed
func AssertDesignSchemaRequired(obj DesignSchema) error {
	elements := map[string]interface{}{
		"name":     obj.Name,
		"roles":    obj.Roles,
		"channels": obj.Channels,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Roles {
		if err := AssertRoleRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Channels {
		if err := AssertChannelRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Connectors {
		if err := AssertConnectorRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertDesignSchemaConstraints checks if the values respects the defined constraints
func AssertDesignSchemaConstraints(obj DesignSchema) error {
	return nil
}
