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
	"errors"

	"github.com/golang/glog"
	pb "github.com/kubeflow/metadata/api"
)

// Service implements the gRPC service MetadataService defined in the metadata
// API spec.
type Service struct{}

// GetResource returns the specified resource in the request.
func (s *Service) GetResource(ctx context.Context, in *pb.GetResourceRequest) (*pb.Resource, error) {
	if in.Name == "" {
		return nil, errors.New("must specify name")
	}

	return &pb.Resource{Name: in.Name}, nil
}
func (s *Service) CreateType(ctx context.Context, in *pb.CreateTypeRequest) (*pb.Type, error) {
	glog.Infof("request: %s", in)
	return &pb.Type{
		Id:     "1",
		Name:   in.Name,
		Schema: string(in.Schema.Value),
	}, nil
}

func (s *Service) GetType(ctx context.Context, in *pb.GetTypeRequest) (*pb.Type, error) {
	return nil, nil
}
