// Code generated by protoc-gen-go.
// source: indexbackend.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	indexbackend.proto
	sourcebackend.proto

It has these top-level messages:
	FilesRequest
	FilesReply
	ReplaceIndexRequest
	ReplaceIndexReply
	FileRequest
	FileReply
	SearchRequest
	Match
	ProgressUpdate
	SearchReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

type FilesRequest struct {
	// Text query (e.g. “i3Font”) which will be translated into a trigram query
	// (e.g. "3Fo" "Fon" "i3F" "ont").
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *FilesRequest) Reset()                    { *m = FilesRequest{} }
func (m *FilesRequest) String() string            { return proto1.CompactTextString(m) }
func (*FilesRequest) ProtoMessage()               {}
func (*FilesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FilesReply struct {
	// A list of paths which match the requested trigram query (likely to match
	// the regular expression from which the trigram query was derived, but can
	// contain false positives).
	Path []string `protobuf:"bytes,1,rep,name=path" json:"path,omitempty"`
}

func (m *FilesReply) Reset()                    { *m = FilesReply{} }
func (m *FilesReply) String() string            { return proto1.CompactTextString(m) }
func (*FilesReply) ProtoMessage()               {}
func (*FilesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ReplaceIndexRequest struct {
	ReplacementPath string `protobuf:"bytes,1,opt,name=replacement_path,json=replacementPath" json:"replacement_path,omitempty"`
}

func (m *ReplaceIndexRequest) Reset()                    { *m = ReplaceIndexRequest{} }
func (m *ReplaceIndexRequest) String() string            { return proto1.CompactTextString(m) }
func (*ReplaceIndexRequest) ProtoMessage()               {}
func (*ReplaceIndexRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ReplaceIndexReply struct {
}

func (m *ReplaceIndexReply) Reset()                    { *m = ReplaceIndexReply{} }
func (m *ReplaceIndexReply) String() string            { return proto1.CompactTextString(m) }
func (*ReplaceIndexReply) ProtoMessage()               {}
func (*ReplaceIndexReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto1.RegisterType((*FilesRequest)(nil), "proto.FilesRequest")
	proto1.RegisterType((*FilesReply)(nil), "proto.FilesReply")
	proto1.RegisterType((*ReplaceIndexRequest)(nil), "proto.ReplaceIndexRequest")
	proto1.RegisterType((*ReplaceIndexReply)(nil), "proto.ReplaceIndexReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for IndexBackend service

type IndexBackendClient interface {
	// Files returns a list of files which match the specified query in the
	// trigram index.
	Files(ctx context.Context, in *FilesRequest, opts ...grpc.CallOption) (*FilesReply, error)
	// Replaces the loaded index with the specified replacement index. On a file
	// system level, the specified file is mv'ed to the file specified by
	// -index_path.
	ReplaceIndex(ctx context.Context, in *ReplaceIndexRequest, opts ...grpc.CallOption) (*ReplaceIndexReply, error)
}

type indexBackendClient struct {
	cc *grpc.ClientConn
}

func NewIndexBackendClient(cc *grpc.ClientConn) IndexBackendClient {
	return &indexBackendClient{cc}
}

func (c *indexBackendClient) Files(ctx context.Context, in *FilesRequest, opts ...grpc.CallOption) (*FilesReply, error) {
	out := new(FilesReply)
	err := grpc.Invoke(ctx, "/proto.IndexBackend/Files", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexBackendClient) ReplaceIndex(ctx context.Context, in *ReplaceIndexRequest, opts ...grpc.CallOption) (*ReplaceIndexReply, error) {
	out := new(ReplaceIndexReply)
	err := grpc.Invoke(ctx, "/proto.IndexBackend/ReplaceIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IndexBackend service

type IndexBackendServer interface {
	// Files returns a list of files which match the specified query in the
	// trigram index.
	Files(context.Context, *FilesRequest) (*FilesReply, error)
	// Replaces the loaded index with the specified replacement index. On a file
	// system level, the specified file is mv'ed to the file specified by
	// -index_path.
	ReplaceIndex(context.Context, *ReplaceIndexRequest) (*ReplaceIndexReply, error)
}

func RegisterIndexBackendServer(s *grpc.Server, srv IndexBackendServer) {
	s.RegisterService(&_IndexBackend_serviceDesc, srv)
}

func _IndexBackend_Files_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(FilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(IndexBackendServer).Files(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _IndexBackend_ReplaceIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ReplaceIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(IndexBackendServer).ReplaceIndex(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _IndexBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.IndexBackend",
	HandlerType: (*IndexBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Files",
			Handler:    _IndexBackend_Files_Handler,
		},
		{
			MethodName: "ReplaceIndex",
			Handler:    _IndexBackend_ReplaceIndex_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcc, 0x4b, 0x49,
	0xad, 0x48, 0x4a, 0x4c, 0xce, 0x4e, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x05, 0x53, 0x4a, 0x2a, 0x5c, 0x3c, 0x6e, 0x99, 0x39, 0xa9, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x22, 0x5c, 0xac, 0x40, 0x46, 0x51, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x84, 0xa3, 0xa4, 0xc0, 0xc5, 0x05, 0x55, 0x55, 0x90, 0x53, 0x29, 0x24, 0xc4, 0xc5, 0x52,
	0x90, 0x58, 0x92, 0x01, 0x54, 0xc2, 0x0c, 0x54, 0x02, 0x66, 0x2b, 0x39, 0x70, 0x09, 0x83, 0x24,
	0x13, 0x93, 0x53, 0x3d, 0x41, 0x76, 0xc1, 0x8c, 0xd3, 0xe4, 0x12, 0x28, 0x82, 0x08, 0xe7, 0xa6,
	0xe6, 0x95, 0xc4, 0x43, 0xb5, 0x81, 0x4c, 0xe6, 0x47, 0x12, 0x0f, 0x00, 0x99, 0x20, 0xcc, 0x25,
	0x88, 0x6a, 0x02, 0xd0, 0x2a, 0xa3, 0x4e, 0x46, 0x2e, 0x1e, 0x30, 0xd7, 0x09, 0xe2, 0x78, 0x21,
	0x43, 0x2e, 0x56, 0xb0, 0x4b, 0x84, 0x84, 0x21, 0xfe, 0xd0, 0x43, 0x76, 0xbd, 0x94, 0x20, 0xaa,
	0x20, 0xd0, 0x04, 0x25, 0x06, 0x21, 0x37, 0x2e, 0x1e, 0x64, 0x83, 0x85, 0xa4, 0xa0, 0x8a, 0xb0,
	0xb8, 0x57, 0x4a, 0x02, 0xab, 0x1c, 0xd8, 0x9c, 0x24, 0x36, 0xb0, 0x94, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xfb, 0x13, 0x94, 0x3a, 0x4e, 0x01, 0x00, 0x00,
}