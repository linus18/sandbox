// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: posting.proto

package posting_api_grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PostingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostingDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=posting_date,json=postingDate,proto3" json:"posting_date,omitempty"`
	Merchant    string                 `protobuf:"bytes,2,opt,name=merchant,proto3" json:"merchant,omitempty"`
	Amount      int64                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	IsCredit    bool                   `protobuf:"varint,4,opt,name=is_credit,json=isCredit,proto3" json:"is_credit,omitempty"`
	AccountId   string                 `protobuf:"bytes,5,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *PostingRequest) Reset() {
	*x = PostingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostingRequest) ProtoMessage() {}

func (x *PostingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostingRequest.ProtoReflect.Descriptor instead.
func (*PostingRequest) Descriptor() ([]byte, []int) {
	return file_posting_proto_rawDescGZIP(), []int{0}
}

func (x *PostingRequest) GetPostingDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PostingDate
	}
	return nil
}

func (x *PostingRequest) GetMerchant() string {
	if x != nil {
		return x.Merchant
	}
	return ""
}

func (x *PostingRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PostingRequest) GetIsCredit() bool {
	if x != nil {
		return x.IsCredit
	}
	return false
}

func (x *PostingRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type PostingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseCode string `protobuf:"bytes,1,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
}

func (x *PostingReply) Reset() {
	*x = PostingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostingReply) ProtoMessage() {}

func (x *PostingReply) ProtoReflect() protoreflect.Message {
	mi := &file_posting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostingReply.ProtoReflect.Descriptor instead.
func (*PostingReply) Descriptor() ([]byte, []int) {
	return file_posting_proto_rawDescGZIP(), []int{1}
}

func (x *PostingReply) GetResponseCode() string {
	if x != nil {
		return x.ResponseCode
	}
	return ""
}

var File_posting_proto protoreflect.FileDescriptor

var file_posting_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x5e, 0x0a, 0x07, 0x50, 0x6f, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x53, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x73, 0x31, 0x38, 0x2f,
	0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_posting_proto_rawDescOnce sync.Once
	file_posting_proto_rawDescData = file_posting_proto_rawDesc
)

func file_posting_proto_rawDescGZIP() []byte {
	file_posting_proto_rawDescOnce.Do(func() {
		file_posting_proto_rawDescData = protoimpl.X.CompressGZIP(file_posting_proto_rawDescData)
	})
	return file_posting_proto_rawDescData
}

var file_posting_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_posting_proto_goTypes = []interface{}{
	(*PostingRequest)(nil),        // 0: posting_api_grpc.PostingRequest
	(*PostingReply)(nil),          // 1: posting_api_grpc.PostingReply
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_posting_proto_depIdxs = []int32{
	2, // 0: posting_api_grpc.PostingRequest.posting_date:type_name -> google.protobuf.Timestamp
	0, // 1: posting_api_grpc.Posting.CreatePosting:input_type -> posting_api_grpc.PostingRequest
	1, // 2: posting_api_grpc.Posting.CreatePosting:output_type -> posting_api_grpc.PostingReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_posting_proto_init() }
func file_posting_proto_init() {
	if File_posting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_posting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostingRequest); i {
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
		file_posting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostingReply); i {
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
			RawDescriptor: file_posting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_posting_proto_goTypes,
		DependencyIndexes: file_posting_proto_depIdxs,
		MessageInfos:      file_posting_proto_msgTypes,
	}.Build()
	File_posting_proto = out.File
	file_posting_proto_rawDesc = nil
	file_posting_proto_goTypes = nil
	file_posting_proto_depIdxs = nil
}