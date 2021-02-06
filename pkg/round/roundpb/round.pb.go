// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/round/roundpb/round.proto

package roundpb

import (
	checkpb "github.com/ScoreTrak/ScoreTrak/pkg/check/checkpb"
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

type Round struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Start  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	Note   string                 `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	Err    string                 `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	Finish *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=finish,proto3" json:"finish,omitempty"`
	Checks []*checkpb.Check       `protobuf:"bytes,6,rep,name=checks,proto3" json:"checks,omitempty"`
}

func (x *Round) Reset() {
	*x = Round{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_round_roundpb_round_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Round) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Round) ProtoMessage() {}

func (x *Round) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_round_roundpb_round_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Round.ProtoReflect.Descriptor instead.
func (*Round) Descriptor() ([]byte, []int) {
	return file_pkg_round_roundpb_round_proto_rawDescGZIP(), []int{0}
}

func (x *Round) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Round) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Round) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Round) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *Round) GetFinish() *timestamppb.Timestamp {
	if x != nil {
		return x.Finish
	}
	return nil
}

func (x *Round) GetChecks() []*checkpb.Check {
	if x != nil {
		return x.Checks
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_round_roundpb_round_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_round_roundpb_round_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_pkg_round_roundpb_round_proto_rawDescGZIP(), []int{1}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rounds []*Round `protobuf:"bytes,1,rep,name=rounds,proto3" json:"rounds,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_round_roundpb_round_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_round_roundpb_round_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_pkg_round_roundpb_round_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllResponse) GetRounds() []*Round {
	if x != nil {
		return x.Rounds
	}
	return nil
}

type GetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIDRequest) Reset() {
	*x = GetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_round_roundpb_round_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDRequest) ProtoMessage() {}

func (x *GetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_round_roundpb_round_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDRequest.ProtoReflect.Descriptor instead.
func (*GetByIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_round_roundpb_round_proto_rawDescGZIP(), []int{3}
}

func (x *GetByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round *Round `protobuf:"bytes,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *GetByIDResponse) Reset() {
	*x = GetByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_round_roundpb_round_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDResponse) ProtoMessage() {}

func (x *GetByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_round_roundpb_round_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDResponse.ProtoReflect.Descriptor instead.
func (*GetByIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_round_roundpb_round_proto_rawDescGZIP(), []int{4}
}

func (x *GetByIDResponse) GetRound() *Round {
	if x != nil {
		return x.Round
	}
	return nil
}

type GetLastRoundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLastRoundRequest) Reset() {
	*x = GetLastRoundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_round_roundpb_round_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastRoundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastRoundRequest) ProtoMessage() {}

func (x *GetLastRoundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_round_roundpb_round_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastRoundRequest.ProtoReflect.Descriptor instead.
func (*GetLastRoundRequest) Descriptor() ([]byte, []int) {
	return file_pkg_round_roundpb_round_proto_rawDescGZIP(), []int{5}
}

type GetLastRoundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round *Round `protobuf:"bytes,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *GetLastRoundResponse) Reset() {
	*x = GetLastRoundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_round_roundpb_round_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastRoundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastRoundResponse) ProtoMessage() {}

func (x *GetLastRoundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_round_roundpb_round_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastRoundResponse.ProtoReflect.Descriptor instead.
func (*GetLastRoundResponse) Descriptor() ([]byte, []int) {
	return file_pkg_round_roundpb_round_proto_rawDescGZIP(), []int{6}
}

func (x *GetLastRoundResponse) GetRound() *Round {
	if x != nil {
		return x.Round
	}
	return nil
}

type GetLastNonElapsingRoundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLastNonElapsingRoundRequest) Reset() {
	*x = GetLastNonElapsingRoundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_round_roundpb_round_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastNonElapsingRoundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastNonElapsingRoundRequest) ProtoMessage() {}

func (x *GetLastNonElapsingRoundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_round_roundpb_round_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastNonElapsingRoundRequest.ProtoReflect.Descriptor instead.
func (*GetLastNonElapsingRoundRequest) Descriptor() ([]byte, []int) {
	return file_pkg_round_roundpb_round_proto_rawDescGZIP(), []int{7}
}

type GetLastNonElapsingRoundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round *Round `protobuf:"bytes,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *GetLastNonElapsingRoundResponse) Reset() {
	*x = GetLastNonElapsingRoundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_round_roundpb_round_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastNonElapsingRoundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastNonElapsingRoundResponse) ProtoMessage() {}

func (x *GetLastNonElapsingRoundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_round_roundpb_round_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastNonElapsingRoundResponse.ProtoReflect.Descriptor instead.
func (*GetLastNonElapsingRoundResponse) Descriptor() ([]byte, []int) {
	return file_pkg_round_roundpb_round_proto_rawDescGZIP(), []int{8}
}

func (x *GetLastNonElapsingRoundResponse) GetRound() *Round {
	if x != nil {
		return x.Round
	}
	return nil
}

var File_pkg_round_roundpb_round_proto protoreflect.FileDescriptor

var file_pkg_round_roundpb_round_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x70, 0x62, 0x2f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x70, 0x62, 0x1a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f,
	0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x06, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70,
	0x62, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x22,
	0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x05, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x05, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x22, 0x20, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x6f,
	0x6e, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74,
	0x4e, 0x6f, 0x6e, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x32, 0x9b, 0x03, 0x0a, 0x0c, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x6f, 0x6e, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x69, 0x6e, 0x67,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x31, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73,
	0x74, 0x4e, 0x6f, 0x6e, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x6f, 0x6e, 0x45, 0x6c, 0x61, 0x70, 0x73, 0x69, 0x6e, 0x67, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x12, 0x26, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x72, 0x61, 0x6b, 0x2f, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x54, 0x72, 0x61, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x2f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_round_roundpb_round_proto_rawDescOnce sync.Once
	file_pkg_round_roundpb_round_proto_rawDescData = file_pkg_round_roundpb_round_proto_rawDesc
)

func file_pkg_round_roundpb_round_proto_rawDescGZIP() []byte {
	file_pkg_round_roundpb_round_proto_rawDescOnce.Do(func() {
		file_pkg_round_roundpb_round_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_round_roundpb_round_proto_rawDescData)
	})
	return file_pkg_round_roundpb_round_proto_rawDescData
}

