// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api_action.proto

package api

import (
	pagination "github.com/iotexproject/iotex-analyser-api/api/pagination"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ActHash    string                 `protobuf:"bytes,2,opt,name=actHash,proto3" json:"actHash,omitempty"`
	Pagination *pagination.Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_api_action_proto_rawDescGZIP(), []int{0}
}

func (x *ActionRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ActionRequest) GetActHash() string {
	if x != nil {
		return x.ActHash
	}
	return ""
}

func (x *ActionRequest) GetPagination() *pagination.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist           bool               `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	Count           uint64             `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	ActionList      []*ActionInfo      `protobuf:"bytes,3,rep,name=actionList,proto3" json:"actionList,omitempty"`
	EvmTransferList []*EvmTransferInfo `protobuf:"bytes,4,rep,name=evmTransferList,proto3" json:"evmTransferList,omitempty"`
	XrcList         []*XrcInfo         `protobuf:"bytes,5,rep,name=xrcList,proto3" json:"xrcList,omitempty"`
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_api_action_proto_rawDescGZIP(), []int{1}
}

func (x *ActionResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *ActionResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ActionResponse) GetActionList() []*ActionInfo {
	if x != nil {
		return x.ActionList
	}
	return nil
}

func (x *ActionResponse) GetEvmTransferList() []*EvmTransferInfo {
	if x != nil {
		return x.EvmTransferList
	}
	return nil
}

func (x *ActionResponse) GetXrcList() []*XrcInfo {
	if x != nil {
		return x.XrcList
	}
	return nil
}

type ActionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActHash   string `protobuf:"bytes,1,opt,name=actHash,proto3" json:"actHash,omitempty"`
	BlkHash   string `protobuf:"bytes,2,opt,name=blkHash,proto3" json:"blkHash,omitempty"`
	ActType   string `protobuf:"bytes,3,opt,name=actType,proto3" json:"actType,omitempty"`
	Sender    string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient string `protobuf:"bytes,5,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount    string `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp uint64 `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GasFee    string `protobuf:"bytes,8,opt,name=gasFee,proto3" json:"gasFee,omitempty"`
	BlkHeight uint64 `protobuf:"varint,9,opt,name=blkHeight,proto3" json:"blkHeight,omitempty"`
}

func (x *ActionInfo) Reset() {
	*x = ActionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionInfo) ProtoMessage() {}

func (x *ActionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionInfo.ProtoReflect.Descriptor instead.
func (*ActionInfo) Descriptor() ([]byte, []int) {
	return file_api_action_proto_rawDescGZIP(), []int{2}
}

func (x *ActionInfo) GetActHash() string {
	if x != nil {
		return x.ActHash
	}
	return ""
}

func (x *ActionInfo) GetBlkHash() string {
	if x != nil {
		return x.BlkHash
	}
	return ""
}

func (x *ActionInfo) GetActType() string {
	if x != nil {
		return x.ActType
	}
	return ""
}

func (x *ActionInfo) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ActionInfo) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *ActionInfo) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *ActionInfo) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ActionInfo) GetGasFee() string {
	if x != nil {
		return x.GasFee
	}
	return ""
}

func (x *ActionInfo) GetBlkHeight() uint64 {
	if x != nil {
		return x.BlkHeight
	}
	return 0
}

type EvmTransferInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActHash   string `protobuf:"bytes,1,opt,name=actHash,proto3" json:"actHash,omitempty"`
	BlkHash   string `protobuf:"bytes,2,opt,name=blkHash,proto3" json:"blkHash,omitempty"`
	From      string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Quantity  string `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	BlkHeight uint64 `protobuf:"varint,6,opt,name=blkHeight,proto3" json:"blkHeight,omitempty"`
	Timestamp uint64 `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *EvmTransferInfo) Reset() {
	*x = EvmTransferInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_action_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvmTransferInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvmTransferInfo) ProtoMessage() {}

func (x *EvmTransferInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_action_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvmTransferInfo.ProtoReflect.Descriptor instead.
func (*EvmTransferInfo) Descriptor() ([]byte, []int) {
	return file_api_action_proto_rawDescGZIP(), []int{3}
}

func (x *EvmTransferInfo) GetActHash() string {
	if x != nil {
		return x.ActHash
	}
	return ""
}

func (x *EvmTransferInfo) GetBlkHash() string {
	if x != nil {
		return x.BlkHash
	}
	return ""
}

func (x *EvmTransferInfo) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *EvmTransferInfo) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *EvmTransferInfo) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *EvmTransferInfo) GetBlkHeight() uint64 {
	if x != nil {
		return x.BlkHeight
	}
	return 0
}

func (x *EvmTransferInfo) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type XrcInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActHash   string `protobuf:"bytes,1,opt,name=actHash,proto3" json:"actHash,omitempty"`
	From      string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Quantity  string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	BlkHeight uint64 `protobuf:"varint,5,opt,name=blkHeight,proto3" json:"blkHeight,omitempty"`
	Timestamp uint64 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Contract  string `protobuf:"bytes,7,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *XrcInfo) Reset() {
	*x = XrcInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_action_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XrcInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XrcInfo) ProtoMessage() {}

