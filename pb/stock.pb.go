// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: stock.proto

package pb

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SearchStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId string `protobuf:"bytes,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
}

func (x *SearchStock) Reset() {
	*x = SearchStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchStock) ProtoMessage() {}

func (x *SearchStock) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchStock.ProtoReflect.Descriptor instead.
func (*SearchStock) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{0}
}

func (x *SearchStock) GetGoodsId() string {
	if x != nil {
		return x.GoodsId
	}
	return ""
}

type StockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId  string `protobuf:"bytes,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	StockNum int32  `protobuf:"varint,2,opt,name=stock_num,json=stockNum,proto3" json:"stock_num,omitempty"`
}

func (x *StockInfo) Reset() {
	*x = StockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockInfo) ProtoMessage() {}

func (x *StockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockInfo.ProtoReflect.Descriptor instead.
func (*StockInfo) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{1}
}

func (x *StockInfo) GetGoodsId() string {
	if x != nil {
		return x.GoodsId
	}
	return ""
}

func (x *StockInfo) GetStockNum() int32 {
	if x != nil {
		return x.StockNum
	}
	return 0
}

type PreDeductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId string `protobuf:"bytes,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *PreDeductRequest) Reset() {
	*x = PreDeductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreDeductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreDeductRequest) ProtoMessage() {}

func (x *PreDeductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreDeductRequest.ProtoReflect.Descriptor instead.
func (*PreDeductRequest) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{2}
}

func (x *PreDeductRequest) GetGoodsId() string {
	if x != nil {
		return x.GoodsId
	}
	return ""
}

func (x *PreDeductRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type PreDeductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId string `protobuf:"bytes,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *PreDeductResponse) Reset() {
	*x = PreDeductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreDeductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreDeductResponse) ProtoMessage() {}

func (x *PreDeductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreDeductResponse.ProtoReflect.Descriptor instead.
func (*PreDeductResponse) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{3}
}

func (x *PreDeductResponse) GetGoodsId() string {
	if x != nil {
		return x.GoodsId
	}
	return ""
}

