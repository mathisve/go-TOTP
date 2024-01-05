// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: totp/totp.proto

package __

import (
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

type GetSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSecretRequest) Reset() {
	*x = GetSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_totp_totp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretRequest) ProtoMessage() {}

func (x *GetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_totp_totp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretRequest.ProtoReflect.Descriptor instead.
func (*GetSecretRequest) Descriptor() ([]byte, []int) {
	return file_totp_totp_proto_rawDescGZIP(), []int{0}
}

type GetSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretId []byte `protobuf:"bytes,1,opt,name=secretId,proto3" json:"secretId,omitempty"`
	Secret   []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *GetSecretResponse) Reset() {
	*x = GetSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_totp_totp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretResponse) ProtoMessage() {}

func (x *GetSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_totp_totp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretResponse.ProtoReflect.Descriptor instead.
func (*GetSecretResponse) Descriptor() ([]byte, []int) {
	return file_totp_totp_proto_rawDescGZIP(), []int{1}
}

func (x *GetSecretResponse) GetSecretId() []byte {
	if x != nil {
		return x.SecretId
	}
	return nil
}

func (x *GetSecretResponse) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

type ChallengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretId []byte `protobuf:"bytes,1,opt,name=secretId,proto3" json:"secretId,omitempty"`
	Topt     []byte `protobuf:"bytes,2,opt,name=topt,proto3" json:"topt,omitempty"`
}

func (x *ChallengeRequest) Reset() {
	*x = ChallengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_totp_totp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeRequest) ProtoMessage() {}

func (x *ChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_totp_totp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeRequest.ProtoReflect.Descriptor instead.
func (*ChallengeRequest) Descriptor() ([]byte, []int) {
	return file_totp_totp_proto_rawDescGZIP(), []int{2}
}

func (x *ChallengeRequest) GetSecretId() []byte {
	if x != nil {
		return x.SecretId
	}
	return nil
}

func (x *ChallengeRequest) GetTopt() []byte {
	if x != nil {
		return x.Topt
	}
	return nil
}

type ChallengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ChallengeResponse) Reset() {
	*x = ChallengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_totp_totp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeResponse) ProtoMessage() {}

func (x *ChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_totp_totp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeResponse.ProtoReflect.Descriptor instead.
func (*ChallengeResponse) Descriptor() ([]byte, []int) {
	return file_totp_totp_proto_rawDescGZIP(), []int{3}
}

func (x *ChallengeResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_totp_totp_proto protoreflect.FileDescriptor

var file_totp_totp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x70, 0x2f, 0x74, 0x6f, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x67, 0x6f, 0x74, 0x6f, 0x70, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x70, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32,
	0x8c, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x70, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x70, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09,
	0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x6f, 0x74, 0x6f,
	0x70, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x70, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04,
	0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_totp_totp_proto_rawDescOnce sync.Once
	file_totp_totp_proto_rawDescData = file_totp_totp_proto_rawDesc
)

func file_totp_totp_proto_rawDescGZIP() []byte {
	file_totp_totp_proto_rawDescOnce.Do(func() {
		file_totp_totp_proto_rawDescData = protoimpl.X.CompressGZIP(file_totp_totp_proto_rawDescData)
	})
	return file_totp_totp_proto_rawDescData
}

var file_totp_totp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_totp_totp_proto_goTypes = []interface{}{
	(*GetSecretRequest)(nil),  // 0: gotopt.GetSecretRequest
	(*GetSecretResponse)(nil), // 1: gotopt.GetSecretResponse
	(*ChallengeRequest)(nil),  // 2: gotopt.ChallengeRequest
	(*ChallengeResponse)(nil), // 3: gotopt.ChallengeResponse
}
var file_totp_totp_proto_depIdxs = []int32{
	0, // 0: gotopt.Server.GetSecret:input_type -> gotopt.GetSecretRequest
	2, // 1: gotopt.Server.Challenge:input_type -> gotopt.ChallengeRequest
	1, // 2: gotopt.Server.GetSecret:output_type -> gotopt.GetSecretResponse
	3, // 3: gotopt.Server.Challenge:output_type -> gotopt.ChallengeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_totp_totp_proto_init() }
func file_totp_totp_proto_init() {
	if File_totp_totp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_totp_totp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretRequest); i {
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
		file_totp_totp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretResponse); i {
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
		file_totp_totp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeRequest); i {
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
		file_totp_totp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeResponse); i {
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
			RawDescriptor: file_totp_totp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_totp_totp_proto_goTypes,
		DependencyIndexes: file_totp_totp_proto_depIdxs,
		MessageInfos:      file_totp_totp_proto_msgTypes,
	}.Build()
	File_totp_totp_proto = out.File
	file_totp_totp_proto_rawDesc = nil
	file_totp_totp_proto_goTypes = nil
	file_totp_totp_proto_depIdxs = nil
}