var file_pkg_round_roundpb_round_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_round_roundpb_round_proto_goTypes = []interface{}{
	(*Round)(nil),                           // 0: pkg.round.roundpb.Round
	(*GetAllRequest)(nil),                   // 1: pkg.round.roundpb.GetAllRequest
	(*GetAllResponse)(nil),                  // 2: pkg.round.roundpb.GetAllResponse
	(*GetByIDRequest)(nil),                  // 3: pkg.round.roundpb.GetByIDRequest
	(*GetByIDResponse)(nil),                 // 4: pkg.round.roundpb.GetByIDResponse
	(*GetLastRoundRequest)(nil),             // 5: pkg.round.roundpb.GetLastRoundRequest
	(*GetLastRoundResponse)(nil),            // 6: pkg.round.roundpb.GetLastRoundResponse
	(*GetLastNonElapsingRoundRequest)(nil),  // 7: pkg.round.roundpb.GetLastNonElapsingRoundRequest
	(*GetLastNonElapsingRoundResponse)(nil), // 8: pkg.round.roundpb.GetLastNonElapsingRoundResponse
	(*timestamppb.Timestamp)(nil),           // 9: google.protobuf.Timestamp
	(*checkpb.Check)(nil),                   // 10: pkg.check.checkpb.Check
}
var file_pkg_round_roundpb_round_proto_depIdxs = []int32{
	9,  // 0: pkg.round.roundpb.Round.start:type_name -> google.protobuf.Timestamp
	9,  // 1: pkg.round.roundpb.Round.finish:type_name -> google.protobuf.Timestamp
	10, // 2: pkg.round.roundpb.Round.checks:type_name -> pkg.check.checkpb.Check
	0,  // 3: pkg.round.roundpb.GetAllResponse.rounds:type_name -> pkg.round.roundpb.Round
	0,  // 4: pkg.round.roundpb.GetByIDResponse.round:type_name -> pkg.round.roundpb.Round
	0,  // 5: pkg.round.roundpb.GetLastRoundResponse.round:type_name -> pkg.round.roundpb.Round
	0,  // 6: pkg.round.roundpb.GetLastNonElapsingRoundResponse.round:type_name -> pkg.round.roundpb.Round
	7,  // 7: pkg.round.roundpb.RoundService.GetLastNonElapsingRound:input_type -> pkg.round.roundpb.GetLastNonElapsingRoundRequest
	1,  // 8: pkg.round.roundpb.RoundService.GetAll:input_type -> pkg.round.roundpb.GetAllRequest
	3,  // 9: pkg.round.roundpb.RoundService.GetByID:input_type -> pkg.round.roundpb.GetByIDRequest
	5,  // 10: pkg.round.roundpb.RoundService.GetLastRound:input_type -> pkg.round.roundpb.GetLastRoundRequest
	8,  // 11: pkg.round.roundpb.RoundService.GetLastNonElapsingRound:output_type -> pkg.round.roundpb.GetLastNonElapsingRoundResponse
	2,  // 12: pkg.round.roundpb.RoundService.GetAll:output_type -> pkg.round.roundpb.GetAllResponse
	4,  // 13: pkg.round.roundpb.RoundService.GetByID:output_type -> pkg.round.roundpb.GetByIDResponse
	6,  // 14: pkg.round.roundpb.RoundService.GetLastRound:output_type -> pkg.round.roundpb.GetLastRoundResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_round_roundpb_round_proto_init() }
func file_pkg_round_roundpb_round_proto_init() {
	if File_pkg_round_roundpb_round_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_round_roundpb_round_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Round); i {
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
		file_pkg_round_roundpb_round_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_pkg_round_roundpb_round_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_pkg_round_roundpb_round_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDRequest); i {
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
		file_pkg_round_roundpb_round_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDResponse); i {
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
		file_pkg_round_roundpb_round_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastRoundRequest); i {
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
		file_pkg_round_roundpb_round_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastRoundResponse); i {
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
		file_pkg_round_roundpb_round_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastNonElapsingRoundRequest); i {
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
		file_pkg_round_roundpb_round_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastNonElapsingRoundResponse); i {
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
			RawDescriptor: file_pkg_round_roundpb_round_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_round_roundpb_round_proto_goTypes,
		DependencyIndexes: file_pkg_round_roundpb_round_proto_depIdxs,
		MessageInfos:      file_pkg_round_roundpb_round_proto_msgTypes,
	}.Build()
	File_pkg_round_roundpb_round_proto = out.File
	file_pkg_round_roundpb_round_proto_rawDesc = nil
	file_pkg_round_roundpb_round_proto_goTypes = nil
	file_pkg_round_roundpb_round_proto_depIdxs = nil
}
