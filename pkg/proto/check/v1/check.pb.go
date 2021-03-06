// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pkg/proto/check/v1/check.proto

package v1

import (
	v1 "github.com/ScoreTrak/ScoreTrak/pkg/proto/proto/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Check struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId *v1.UUID              `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	RoundId   uint64                `protobuf:"varint,2,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
	Log       string                `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Err       string                `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	Passed    *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=passed,proto3" json:"passed,omitempty"`
}

func (x *Check) Reset() {
	*x = Check{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_check_v1_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Check) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Check) ProtoMessage() {}

func (x *Check) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_check_v1_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Check.ProtoReflect.Descriptor instead.
func (*Check) Descriptor() ([]byte, []int) {
	return file_pkg_proto_check_v1_check_proto_rawDescGZIP(), []int{0}
}

func (x *Check) GetServiceId() *v1.UUID {
	if x != nil {
		return x.ServiceId
	}
	return nil
}

func (x *Check) GetRoundId() uint64 {
	if x != nil {
		return x.RoundId
	}
	return 0
}

func (x *Check) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *Check) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *Check) GetPassed() *wrapperspb.BoolValue {
	if x != nil {
		return x.Passed
	}
	return nil
}

type GetAllByRoundIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundId uint64 `protobuf:"varint,1,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
}

func (x *GetAllByRoundIDRequest) Reset() {
	*x = GetAllByRoundIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_check_v1_check_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByRoundIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByRoundIDRequest) ProtoMessage() {}

func (x *GetAllByRoundIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_check_v1_check_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByRoundIDRequest.ProtoReflect.Descriptor instead.
func (*GetAllByRoundIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_check_v1_check_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllByRoundIDRequest) GetRoundId() uint64 {
	if x != nil {
		return x.RoundId
	}
	return 0
}

type GetAllByRoundIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Checks []*Check `protobuf:"bytes,1,rep,name=checks,proto3" json:"checks,omitempty"`
}

func (x *GetAllByRoundIDResponse) Reset() {
	*x = GetAllByRoundIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_check_v1_check_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByRoundIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByRoundIDResponse) ProtoMessage() {}

func (x *GetAllByRoundIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_check_v1_check_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByRoundIDResponse.ProtoReflect.Descriptor instead.
func (*GetAllByRoundIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_check_v1_check_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllByRoundIDResponse) GetChecks() []*Check {
	if x != nil {
		return x.Checks
	}
	return nil
}

type GetByRoundServiceIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId *v1.UUID `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	RoundId   uint64   `protobuf:"varint,2,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
}

func (x *GetByRoundServiceIDRequest) Reset() {
	*x = GetByRoundServiceIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_check_v1_check_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByRoundServiceIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByRoundServiceIDRequest) ProtoMessage() {}

func (x *GetByRoundServiceIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_check_v1_check_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByRoundServiceIDRequest.ProtoReflect.Descriptor instead.
func (*GetByRoundServiceIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_check_v1_check_proto_rawDescGZIP(), []int{3}
}

func (x *GetByRoundServiceIDRequest) GetServiceId() *v1.UUID {
	if x != nil {
		return x.ServiceId
	}
	return nil
}

func (x *GetByRoundServiceIDRequest) GetRoundId() uint64 {
	if x != nil {
		return x.RoundId
	}
	return 0
}

type GetByRoundServiceIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Check *Check `protobuf:"bytes,1,opt,name=check,proto3" json:"check,omitempty"`
}

func (x *GetByRoundServiceIDResponse) Reset() {
	*x = GetByRoundServiceIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_check_v1_check_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByRoundServiceIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByRoundServiceIDResponse) ProtoMessage() {}

func (x *GetByRoundServiceIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_check_v1_check_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByRoundServiceIDResponse.ProtoReflect.Descriptor instead.
func (*GetByRoundServiceIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_check_v1_check_proto_rawDescGZIP(), []int{4}
}

func (x *GetByRoundServiceIDResponse) GetCheck() *Check {
	if x != nil {
		return x.Check
	}
	return nil
}

type GetAllByServiceIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId *v1.UUID `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
}

func (x *GetAllByServiceIDRequest) Reset() {
	*x = GetAllByServiceIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_check_v1_check_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByServiceIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByServiceIDRequest) ProtoMessage() {}

func (x *GetAllByServiceIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_check_v1_check_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByServiceIDRequest.ProtoReflect.Descriptor instead.
func (*GetAllByServiceIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_check_v1_check_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllByServiceIDRequest) GetServiceId() *v1.UUID {
	if x != nil {
		return x.ServiceId
	}
	return nil
}

type GetAllByServiceIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Checks []*Check `protobuf:"bytes,1,rep,name=checks,proto3" json:"checks,omitempty"`
}

func (x *GetAllByServiceIDResponse) Reset() {
	*x = GetAllByServiceIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_check_v1_check_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllByServiceIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllByServiceIDResponse) ProtoMessage() {}

