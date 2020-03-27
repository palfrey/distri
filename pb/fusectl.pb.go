// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: fusectl.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fusectl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fusectl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_fusectl_proto_rawDescGZIP(), []int{0}
}

type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fusectl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_fusectl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_fusectl_proto_rawDescGZIP(), []int{1}
}

type MkdirAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dir *string `protobuf:"bytes,1,opt,name=dir" json:"dir,omitempty"`
}

func (x *MkdirAllRequest) Reset() {
	*x = MkdirAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fusectl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MkdirAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MkdirAllRequest) ProtoMessage() {}

func (x *MkdirAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fusectl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MkdirAllRequest.ProtoReflect.Descriptor instead.
func (*MkdirAllRequest) Descriptor() ([]byte, []int) {
	return file_fusectl_proto_rawDescGZIP(), []int{2}
}

func (x *MkdirAllRequest) GetDir() string {
	if x != nil && x.Dir != nil {
		return *x.Dir
	}
	return ""
}

type MkdirAllReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MkdirAllReply) Reset() {
	*x = MkdirAllReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fusectl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MkdirAllReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MkdirAllReply) ProtoMessage() {}

func (x *MkdirAllReply) ProtoReflect() protoreflect.Message {
	mi := &file_fusectl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MkdirAllReply.ProtoReflect.Descriptor instead.
func (*MkdirAllReply) Descriptor() ([]byte, []int) {
	return file_fusectl_proto_rawDescGZIP(), []int{3}
}

type ScanPackagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScanPackagesRequest) Reset() {
	*x = ScanPackagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fusectl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanPackagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanPackagesRequest) ProtoMessage() {}

func (x *ScanPackagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fusectl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanPackagesRequest.ProtoReflect.Descriptor instead.
func (*ScanPackagesRequest) Descriptor() ([]byte, []int) {
	return file_fusectl_proto_rawDescGZIP(), []int{4}
}

type ScanPackagesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScanPackagesReply) Reset() {
	*x = ScanPackagesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fusectl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanPackagesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanPackagesReply) ProtoMessage() {}

func (x *ScanPackagesReply) ProtoReflect() protoreflect.Message {
	mi := &file_fusectl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanPackagesReply.ProtoReflect.Descriptor instead.
func (*ScanPackagesReply) Descriptor() ([]byte, []int) {
	return file_fusectl_proto_rawDescGZIP(), []int{5}
}

var File_fusectl_proto protoreflect.FileDescriptor

var file_fusectl_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x75, 0x73, 0x65, 0x63, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x23, 0x0a, 0x0f, 0x4d, 0x6b, 0x64, 0x69, 0x72, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x64, 0x69, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x4d, 0x6b, 0x64, 0x69, 0x72, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11,
	0x53, 0x63, 0x61, 0x6e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0xa8, 0x01, 0x0a, 0x04, 0x46, 0x55, 0x53, 0x45, 0x12, 0x28, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x4d, 0x6b, 0x64, 0x69, 0x72, 0x41, 0x6c, 0x6c,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6b, 0x64, 0x69, 0x72, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6b, 0x64, 0x69, 0x72,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0c, 0x53, 0x63,
	0x61, 0x6e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x3b, 0x70, 0x62,
}

var (
	file_fusectl_proto_rawDescOnce sync.Once
	file_fusectl_proto_rawDescData = file_fusectl_proto_rawDesc
)

func file_fusectl_proto_rawDescGZIP() []byte {
	file_fusectl_proto_rawDescOnce.Do(func() {
		file_fusectl_proto_rawDescData = protoimpl.X.CompressGZIP(file_fusectl_proto_rawDescData)
	})
	return file_fusectl_proto_rawDescData
}

