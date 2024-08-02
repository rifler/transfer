// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/priv/iam/v1/token_agent.proto

package iam

import (
	_ "github.com/doublecloud/tross/cloud/bitbucket/private-api/yandex/cloud/priv"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An additional service qualifier for obtaining different tokens
	// for different services running on behalf of the same user.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_token_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_token_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokenRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

// The response message containing the token
type GetTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IamToken  string                 `protobuf:"bytes,1,opt,name=iam_token,json=iamToken,proto3" json:"iam_token,omitempty"`
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *GetTokenResponse) Reset() {
	*x = GetTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_token_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenResponse) ProtoMessage() {}

func (x *GetTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_token_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenResponse.ProtoReflect.Descriptor instead.
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenResponse) GetIamToken() string {
	if x != nil {
		return x.IamToken
	}
	return ""
}

func (x *GetTokenResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

var File_yandex_cloud_priv_iam_v1_token_agent_proto protoreflect.FileDescriptor

var file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22,
	0x74, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x69, 0x61, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xc8, 0x8f, 0x31, 0x01, 0xd0, 0x8f, 0x31, 0x02,
	0x52, 0x08, 0x69, 0x61, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x41, 0x74, 0x32, 0x71, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x51, 0x42, 0x04, 0x50, 0x54, 0x41, 0x53,
	0x5a, 0x49, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x69, 0x76,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDescOnce sync.Once
	file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDescData = file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDesc
)

func file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDescGZIP() []byte {
	file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDescData)
	})
	return file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDescData
}

var file_yandex_cloud_priv_iam_v1_token_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_priv_iam_v1_token_agent_proto_goTypes = []interface{}{
	(*GetTokenRequest)(nil),       // 0: yandex.cloud.priv.iam.v1.GetTokenRequest
	(*GetTokenResponse)(nil),      // 1: yandex.cloud.priv.iam.v1.GetTokenResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_yandex_cloud_priv_iam_v1_token_agent_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.priv.iam.v1.GetTokenResponse.expires_at:type_name -> google.protobuf.Timestamp
	0, // 1: yandex.cloud.priv.iam.v1.TokenAgent.GetToken:input_type -> yandex.cloud.priv.iam.v1.GetTokenRequest
	1, // 2: yandex.cloud.priv.iam.v1.TokenAgent.GetToken:output_type -> yandex.cloud.priv.iam.v1.GetTokenResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_priv_iam_v1_token_agent_proto_init() }
func file_yandex_cloud_priv_iam_v1_token_agent_proto_init() {
	if File_yandex_cloud_priv_iam_v1_token_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_priv_iam_v1_token_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenRequest); i {
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
		file_yandex_cloud_priv_iam_v1_token_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenResponse); i {
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
			RawDescriptor: file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_priv_iam_v1_token_agent_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_priv_iam_v1_token_agent_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_priv_iam_v1_token_agent_proto_msgTypes,
	}.Build()
	File_yandex_cloud_priv_iam_v1_token_agent_proto = out.File
	file_yandex_cloud_priv_iam_v1_token_agent_proto_rawDesc = nil
	file_yandex_cloud_priv_iam_v1_token_agent_proto_goTypes = nil
	file_yandex_cloud_priv_iam_v1_token_agent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokenAgentClient is the client API for TokenAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenAgentClient interface {
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
}

type tokenAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenAgentClient(cc grpc.ClientConnInterface) TokenAgentClient {
	return &tokenAgentClient{cc}
}

func (c *tokenAgentClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.priv.iam.v1.TokenAgent/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenAgentServer is the server API for TokenAgent service.
type TokenAgentServer interface {
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
}

// UnimplementedTokenAgentServer can be embedded to have forward compatible implementations.
type UnimplementedTokenAgentServer struct {
}

func (*UnimplementedTokenAgentServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}

func RegisterTokenAgentServer(s *grpc.Server, srv TokenAgentServer) {
	s.RegisterService(&_TokenAgent_serviceDesc, srv)
}

func _TokenAgent_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenAgentServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.priv.iam.v1.TokenAgent/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenAgentServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.priv.iam.v1.TokenAgent",
	HandlerType: (*TokenAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _TokenAgent_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/priv/iam/v1/token_agent.proto",
}