func (x *GetAllByServiceIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_check_v1_check_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllByServiceIDResponse.ProtoReflect.Descriptor instead.
func (*GetAllByServiceIDResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_check_v1_check_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllByServiceIDResponse) GetChecks() []*Check {
	if x != nil {
		return x.Checks
	}
	return nil
}

var File_pkg_proto_check_v1_check_proto protoreflect.FileDescriptor

var file_pkg_proto_check_v1_check_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x37, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6c, 0x6f, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x22, 0x33, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x4c,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x22, 0x70, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x4e,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x53,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x06, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x32, 0xea, 0x02, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x2a, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42,
	0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x78, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x2e, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x2c, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x54, 0x72, 0x61, 0x6b, 0x2f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x72,
	0x61, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_check_v1_check_proto_rawDescOnce sync.Once
	file_pkg_proto_check_v1_check_proto_rawDescData = file_pkg_proto_check_v1_check_proto_rawDesc
)

func file_pkg_proto_check_v1_check_proto_rawDescGZIP() []byte {
	file_pkg_proto_check_v1_check_proto_rawDescOnce.Do(func() {
		file_pkg_proto_check_v1_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_check_v1_check_proto_rawDescData)
	})
	return file_pkg_proto_check_v1_check_proto_rawDescData
}

var file_pkg_proto_check_v1_check_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_check_v1_check_proto_goTypes = []interface{}{
	(*Check)(nil),                       // 0: pkg.proto.check.v1.Check
	(*GetAllByRoundIDRequest)(nil),      // 1: pkg.proto.check.v1.GetAllByRoundIDRequest
	(*GetAllByRoundIDResponse)(nil),     // 2: pkg.proto.check.v1.GetAllByRoundIDResponse
	(*GetByRoundServiceIDRequest)(nil),  // 3: pkg.proto.check.v1.GetByRoundServiceIDRequest
	(*GetByRoundServiceIDResponse)(nil), // 4: pkg.proto.check.v1.GetByRoundServiceIDResponse
	(*GetAllByServiceIDRequest)(nil),    // 5: pkg.proto.check.v1.GetAllByServiceIDRequest
	(*GetAllByServiceIDResponse)(nil),   // 6: pkg.proto.check.v1.GetAllByServiceIDResponse
	(*v1.UUID)(nil),                     // 7: pkg.proto.proto.v1.UUID
	(*wrapperspb.BoolValue)(nil),        // 8: google.protobuf.BoolValue
}
var file_pkg_proto_check_v1_check_proto_depIdxs = []int32{
	7,  // 0: pkg.proto.check.v1.Check.service_id:type_name -> pkg.proto.proto.v1.UUID
	8,  // 1: pkg.proto.check.v1.Check.passed:type_name -> google.protobuf.BoolValue
	0,  // 2: pkg.proto.check.v1.GetAllByRoundIDResponse.checks:type_name -> pkg.proto.check.v1.Check
	7,  // 3: pkg.proto.check.v1.GetByRoundServiceIDRequest.service_id:type_name -> pkg.proto.proto.v1.UUID
	0,  // 4: pkg.proto.check.v1.GetByRoundServiceIDResponse.check:type_name -> pkg.proto.check.v1.Check
	7,  // 5: pkg.proto.check.v1.GetAllByServiceIDRequest.service_id:type_name -> pkg.proto.proto.v1.UUID
	0,  // 6: pkg.proto.check.v1.GetAllByServiceIDResponse.checks:type_name -> pkg.proto.check.v1.Check
	1,  // 7: pkg.proto.check.v1.CheckService.GetAllByRoundID:input_type -> pkg.proto.check.v1.GetAllByRoundIDRequest
	3,  // 8: pkg.proto.check.v1.CheckService.GetByRoundServiceID:input_type -> pkg.proto.check.v1.GetByRoundServiceIDRequest
	5,  // 9: pkg.proto.check.v1.CheckService.GetAllByServiceID:input_type -> pkg.proto.check.v1.GetAllByServiceIDRequest
	2,  // 10: pkg.proto.check.v1.CheckService.GetAllByRoundID:output_type -> pkg.proto.check.v1.GetAllByRoundIDResponse
	4,  // 11: pkg.proto.check.v1.CheckService.GetByRoundServiceID:output_type -> pkg.proto.check.v1.GetByRoundServiceIDResponse
	6,  // 12: pkg.proto.check.v1.CheckService.GetAllByServiceID:output_type -> pkg.proto.check.v1.GetAllByServiceIDResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_proto_check_v1_check_proto_init() }
func file_pkg_proto_check_v1_check_proto_init() {
	if File_pkg_proto_check_v1_check_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_check_v1_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Check); i {
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
		file_pkg_proto_check_v1_check_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByRoundIDRequest); i {
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
		file_pkg_proto_check_v1_check_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByRoundIDResponse); i {
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
		file_pkg_proto_check_v1_check_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByRoundServiceIDRequest); i {
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
		file_pkg_proto_check_v1_check_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByRoundServiceIDResponse); i {
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
		file_pkg_proto_check_v1_check_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByServiceIDRequest); i {
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
		file_pkg_proto_check_v1_check_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllByServiceIDResponse); i {
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
			RawDescriptor: file_pkg_proto_check_v1_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_check_v1_check_proto_goTypes,
		DependencyIndexes: file_pkg_proto_check_v1_check_proto_depIdxs,
		MessageInfos:      file_pkg_proto_check_v1_check_proto_msgTypes,
	}.Build()
	File_pkg_proto_check_v1_check_proto = out.File
	file_pkg_proto_check_v1_check_proto_rawDesc = nil
	file_pkg_proto_check_v1_check_proto_goTypes = nil
	file_pkg_proto_check_v1_check_proto_depIdxs = nil
}
