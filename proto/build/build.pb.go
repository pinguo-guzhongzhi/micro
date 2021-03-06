// Code generated by protoc-gen-go. DO NOT EDIT.
// source: build/build.proto

package build

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BuildRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Options              *Options `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildRequest) Reset()         { *m = BuildRequest{} }
func (m *BuildRequest) String() string { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()    {}
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d8124b77a3c5414, []int{0}
}

func (m *BuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildRequest.Unmarshal(m, b)
}
func (m *BuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildRequest.Marshal(b, m, deterministic)
}
func (m *BuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildRequest.Merge(m, src)
}
func (m *BuildRequest) XXX_Size() int {
	return xxx_messageInfo_BuildRequest.Size(m)
}
func (m *BuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildRequest proto.InternalMessageInfo

func (m *BuildRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BuildRequest) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

type Options struct {
	Archive              string   `protobuf:"bytes,1,opt,name=archive,proto3" json:"archive,omitempty"`
	Entrypoint           string   `protobuf:"bytes,2,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d8124b77a3c5414, []int{1}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetArchive() string {
	if m != nil {
		return m.Archive
	}
	return ""
}

func (m *Options) GetEntrypoint() string {
	if m != nil {
		return m.Entrypoint
	}
	return ""
}

type Result struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d8124b77a3c5414, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildRequest)(nil), "build.BuildRequest")
	proto.RegisterType((*Options)(nil), "build.Options")
	proto.RegisterType((*Result)(nil), "build.Result")
}

func init() { proto.RegisterFile("build/build.proto", fileDescriptor_3d8124b77a3c5414) }

var fileDescriptor_3d8124b77a3c5414 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2a, 0xcd, 0xcc,
	0x49, 0xd1, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x92, 0x0f,
	0x17, 0x8f, 0x13, 0x88, 0x11, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc4, 0xc5, 0x92,
	0x92, 0x58, 0x92, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x66, 0x0b, 0x69, 0x70, 0xb1,
	0xe7, 0x17, 0x94, 0x64, 0xe6, 0xe7, 0x15, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0xf1, 0xe9,
	0x41, 0x4c, 0xf2, 0x87, 0x88, 0x06, 0xc1, 0xa4, 0x95, 0x9c, 0xb9, 0xd8, 0xa1, 0x62, 0x42, 0x12,
	0x5c, 0xec, 0x89, 0x45, 0xc9, 0x19, 0x99, 0x65, 0xa9, 0x60, 0xb3, 0x38, 0x83, 0x60, 0x5c, 0x21,
	0x39, 0x2e, 0xae, 0xd4, 0xbc, 0x92, 0xa2, 0xca, 0x82, 0xfc, 0xcc, 0xbc, 0x12, 0xb0, 0x89, 0x9c,
	0x41, 0x48, 0x22, 0x4a, 0x32, 0x5c, 0x6c, 0x41, 0xa9, 0xc5, 0xa5, 0x39, 0x58, 0x1d, 0x63, 0x64,
	0xc5, 0xc5, 0x0a, 0x76, 0xb0, 0x90, 0x21, 0x8c, 0x21, 0x0c, 0x75, 0x0d, 0xb2, 0x3f, 0xa4, 0x78,
	0xa1, 0x82, 0x10, 0x93, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x9d, 0x74, 0xa3, 0xb4, 0xd3, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x73, 0x33, 0x93, 0x8b, 0xf2, 0xa1, 0x64,
	0x99, 0xb1, 0x3e, 0x38, 0x58, 0x20, 0x41, 0x64, 0x0d, 0x26, 0x93, 0xd8, 0xc0, 0x42, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xa0, 0xb7, 0x20, 0x3e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildClient is the client API for Build service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildClient interface {
	Build(ctx context.Context, opts ...grpc.CallOption) (Build_BuildClient, error)
}

type buildClient struct {
	cc *grpc.ClientConn
}

func NewBuildClient(cc *grpc.ClientConn) BuildClient {
	return &buildClient{cc}
}

func (c *buildClient) Build(ctx context.Context, opts ...grpc.CallOption) (Build_BuildClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Build_serviceDesc.Streams[0], "/build.Build/Build", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildBuildClient{stream}
	return x, nil
}

type Build_BuildClient interface {
	Send(*BuildRequest) error
	Recv() (*Result, error)
	grpc.ClientStream
}

type buildBuildClient struct {
	grpc.ClientStream
}

func (x *buildBuildClient) Send(m *BuildRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *buildBuildClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BuildServer is the server API for Build service.
type BuildServer interface {
	Build(Build_BuildServer) error
}

func RegisterBuildServer(s *grpc.Server, srv BuildServer) {
	s.RegisterService(&_Build_serviceDesc, srv)
}

func _Build_Build_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BuildServer).Build(&buildBuildServer{stream})
}

type Build_BuildServer interface {
	Send(*Result) error
	Recv() (*BuildRequest, error)
	grpc.ServerStream
}

type buildBuildServer struct {
	grpc.ServerStream
}

func (x *buildBuildServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *buildBuildServer) Recv() (*BuildRequest, error) {
	m := new(BuildRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Build_serviceDesc = grpc.ServiceDesc{
	ServiceName: "build.Build",
	HandlerType: (*BuildServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Build",
			Handler:       _Build_Build_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "build/build.proto",
}
