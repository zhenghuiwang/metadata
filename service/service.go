// Copyright 2019 Google LLC
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

// Package service exports the type Service which implements the API defined by
// MetadataService.
package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	pb "github.com/kubeflow/metadata/api"
	"github.com/kubeflow/metadata/schemaparser"
	"github.com/xeipuuv/gojsonschema"

	"ml_metadata/metadata_store/mlmetadata"
	mlpb "ml_metadata/proto/metadata_store_go_proto"
)

const payloadKey = "__PAYLOAD__"

// Service implements the gRPC service MetadataService defined in the metadata
// API spec.
type Service struct {
	schemaset        *schemaparser.SchemaSet
	mlmd             *mlmetadata.Store
	typeNameToID     map[string]string
	typeNameToMLMDID map[string]int64
}

// NewService returns a metadata server
func NewService(schemaRootDir string) (*Service, error) {
	ss, err := schemaparser.NewSchemaSetFromADir(schemaRootDir)
	if err != nil {
		return nil, err
	}
	cfg := &mlpb.ConnectionConfig{Config: &mlpb.ConnectionConfig_Mysql{
		&mlpb.MySQLDatabaseConfig{
			Host:     proto.String("localhost"),
			Port:     proto.Uint32(3306),
			Database: proto.String("metadb"),
			User:     proto.String("guest"),
		},
	},
	}
	mlmd, err := mlmetadata.NewStore(cfg)
	typeNameToID := make(map[string]string)
	typeNameToMLMDID := make(map[string]int64)
	if err != nil {
		return nil, fmt.Errorf("failed to connect MySQL. config: %v, error: %s", cfg, err)
	}
	glog.Infof("len(schemas) = %d", len(ss.Schemas))
	for id := range ss.Schemas {
		tn, err := ss.TypeName(id)
		if err != nil {
			glog.Errorf("Ignored schema for missing 'kind', 'apiversion', or 'namespace'. schema $id = %s\n", id)
			continue
		}
		glog.Infof("Processng schema %s\n", tn)
		artifactType, err := getMLMDType(ss, id, tn)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to MLMD type: %s", err)
		}
		mlmdID, err := mlmd.PutArtifactType(
			artifactType,
			&mlmetadata.PutTypeOptions{AllFieldsMustMatch: true},
		)
		if err != nil {
			return nil, fmt.Errorf("error response from MLMD server: %s", err)
		}
		typeNameToID[tn] = id
		typeNameToMLMDID[tn] = int64(mlmdID)
		glog.Infof("Registered type: %+v\n", artifactType)

	}
	return &Service{
		schemaset:        ss,
		mlmd:             mlmd,
		typeNameToID:     typeNameToID,
		typeNameToMLMDID: typeNameToMLMDID,
	}, nil
}

func getMLMDType(ss *schemaparser.SchemaSet, id string, name string) (*mlpb.ArtifactType, error) {
	properties, err := ss.SimpleProperties(id)
	if err != nil {
		return nil, err
	}
	artifactType := &mlpb.ArtifactType{
		Name:       &name,
		Properties: make(map[string]mlpb.PropertyType),
	}
	for pname, ptype := range properties {
		switch ptype {
		case schemaparser.StringType:
			artifactType.Properties[pname] = mlpb.PropertyType_STRING
		case schemaparser.IntegerType:
			artifactType.Properties[pname] = mlpb.PropertyType_INT
		case schemaparser.NumberType:
			artifactType.Properties[pname] = mlpb.PropertyType_DOUBLE
		default:
			return nil, fmt.Errorf("unknown type %q for property %q", ptype, pname)
		}
	}
	artifactType.Properties[payloadKey] = mlpb.PropertyType_STRING
	return artifactType, nil
}

// GetResource returns the specified resource in the request.
func (s *Service) GetResource(ctx context.Context, in *pb.GetResourceRequest) (*pb.Resource, error) {
	if in.Name == "" {
		return nil, errors.New("must specify name")
	}

	return &pb.Resource{Name: in.Name}, nil
}

// CreateType registers a cutsomized type.
func (s *Service) CreateType(ctx context.Context, in *pb.CreateTypeRequest) (*pb.Type, error) {
	return &pb.Type{
		Id:     "1",
		Name:   in.Name,
		Schema: string(in.Schema),
	}, nil

}

// CreateArtifact creates an artifact.
func (s *Service) CreateArtifact(ctx context.Context, in *pb.Artifact) (*pb.Artifact, error) {
	id, exists := s.typeNameToID[in.GetTypeName()]
	if !exists {
		return nil, fmt.Errorf("type %q not registered", in.GetTypeName())
	}
	payload := make(map[string]interface{})
	err := json.Unmarshal(in.GetPayload(), &payload)
	if err != nil {
		return nil, fmt.Errorf("HTTP body is invalid JSON: %s", err)
	}
	if _, exists := payload["id"]; !exists {
		payload["id"] = ""
	}
	result, err := s.schemaset.Schemas[id].Validator.Validate(gojsonschema.NewBytesLoader(in.GetPayload()))
	if err != nil {
		return nil, fmt.Errorf("failed to validate payload: %s", err)
	}
	if !result.Valid() {
		return nil, fmt.Errorf("invalid payload: %+v", result.Errors())
	}

	simpleProperties, err := s.schemaset.SimpleProperties(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get SimpleProperties from schema. %s", err)
	}
	artifact := &mlpb.Artifact{
		TypeId:     proto.Int64(s.typeNameToMLMDID[in.GetTypeName()]),
		Properties: make(map[string]*mlpb.Value),
	}
	p := artifact.Properties
	for pname, ptype := range simpleProperties {
		val, exists := payload[pname]
		if !exists {
			continue
		}
		switch ptype {
		case schemaparser.StringType:
			p[pname] = &mlpb.Value{
				Value: &mlpb.Value_StringValue{
					StringValue: val.(string),
				},
			}
		case schemaparser.IntegerType:
			p[pname] = &mlpb.Value{
				Value: &mlpb.Value_IntValue{
					IntValue: val.(int64),
				},
			}
		case schemaparser.NumberType:
			p[pname] = &mlpb.Value{
				Value: &mlpb.Value_DoubleValue{
					DoubleValue: val.(float64),
				},
			}
		default:
			return nil, fmt.Errorf("unknown type %q of property %q", ptype, pname)
		}
	}
	p[payloadKey] = &mlpb.Value{
		Value: &mlpb.Value_StringValue{
			StringValue: string(in.GetPayload()),
		},
	}
	ids, err := s.mlmd.PutArtifacts([]*mlpb.Artifact{artifact})
	if err != nil {
		return nil, fmt.Errorf("Failed to put artifact into MLMD: %s", err)
	}
	return &pb.Artifact{
		Id:       fmt.Sprintf("%s", ids[0]),
		TypeName: in.GetTypeName(),
	}, nil
}
