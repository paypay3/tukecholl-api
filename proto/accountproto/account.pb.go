// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/accountproto/account.proto

package accountproto

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

type CreateStandardBudgetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateStandardBudgetsRequest) Reset() {
	*x = CreateStandardBudgetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_accountproto_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStandardBudgetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStandardBudgetsRequest) ProtoMessage() {}

func (x *CreateStandardBudgetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_accountproto_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStandardBudgetsRequest.ProtoReflect.Descriptor instead.
func (*CreateStandardBudgetsRequest) Descriptor() ([]byte, []int) {
	return file_proto_accountproto_account_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStandardBudgetsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateStandardBudgetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateStandardBudgetsResponse) Reset() {
	*x = CreateStandardBudgetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_accountproto_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStandardBudgetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStandardBudgetsResponse) ProtoMessage() {}

func (x *CreateStandardBudgetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_accountproto_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStandardBudgetsResponse.ProtoReflect.Descriptor instead.
func (*CreateStandardBudgetsResponse) Descriptor() ([]byte, []int) {
	return file_proto_accountproto_account_proto_rawDescGZIP(), []int{1}
}

var File_proto_accountproto_account_proto protoreflect.FileDescriptor

var file_proto_accountproto_account_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x37, 0x0a, 0x1c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x78, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x12, 0x25, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61,
	0x79, 0x70, 0x61, 0x79, 0x33, 0x2f, 0x74, 0x75, 0x6b, 0x65, 0x63, 0x68, 0x6f, 0x6c, 0x6c, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_accountproto_account_proto_rawDescOnce sync.Once
	file_proto_accountproto_account_proto_rawDescData = file_proto_accountproto_account_proto_rawDesc
)

func file_proto_accountproto_account_proto_rawDescGZIP() []byte {
	file_proto_accountproto_account_proto_rawDescOnce.Do(func() {
		file_proto_accountproto_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_accountproto_account_proto_rawDescData)
	})
	return file_proto_accountproto_account_proto_rawDescData
}

var file_proto_accountproto_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_accountproto_account_proto_goTypes = []interface{}{
	(*CreateStandardBudgetsRequest)(nil),  // 0: account.CreateStandardBudgetsRequest
	(*CreateStandardBudgetsResponse)(nil), // 1: account.CreateStandardBudgetsResponse
}
var file_proto_accountproto_account_proto_depIdxs = []int32{
	0, // 0: account.AccountService.CreateStandardBudgets:input_type -> account.CreateStandardBudgetsRequest
	1, // 1: account.AccountService.CreateStandardBudgets:output_type -> account.CreateStandardBudgetsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_accountproto_account_proto_init() }
func file_proto_accountproto_account_proto_init() {
	if File_proto_accountproto_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_accountproto_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStandardBudgetsRequest); i {
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
		file_proto_accountproto_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStandardBudgetsResponse); i {
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
			RawDescriptor: file_proto_accountproto_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_accountproto_account_proto_goTypes,
		DependencyIndexes: file_proto_accountproto_account_proto_depIdxs,
		MessageInfos:      file_proto_accountproto_account_proto_msgTypes,
	}.Build()
	File_proto_accountproto_account_proto = out.File
	file_proto_accountproto_account_proto_rawDesc = nil
	file_proto_accountproto_account_proto_goTypes = nil
	file_proto_accountproto_account_proto_depIdxs = nil
}