var file_fusectl_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_fusectl_proto_goTypes = []interface{}{
	(*PingRequest)(nil),         // 0: pb.PingRequest
	(*PingReply)(nil),           // 1: pb.PingReply
	(*MkdirAllRequest)(nil),     // 2: pb.MkdirAllRequest
	(*MkdirAllReply)(nil),       // 3: pb.MkdirAllReply
	(*ScanPackagesRequest)(nil), // 4: pb.ScanPackagesRequest
	(*ScanPackagesReply)(nil),   // 5: pb.ScanPackagesReply
}
var file_fusectl_proto_depIdxs = []int32{
	0, // 0: pb.FUSE.Ping:input_type -> pb.PingRequest
	2, // 1: pb.FUSE.MkdirAll:input_type -> pb.MkdirAllRequest
	4, // 2: pb.FUSE.ScanPackages:input_type -> pb.ScanPackagesRequest
	1, // 3: pb.FUSE.Ping:output_type -> pb.PingReply
	3, // 4: pb.FUSE.MkdirAll:output_type -> pb.MkdirAllReply
	5, // 5: pb.FUSE.ScanPackages:output_type -> pb.ScanPackagesReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fusectl_proto_init() }
func file_fusectl_proto_init() {
	if File_fusectl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fusectl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fusectl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fusectl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MkdirAllRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fusectl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MkdirAllReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fusectl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanPackagesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fusectl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanPackagesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fusectl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fusectl_proto_goTypes,
		DependencyIndexes: file_fusectl_proto_depIdxs,
		MessageInfos:      file_fusectl_proto_msgTypes,
	}.Build()
	File_fusectl_proto = out.File
	file_fusectl_proto_rawDesc = nil
	file_fusectl_proto_goTypes = nil
	file_fusectl_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FUSEClient is the client API for FUSE service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FUSEClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	// MkdirAll creates the specified directory in the root of the mountpoint
	// (e.g. /ro/systemd-amd64-239). This is useful for bind-mounting
	// DESTDIR/PREFIX to PREFIX when building packages.
	MkdirAll(ctx context.Context, in *MkdirAllRequest, opts ...grpc.CallOption) (*MkdirAllReply, error)
	// ScanPackages discovers new packages in the mounted repository. This is
	// called by “distri install”.
	ScanPackages(ctx context.Context, in *ScanPackagesRequest, opts ...grpc.CallOption) (*ScanPackagesReply, error)
}

type fUSEClient struct {
	cc grpc.ClientConnInterface
}

func NewFUSEClient(cc grpc.ClientConnInterface) FUSEClient {
	return &fUSEClient{cc}
}

func (c *fUSEClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/pb.FUSE/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fUSEClient) MkdirAll(ctx context.Context, in *MkdirAllRequest, opts ...grpc.CallOption) (*MkdirAllReply, error) {
	out := new(MkdirAllReply)
	err := c.cc.Invoke(ctx, "/pb.FUSE/MkdirAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fUSEClient) ScanPackages(ctx context.Context, in *ScanPackagesRequest, opts ...grpc.CallOption) (*ScanPackagesReply, error) {
	out := new(ScanPackagesReply)
	err := c.cc.Invoke(ctx, "/pb.FUSE/ScanPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FUSEServer is the server API for FUSE service.
type FUSEServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	// MkdirAll creates the specified directory in the root of the mountpoint
	// (e.g. /ro/systemd-amd64-239). This is useful for bind-mounting
	// DESTDIR/PREFIX to PREFIX when building packages.
	MkdirAll(context.Context, *MkdirAllRequest) (*MkdirAllReply, error)
	// ScanPackages discovers new packages in the mounted repository. This is
	// called by “distri install”.
	ScanPackages(context.Context, *ScanPackagesRequest) (*ScanPackagesReply, error)
}

// UnimplementedFUSEServer can be embedded to have forward compatible implementations.
type UnimplementedFUSEServer struct {
}

func (*UnimplementedFUSEServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedFUSEServer) MkdirAll(context.Context, *MkdirAllRequest) (*MkdirAllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MkdirAll not implemented")
}
func (*UnimplementedFUSEServer) ScanPackages(context.Context, *ScanPackagesRequest) (*ScanPackagesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanPackages not implemented")
}

func RegisterFUSEServer(s *grpc.Server, srv FUSEServer) {
	s.RegisterService(&_FUSE_serviceDesc, srv)
}

func _FUSE_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FUSEServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FUSE/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FUSEServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FUSE_MkdirAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkdirAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FUSEServer).MkdirAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FUSE/MkdirAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FUSEServer).MkdirAll(ctx, req.(*MkdirAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FUSE_ScanPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FUSEServer).ScanPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FUSE/ScanPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FUSEServer).ScanPackages(ctx, req.(*ScanPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FUSE_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FUSE",
	HandlerType: (*FUSEServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _FUSE_Ping_Handler,
		},
		{
			MethodName: "MkdirAll",
			Handler:    _FUSE_MkdirAll_Handler,
		},
		{
			MethodName: "ScanPackages",
			Handler:    _FUSE_ScanPackages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fusectl.proto",
}
