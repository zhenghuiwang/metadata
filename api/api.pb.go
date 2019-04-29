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

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Resource struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f90b1c968afc262a, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (dst *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(dst, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResourceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResourceRequest) Reset()         { *m = GetResourceRequest{} }
func (m *GetResourceRequest) String() string { return proto.CompactTextString(m) }
func (*GetResourceRequest) ProtoMessage()    {}
func (*GetResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f90b1c968afc262a, []int{1}
}
func (m *GetResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResourceRequest.Unmarshal(m, b)
}
func (m *GetResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResourceRequest.Marshal(b, m, deterministic)
}
func (dst *GetResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourceRequest.Merge(dst, src)
}
func (m *GetResourceRequest) XXX_Size() int {
	return xxx_messageInfo_GetResourceRequest.Size(m)
}
func (m *GetResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourceRequest proto.InternalMessageInfo

func (m *GetResourceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Type struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Schema               string   `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f90b1c968afc262a, []int{2}
}
func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (dst *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(dst, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

func (m *Type) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Type) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

type CreateTypeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Schema               []byte   `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTypeRequest) Reset()         { *m = CreateTypeRequest{} }
func (m *CreateTypeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTypeRequest) ProtoMessage()    {}
func (*CreateTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f90b1c968afc262a, []int{3}
}
func (m *CreateTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTypeRequest.Unmarshal(m, b)
}
func (m *CreateTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTypeRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTypeRequest.Merge(dst, src)
}
func (m *CreateTypeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTypeRequest.Size(m)
}
func (m *CreateTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTypeRequest proto.InternalMessageInfo

func (m *CreateTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTypeRequest) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

type Artifact struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeName             string   `protobuf:"bytes,2,opt,name=typeName,proto3" json:"typeName,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f90b1c968afc262a, []int{4}
}
func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (dst *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(dst, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Artifact) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

func (m *Artifact) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Resource)(nil), "api.Resource")
	proto.RegisterType((*GetResourceRequest)(nil), "api.GetResourceRequest")
	proto.RegisterType((*Type)(nil), "api.Type")
	proto.RegisterType((*CreateTypeRequest)(nil), "api.CreateTypeRequest")
	proto.RegisterType((*Artifact)(nil), "api.Artifact")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetadataServiceClient is the client API for MetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataServiceClient interface {
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error)
	CreateType(ctx context.Context, in *CreateTypeRequest, opts ...grpc.CallOption) (*Type, error)
	CreateArtifact(ctx context.Context, in *Artifact, opts ...grpc.CallOption) (*Artifact, error)
}

type metadataServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetadataServiceClient(cc *grpc.ClientConn) MetadataServiceClient {
	return &metadataServiceClient{cc}
}

func (c *metadataServiceClient) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := c.cc.Invoke(ctx, "/api.MetadataService/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) CreateType(ctx context.Context, in *CreateTypeRequest, opts ...grpc.CallOption) (*Type, error) {
	out := new(Type)
	err := c.cc.Invoke(ctx, "/api.MetadataService/CreateType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) CreateArtifact(ctx context.Context, in *Artifact, opts ...grpc.CallOption) (*Artifact, error) {
	out := new(Artifact)
	err := c.cc.Invoke(ctx, "/api.MetadataService/CreateArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServiceServer is the server API for MetadataService service.
type MetadataServiceServer interface {
	GetResource(context.Context, *GetResourceRequest) (*Resource, error)
	CreateType(context.Context, *CreateTypeRequest) (*Type, error)
	CreateArtifact(context.Context, *Artifact) (*Artifact, error)
}

func RegisterMetadataServiceServer(s *grpc.Server, srv MetadataServiceServer) {
	s.RegisterService(&_MetadataService_serviceDesc, srv)
}

func _MetadataService_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MetadataService/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetResource(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_CreateType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).CreateType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MetadataService/CreateType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).CreateType(ctx, req.(*CreateTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_CreateArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Artifact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).CreateArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MetadataService/CreateArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).CreateArtifact(ctx, req.(*Artifact))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetadataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MetadataService",
	HandlerType: (*MetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResource",
			Handler:    _MetadataService_GetResource_Handler,
		},
		{
			MethodName: "CreateType",
			Handler:    _MetadataService_CreateType_Handler,
		},
		{
			MethodName: "CreateArtifact",
			Handler:    _MetadataService_CreateArtifact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_api_f90b1c968afc262a) }

var fileDescriptor_api_f90b1c968afc262a = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x15, 0x06, 0x51, 0x98, 0x02, 0x55, 0xf7, 0x00, 0xd4, 0xaa, 0xda, 0x6a, 0x4f, 0x88, 0x03,
	0xab, 0x96, 0x5b, 0x2b, 0x54, 0x25, 0x48, 0xc9, 0x25, 0x89, 0x22, 0x27, 0xca, 0x7d, 0x63, 0x4f,
	0xc8, 0x2a, 0xe0, 0xdd, 0xd8, 0x0b, 0x12, 0x22, 0x5c, 0xf2, 0x0b, 0xf9, 0xb4, 0x9c, 0x73, 0xcb,
	0x87, 0x44, 0xbb, 0xf6, 0x1a, 0x22, 0xa2, 0xe4, 0x62, 0x7b, 0xe6, 0x3d, 0xbf, 0x37, 0xf3, 0x76,
	0xa1, 0xc9, 0x95, 0x60, 0x5c, 0x89, 0x81, 0x4a, 0xa4, 0x96, 0xa4, 0xcc, 0x95, 0xf0, 0xbf, 0x4f,
	0xa4, 0x9c, 0x4c, 0x91, 0x59, 0x28, 0x8e, 0xa5, 0xe6, 0x5a, 0xc8, 0x38, 0xcd, 0x28, 0xf4, 0x07,
	0xd4, 0x02, 0x4c, 0xe5, 0x3c, 0x09, 0x91, 0x10, 0xa8, 0xc4, 0x7c, 0x86, 0xdd, 0xd2, 0xaf, 0x52,
	0xaf, 0x1e, 0xd8, 0x6f, 0xda, 0x03, 0x72, 0x88, 0xda, 0x51, 0x02, 0xbc, 0x9d, 0x63, 0xaa, 0xdf,
	0x64, 0xee, 0x43, 0xe5, 0x7c, 0xa9, 0x90, 0xb4, 0xc0, 0x13, 0x51, 0x8e, 0x78, 0x22, 0x2a, 0xb8,
	0xde, 0x86, 0x4b, 0xda, 0x50, 0x4d, 0xc3, 0x6b, 0x9c, 0xf1, 0x6e, 0xd9, 0x76, 0xf3, 0x8a, 0xfe,
	0x87, 0xaf, 0xe3, 0x04, 0xb9, 0x46, 0xa3, 0xf4, 0x8e, 0xd9, 0x96, 0x80, 0x91, 0x6d, 0x14, 0x02,
	0xa7, 0x50, 0xdb, 0x4b, 0xb4, 0xb8, 0xe2, 0xa1, 0xde, 0x19, 0xc4, 0x87, 0x9a, 0x5e, 0x2a, 0x3c,
	0xd9, 0x0c, 0x53, 0xd4, 0xa4, 0x0b, 0x9f, 0x14, 0x5f, 0x4e, 0x25, 0x8f, 0xec, 0x44, 0x8d, 0xc0,
	0x95, 0x7f, 0x9e, 0x3c, 0xf8, 0x72, 0x8c, 0x9a, 0x47, 0x5c, 0xf3, 0x33, 0x4c, 0x16, 0x22, 0x44,
	0x72, 0x01, 0x9f, 0xb7, 0x42, 0x21, 0x9d, 0x81, 0x89, 0x7c, 0x37, 0x26, 0xbf, 0x69, 0x01, 0xd7,
	0xa5, 0x3f, 0xef, 0x1f, 0x9f, 0x1f, 0xbc, 0x6f, 0xa4, 0x63, 0xcf, 0x62, 0xf1, 0x9b, 0x25, 0x39,
	0xc2, 0x56, 0x66, 0xa9, 0x35, 0xd1, 0x00, 0x9b, 0xf5, 0x49, 0xdb, 0xfe, 0xbd, 0x93, 0x87, 0x5f,
	0xb7, 0x7d, 0xd3, 0xa1, 0x63, 0xab, 0x38, 0xa2, 0x43, 0xa7, 0x68, 0x16, 0xca, 0xd4, 0x46, 0xe6,
	0x91, 0x2a, 0x1e, 0x62, 0xca, 0xfa, 0xec, 0x46, 0xc4, 0x91, 0x79, 0x2f, 0x30, 0x49, 0xcd, 0xe9,
	0xb3, 0xfe, 0xfa, 0x6f, 0x9e, 0x19, 0xb9, 0x83, 0x56, 0x66, 0x52, 0x24, 0x97, 0xcd, 0xed, 0x4a,
	0xff, 0x75, 0x49, 0x8f, 0xac, 0xe9, 0x01, 0xfd, 0xe7, 0x4c, 0x79, 0x8e, 0xb0, 0x95, 0xcb, 0xf3,
	0x63, 0x73, 0x97, 0xef, 0x65, 0xd5, 0xde, 0xc3, 0xe1, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8,
	0x80, 0xbf, 0x78, 0xbb, 0x02, 0x00, 0x00,
}
