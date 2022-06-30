// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api_voting.proto

package api

import (
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

type CandidateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartEpoch uint64 `protobuf:"varint,1,opt,name=startEpoch,proto3" json:"startEpoch,omitempty"` // starting epoch number
	EpochCount uint64 `protobuf:"varint,2,opt,name=epochCount,proto3" json:"epochCount,omitempty"` // epoch count
}

func (x *CandidateInfoRequest) Reset() {
	*x = CandidateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateInfoRequest) ProtoMessage() {}

func (x *CandidateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateInfoRequest.ProtoReflect.Descriptor instead.
func (*CandidateInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_voting_proto_rawDescGZIP(), []int{0}
}

func (x *CandidateInfoRequest) GetStartEpoch() uint64 {
	if x != nil {
		return x.StartEpoch
	}
	return 0
}

func (x *CandidateInfoRequest) GetEpochCount() uint64 {
	if x != nil {
		return x.EpochCount
	}
	return 0
}

type CandidateInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidateInfo []*CandidateInfoResponse_CandidateInfo `protobuf:"bytes,1,rep,name=candidateInfo,proto3" json:"candidateInfo,omitempty"`
}

func (x *CandidateInfoResponse) Reset() {
	*x = CandidateInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateInfoResponse) ProtoMessage() {}

func (x *CandidateInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_voting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateInfoResponse.ProtoReflect.Descriptor instead.
func (*CandidateInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_voting_proto_rawDescGZIP(), []int{1}
}

func (x *CandidateInfoResponse) GetCandidateInfo() []*CandidateInfoResponse_CandidateInfo {
	if x != nil {
		return x.CandidateInfo
	}
	return nil
}

type RewardSourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartEpoch       uint64 `protobuf:"varint,1,opt,name=startEpoch,proto3" json:"startEpoch,omitempty"`            // starting epoch number
	EpochCount       uint64 `protobuf:"varint,2,opt,name=epochCount,proto3" json:"epochCount,omitempty"`            // epoch count
	VoterIotxAddress string `protobuf:"bytes,3,opt,name=voterIotxAddress,proto3" json:"voterIotxAddress,omitempty"` // voter IoTeX address
}

func (x *RewardSourcesRequest) Reset() {
	*x = RewardSourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardSourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardSourcesRequest) ProtoMessage() {}

func (x *RewardSourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_voting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardSourcesRequest.ProtoReflect.Descriptor instead.
func (*RewardSourcesRequest) Descriptor() ([]byte, []int) {
	return file_api_voting_proto_rawDescGZIP(), []int{2}
}

func (x *RewardSourcesRequest) GetStartEpoch() uint64 {
	if x != nil {
		return x.StartEpoch
	}
	return 0
}

func (x *RewardSourcesRequest) GetEpochCount() uint64 {
	if x != nil {
		return x.EpochCount
	}
	return 0
}

func (x *RewardSourcesRequest) GetVoterIotxAddress() string {
	if x != nil {
		return x.VoterIotxAddress
	}
	return ""
}

type RewardSourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist                 bool                                           `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"` // whether the voter has reward information within the specified epoch range
	DelegateDistributions []*RewardSourcesResponse_DelegateDistributions `protobuf:"bytes,2,rep,name=delegateDistributions,proto3" json:"delegateDistributions,omitempty"`
}

func (x *RewardSourcesResponse) Reset() {
	*x = RewardSourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardSourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardSourcesResponse) ProtoMessage() {}