func (x *PreDeductResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CancelTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId string `protobuf:"bytes,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CancelTransactionRequest) Reset() {
	*x = CancelTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTransactionRequest) ProtoMessage() {}

func (x *CancelTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTransactionRequest.ProtoReflect.Descriptor instead.
func (*CancelTransactionRequest) Descriptor() ([]byte, []int) {
	return file_stock_proto_rawDescGZIP(), []int{4}
}

func (x *CancelTransactionRequest) GetGoodsId() string {
	if x != nil {
		return x.GoodsId
	}
	return ""
}

func (x *CancelTransactionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_stock_proto protoreflect.FileDescriptor

var file_stock_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x09,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x22, 0x46, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x11, 0x50, 0x72, 0x65,
	0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x4e, 0x0a, 0x18, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x32, 0xa1, 0x02, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x4a, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x10, 0x2e, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x59, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x44,
	0x65, 0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72,
	0x65, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x65, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x2f, 0x64, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x50, 0x72, 0x65, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stock_proto_rawDescOnce sync.Once
	file_stock_proto_rawDescData = file_stock_proto_rawDesc
)

func file_stock_proto_rawDescGZIP() []byte {
	file_stock_proto_rawDescOnce.Do(func() {
		file_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_proto_rawDescData)
	})
	return file_stock_proto_rawDescData
}

var file_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_stock_proto_goTypes = []interface{}{
	(*SearchStock)(nil),              // 0: stock.SearchStock
	(*StockInfo)(nil),                // 1: stock.StockInfo
	(*PreDeductRequest)(nil),         // 2: stock.PreDeductRequest
	(*PreDeductResponse)(nil),        // 3: stock.PreDeductResponse
	(*CancelTransactionRequest)(nil), // 4: stock.CancelTransactionRequest
}
var file_stock_proto_depIdxs = []int32{
	0, // 0: stock.stock.GetStock:input_type -> stock.SearchStock
	2, // 1: stock.stock.PreDeduct:input_type -> stock.PreDeductRequest
	4, // 2: stock.stock.CancelTransaction:input_type -> stock.CancelTransactionRequest
	1, // 3: stock.stock.GetStock:output_type -> stock.StockInfo
	3, // 4: stock.stock.PreDeduct:output_type -> stock.PreDeductResponse
	3, // 5: stock.stock.CancelTransaction:output_type -> stock.PreDeductResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stock_proto_init() }
func file_stock_proto_init() {
	if File_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchStock); i {
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
		file_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockInfo); i {
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
		file_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreDeductRequest); i {
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
		file_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreDeductResponse); i {
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
		file_stock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTransactionRequest); i {
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
			RawDescriptor: file_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_proto_goTypes,
		DependencyIndexes: file_stock_proto_depIdxs,
		MessageInfos:      file_stock_proto_msgTypes,
	}.Build()
	File_stock_proto = out.File
	file_stock_proto_rawDesc = nil
	file_stock_proto_goTypes = nil
	file_stock_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StockClient is the client API for Stock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockClient interface {
	GetStock(ctx context.Context, in *SearchStock, opts ...grpc.CallOption) (*StockInfo, error)
	PreDeduct(ctx context.Context, in *PreDeductRequest, opts ...grpc.CallOption) (*PreDeductResponse, error)
	CancelTransaction(ctx context.Context, in *CancelTransactionRequest, opts ...grpc.CallOption) (*PreDeductResponse, error)
}

type stockClient struct {
	cc grpc.ClientConnInterface
}

func NewStockClient(cc grpc.ClientConnInterface) StockClient {
	return &stockClient{cc}
}

func (c *stockClient) GetStock(ctx context.Context, in *SearchStock, opts ...grpc.CallOption) (*StockInfo, error) {
	out := new(StockInfo)
	err := c.cc.Invoke(ctx, "/stock.stock/GetStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) PreDeduct(ctx context.Context, in *PreDeductRequest, opts ...grpc.CallOption) (*PreDeductResponse, error) {
	out := new(PreDeductResponse)
	err := c.cc.Invoke(ctx, "/stock.stock/PreDeduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) CancelTransaction(ctx context.Context, in *CancelTransactionRequest, opts ...grpc.CallOption) (*PreDeductResponse, error) {
	out := new(PreDeductResponse)
	err := c.cc.Invoke(ctx, "/stock.stock/CancelTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServer is the server API for Stock service.
type StockServer interface {
	GetStock(context.Context, *SearchStock) (*StockInfo, error)
	PreDeduct(context.Context, *PreDeductRequest) (*PreDeductResponse, error)
	CancelTransaction(context.Context, *CancelTransactionRequest) (*PreDeductResponse, error)
}

// UnimplementedStockServer can be embedded to have forward compatible implementations.
type UnimplementedStockServer struct {
}

func (*UnimplementedStockServer) GetStock(context.Context, *SearchStock) (*StockInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}
func (*UnimplementedStockServer) PreDeduct(context.Context, *PreDeductRequest) (*PreDeductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreDeduct not implemented")
}
func (*UnimplementedStockServer) CancelTransaction(context.Context, *CancelTransactionRequest) (*PreDeductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTransaction not implemented")
}

func RegisterStockServer(s *grpc.Server, srv StockServer) {
	s.RegisterService(&_Stock_serviceDesc, srv)
}

func _Stock_GetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchStock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).GetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.stock/GetStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).GetStock(ctx, req.(*SearchStock))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_PreDeduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreDeductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).PreDeduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.stock/PreDeduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).PreDeduct(ctx, req.(*PreDeductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_CancelTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).CancelTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.stock/CancelTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).CancelTransaction(ctx, req.(*CancelTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stock.stock",
	HandlerType: (*StockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStock",
			Handler:    _Stock_GetStock_Handler,
		},
		{
			MethodName: "PreDeduct",
			Handler:    _Stock_PreDeduct_Handler,
		},
		{
			MethodName: "CancelTransaction",
			Handler:    _Stock_CancelTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock.proto",
}