func (x *XrcInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_action_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XrcInfo.ProtoReflect.Descriptor instead.
func (*XrcInfo) Descriptor() ([]byte, []int) {
	return file_api_action_proto_rawDescGZIP(), []int{4}
}

func (x *XrcInfo) GetActHash() string {
	if x != nil {
		return x.ActHash
	}
	return ""
}

func (x *XrcInfo) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *XrcInfo) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *XrcInfo) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *XrcInfo) GetBlkHeight() uint64 {
	if x != nil {
		return x.BlkHeight
	}
	return 0
}

func (x *XrcInfo) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *XrcInfo) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

var File_api_action_proto protoreflect.FileDescriptor

var file_api_action_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x36, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xd5, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0f, 0x65, 0x76, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x76, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0f, 0x65, 0x76, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x78, 0x72, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x58, 0x72, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x78, 0x72, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xfc, 0x01, 0x0a, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x73, 0x46, 0x65, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x73, 0x46, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x6c, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x62, 0x6c, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xc1, 0x01, 0x0a, 0x0f, 0x45,
	0x76, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x6c, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xbb,
	0x01, 0x0a, 0x07, 0x58, 0x72, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x6c, 0x6b, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x32, 0xb3, 0x03, 0x0a,
	0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80,
	0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x56, 0x6f,
	0x74, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0xba, 0x43,
	0x12, 0x12, 0x10, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x56, 0x6f,
	0x74, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x3a, 0x01,
	0x2a, 0x12, 0x98, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x76, 0x6d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x73, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0xba, 0x43, 0x1a, 0x12, 0x18, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x42, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x22, 0x2b, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73,
	0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x83, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x58, 0x72, 0x63, 0x32, 0x30, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0xba, 0x43, 0x13,
	0x12, 0x11, 0x47, 0x65, 0x74, 0x58, 0x72, 0x63, 0x32, 0x30, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x58, 0x72, 0x63, 0x32, 0x30, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a,
	0x01, 0x2a, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_action_proto_rawDescOnce sync.Once
	file_api_action_proto_rawDescData = file_api_action_proto_rawDesc
)

func file_api_action_proto_rawDescGZIP() []byte {
	file_api_action_proto_rawDescOnce.Do(func() {
		file_api_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_action_proto_rawDescData)
	})
	return file_api_action_proto_rawDescData
}

var file_api_action_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_action_proto_goTypes = []interface{}{
	(*ActionRequest)(nil),         // 0: api.ActionRequest
	(*ActionResponse)(nil),        // 1: api.ActionResponse
	(*ActionInfo)(nil),            // 2: api.ActionInfo
	(*EvmTransferInfo)(nil),       // 3: api.EvmTransferInfo
	(*XrcInfo)(nil),               // 4: api.XrcInfo
	(*pagination.Pagination)(nil), // 5: pagination.Pagination
}
var file_api_action_proto_depIdxs = []int32{
	5, // 0: api.ActionRequest.pagination:type_name -> pagination.Pagination
	2, // 1: api.ActionResponse.actionList:type_name -> api.ActionInfo
	3, // 2: api.ActionResponse.evmTransferList:type_name -> api.EvmTransferInfo
	4, // 3: api.ActionResponse.xrcList:type_name -> api.XrcInfo
	0, // 4: api.ActionService.GetActionByVoter:input_type -> api.ActionRequest
	0, // 5: api.ActionService.GetEvmTransfersByAddress:input_type -> api.ActionRequest
	0, // 6: api.ActionService.GetXrc20ByAddress:input_type -> api.ActionRequest
	1, // 7: api.ActionService.GetActionByVoter:output_type -> api.ActionResponse
	1, // 8: api.ActionService.GetEvmTransfersByAddress:output_type -> api.ActionResponse
	1, // 9: api.ActionService.GetXrc20ByAddress:output_type -> api.ActionResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_action_proto_init() }
func file_api_action_proto_init() {
	if File_api_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRequest); i {
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
		file_api_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResponse); i {
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
		file_api_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionInfo); i {
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
		file_api_action_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvmTransferInfo); i {
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
		file_api_action_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XrcInfo); i {
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
			RawDescriptor: file_api_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_action_proto_goTypes,
		DependencyIndexes: file_api_action_proto_depIdxs,
		MessageInfos:      file_api_action_proto_msgTypes,
	}.Build()
	File_api_action_proto = out.File
	file_api_action_proto_rawDesc = nil
	file_api_action_proto_goTypes = nil
	file_api_action_proto_depIdxs = nil
}