func (x *RewardSourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_voting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardSourcesResponse.ProtoReflect.Descriptor instead.
func (*RewardSourcesResponse) Descriptor() ([]byte, []int) {
	return file_api_voting_proto_rawDescGZIP(), []int{3}
}

func (x *RewardSourcesResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *RewardSourcesResponse) GetDelegateDistributions() []*RewardSourcesResponse_DelegateDistributions {
	if x != nil {
		return x.DelegateDistributions
	}
	return nil
}

type CandidateInfoResponse_Candidates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                             // candidate name
	Address            string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`                       // canddiate address
	TotalWeightedVotes string `protobuf:"bytes,3,opt,name=totalWeightedVotes,proto3" json:"totalWeightedVotes,omitempty"` // total weighted votes
	SelfStakingTokens  string `protobuf:"bytes,4,opt,name=selfStakingTokens,proto3" json:"selfStakingTokens,omitempty"`   // candidate self-staking tokens
	OperatorAddress    string `protobuf:"bytes,5,opt,name=operatorAddress,proto3" json:"operatorAddress,omitempty"`       // candidate operator address
	RewardAddress      string `protobuf:"bytes,6,opt,name=rewardAddress,proto3" json:"rewardAddress,omitempty"`           // candidate reward address
}

func (x *CandidateInfoResponse_Candidates) Reset() {
	*x = CandidateInfoResponse_Candidates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateInfoResponse_Candidates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateInfoResponse_Candidates) ProtoMessage() {}

func (x *CandidateInfoResponse_Candidates) ProtoReflect() protoreflect.Message {
	mi := &file_api_voting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateInfoResponse_Candidates.ProtoReflect.Descriptor instead.
func (*CandidateInfoResponse_Candidates) Descriptor() ([]byte, []int) {
	return file_api_voting_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CandidateInfoResponse_Candidates) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CandidateInfoResponse_Candidates) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CandidateInfoResponse_Candidates) GetTotalWeightedVotes() string {
	if x != nil {
		return x.TotalWeightedVotes
	}
	return ""
}

func (x *CandidateInfoResponse_Candidates) GetSelfStakingTokens() string {
	if x != nil {
		return x.SelfStakingTokens
	}
	return ""
}

func (x *CandidateInfoResponse_Candidates) GetOperatorAddress() string {
	if x != nil {
		return x.OperatorAddress
	}
	return ""
}

func (x *CandidateInfoResponse_Candidates) GetRewardAddress() string {
	if x != nil {
		return x.RewardAddress
	}
	return ""
}

type CandidateInfoResponse_CandidateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpochNumber uint64                              `protobuf:"varint,1,opt,name=epochNumber,proto3" json:"epochNumber,omitempty"` // epoch number
	Candidates  []*CandidateInfoResponse_Candidates `protobuf:"bytes,2,rep,name=candidates,proto3" json:"candidates,omitempty"`
}

func (x *CandidateInfoResponse_CandidateInfo) Reset() {
	*x = CandidateInfoResponse_CandidateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateInfoResponse_CandidateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateInfoResponse_CandidateInfo) ProtoMessage() {}

func (x *CandidateInfoResponse_CandidateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_voting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateInfoResponse_CandidateInfo.ProtoReflect.Descriptor instead.
func (*CandidateInfoResponse_CandidateInfo) Descriptor() ([]byte, []int) {
	return file_api_voting_proto_rawDescGZIP(), []int{1, 1}
}

func (x *CandidateInfoResponse_CandidateInfo) GetEpochNumber() uint64 {
	if x != nil {
		return x.EpochNumber
	}
	return 0
}

func (x *CandidateInfoResponse_CandidateInfo) GetCandidates() []*CandidateInfoResponse_Candidates {
	if x != nil {
		return x.Candidates
	}
	return nil
}

type RewardSourcesResponse_DelegateDistributions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegateName string `protobuf:"bytes,1,opt,name=delegateName,proto3" json:"delegateName,omitempty"` // delegate name
	Amount       string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`             // amount of reward distribution
}

func (x *RewardSourcesResponse_DelegateDistributions) Reset() {
	*x = RewardSourcesResponse_DelegateDistributions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_voting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardSourcesResponse_DelegateDistributions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardSourcesResponse_DelegateDistributions) ProtoMessage() {}

