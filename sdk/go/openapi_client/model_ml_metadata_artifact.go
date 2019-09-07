/*
 * api/service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_client

type MlMetadataArtifact struct {
	// The id of the artifact.
	Id string `json:"id,omitempty"`
	// The id of an ArtifactType. Type must be specified when an artifact is created, and it cannot be changed.
	TypeId string `json:"type_id,omitempty"`
	// The uniform resource identifier of the physical artifact. May be empty if there is no physical artifact.
	Uri string `json:"uri,omitempty"`
	// Properties of the artifact. Properties must be specified in the ArtifactType.
	Properties map[string]MlMetadataValue `json:"properties,omitempty"`
	// User provided custom properties which are not defined by its type.
	CustomProperties map[string]MlMetadataValue `json:"custom_properties,omitempty"`
}
