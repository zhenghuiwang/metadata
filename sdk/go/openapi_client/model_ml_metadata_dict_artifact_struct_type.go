/*
 * api/service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_client

type MlMetadataDictArtifactStructType struct {
	// Underlying properties for the type.
	Properties map[string]MlMetadataArtifactStructType `json:"properties,omitempty"`
	// If true, then if properties[\"foo\"] can be None, then that key is not required.
	NoneTypeNotRequired bool `json:"none_type_not_required,omitempty"`
	ExtraPropertiesType MlMetadataArtifactStructType `json:"extra_properties_type,omitempty"`
}