func (x *RewardSourcesResponse_DelegateDistributions) ProtoReflect() protoreflect.Message {
	mi := &file_api_voting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardSourcesResponse_DelegateDistributions.ProtoReflect.Descriptor instead.
func (*RewardSourcesResponse_DelegateDistributions) Descriptor() ([]byte, []int) {
	return file_api_voting_proto_rawDescGZIP(), []int{3, 0}
}

func (x *RewardSourcesResponse_DelegateDistributions) GetDelegateName() string {
	if x != nil {
		return x.DelegateName
	}
	return ""
}

func (x *RewardSourcesResponse_DelegateDistributions) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_api_voting_proto protoreflect.FileDescriptor

var file_api_voting_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x14,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xcc, 0x03, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0xe8,
	0x01, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x73,
	0x65, 0x6c, 0x66, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x6c, 0x66, 0x53, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x78, 0x0a, 0x0d, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0a,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x76, 0x6f, 0x74, 0x65, 0x72, 0x49, 0x6f, 0x74, 0x78, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x49, 0x6f, 0x74,
	0x78, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xea, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x66, 0x0a, 0x15, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x15, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x53, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x9f, 0x02, 0x0a, 0x0d, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3d, 0xba, 0x43, 0x0f, 0x12, 0x0d, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x3a, 0x01, 0x2a, 0x12,
	0x85, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0xba, 0x43, 0x0f, 0x12, 0x0d, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x25, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_voting_proto_rawDescOnce sync.Once
	file_api_voting_proto_rawDescData = file_api_voting_proto_rawDesc
)

func file_api_voting_proto_rawDescGZIP() []byte {
	file_api_voting_proto_rawDescOnce.Do(func() {
		file_api_voting_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_voting_proto_rawDescData)
	})
	return file_api_voting_proto_rawDescData
}

var file_api_voting_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_voting_proto_goTypes = []interface{}{
	(*CandidateInfoRequest)(nil),                        // 0: api.CandidateInfoRequest
	(*CandidateInfoResponse)(nil),                       // 1: api.CandidateInfoResponse
	(*RewardSourcesRequest)(nil),                        // 2: api.RewardSourcesRequest
	(*RewardSourcesResponse)(nil),                       // 3: api.RewardSourcesResponse
	(*CandidateInfoResponse_Candidates)(nil),            // 4: api.CandidateInfoResponse.Candidates
	(*CandidateInfoResponse_CandidateInfo)(nil),         // 5: api.CandidateInfoResponse.CandidateInfo
	(*RewardSourcesResponse_DelegateDistributions)(nil), // 6: api.RewardSourcesResponse.DelegateDistributions
}
var file_api_voting_proto_depIdxs = []int32{
	5, // 0: api.CandidateInfoResponse.candidateInfo:type_name -> api.CandidateInfoResponse.CandidateInfo
	6, // 1: api.RewardSourcesResponse.delegateDistributions:type_name -> api.RewardSourcesResponse.DelegateDistributions
	4, // 2: api.CandidateInfoResponse.CandidateInfo.candidates:type_name -> api.CandidateInfoResponse.Candidates
	0, // 3: api.VotingService.CandidateInfo:input_type -> api.CandidateInfoRequest
	2, // 4: api.VotingService.RewardSources:input_type -> api.RewardSourcesRequest
	1, // 5: api.VotingService.CandidateInfo:output_type -> api.CandidateInfoResponse
	3, // 6: api.VotingService.RewardSources:output_type -> api.RewardSourcesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_voting_proto_init() }
func file_api_voting_proto_init() {
	if File_api_voting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_voting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateInfoRequest); i {
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
		file_api_voting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateInfoResponse); i {
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
		file_api_voting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardSourcesRequest); i {
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
		file_api_voting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardSourcesResponse); i {
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
		file_api_voting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateInfoResponse_Candidates); i {
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
		file_api_voting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateInfoResponse_CandidateInfo); i {
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
		file_api_voting_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardSourcesResponse_DelegateDistributions); i {
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
			RawDescriptor: file_api_voting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_voting_proto_goTypes,
		DependencyIndexes: file_api_voting_proto_depIdxs,
		MessageInfos:      file_api_voting_proto_msgTypes,
	}.Build()
	File_api_voting_proto = out.File
	file_api_voting_proto_rawDesc = nil
	file_api_voting_proto_goTypes = nil
	file_api_voting_proto_depIdxs = nil
}
