// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: saia/v1/saia.proto

package saiav1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Read request messages
type ReadInputRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       uint32                 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadInputRequest) Reset() {
	*x = ReadInputRequest{}
	mi := &file_saia_v1_saia_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadInputRequest) ProtoMessage() {}

func (x *ReadInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadInputRequest.ProtoReflect.Descriptor instead.
func (*ReadInputRequest) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{0}
}

func (x *ReadInputRequest) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

type ReadOutputRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       uint32                 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadOutputRequest) Reset() {
	*x = ReadOutputRequest{}
	mi := &file_saia_v1_saia_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadOutputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOutputRequest) ProtoMessage() {}

func (x *ReadOutputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOutputRequest.ProtoReflect.Descriptor instead.
func (*ReadOutputRequest) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{1}
}

func (x *ReadOutputRequest) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

type ReadFlagRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       uint32                 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadFlagRequest) Reset() {
	*x = ReadFlagRequest{}
	mi := &file_saia_v1_saia_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadFlagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFlagRequest) ProtoMessage() {}

func (x *ReadFlagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFlagRequest.ProtoReflect.Descriptor instead.
func (*ReadFlagRequest) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{2}
}

func (x *ReadFlagRequest) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

type ReadCounterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       uint32                 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadCounterRequest) Reset() {
	*x = ReadCounterRequest{}
	mi := &file_saia_v1_saia_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadCounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCounterRequest) ProtoMessage() {}

func (x *ReadCounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCounterRequest.ProtoReflect.Descriptor instead.
func (*ReadCounterRequest) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{3}
}

func (x *ReadCounterRequest) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

type ReadTimerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       uint32                 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadTimerRequest) Reset() {
	*x = ReadTimerRequest{}
	mi := &file_saia_v1_saia_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadTimerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadTimerRequest) ProtoMessage() {}

func (x *ReadTimerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadTimerRequest.ProtoReflect.Descriptor instead.
func (*ReadTimerRequest) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{4}
}

func (x *ReadTimerRequest) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

type ReadRegisterRequest struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Address uint32                 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are valid to be assigned to DataType:
	//
	//	*ReadRegisterRequest_AsInt
	//	*ReadRegisterRequest_AsFloat
	DataType      isReadRegisterRequest_DataType `protobuf_oneof:"data_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadRegisterRequest) Reset() {
	*x = ReadRegisterRequest{}
	mi := &file_saia_v1_saia_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRegisterRequest) ProtoMessage() {}

func (x *ReadRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRegisterRequest.ProtoReflect.Descriptor instead.
func (*ReadRegisterRequest) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{5}
}

func (x *ReadRegisterRequest) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *ReadRegisterRequest) GetDataType() isReadRegisterRequest_DataType {
	if x != nil {
		return x.DataType
	}
	return nil
}

func (x *ReadRegisterRequest) GetAsInt() bool {
	if x != nil {
		if x, ok := x.DataType.(*ReadRegisterRequest_AsInt); ok {
			return x.AsInt
		}
	}
	return false
}

func (x *ReadRegisterRequest) GetAsFloat() bool {
	if x != nil {
		if x, ok := x.DataType.(*ReadRegisterRequest_AsFloat); ok {
			return x.AsFloat
		}
	}
	return false
}

type isReadRegisterRequest_DataType interface {
	isReadRegisterRequest_DataType()
}

type ReadRegisterRequest_AsInt struct {
	AsInt bool `protobuf:"varint,2,opt,name=as_int,json=asInt,proto3,oneof"` // Set to true to read as integer
}

type ReadRegisterRequest_AsFloat struct {
	AsFloat bool `protobuf:"varint,3,opt,name=as_float,json=asFloat,proto3,oneof"` // Set to true to read as float
}

func (*ReadRegisterRequest_AsInt) isReadRegisterRequest_DataType() {}

func (*ReadRegisterRequest_AsFloat) isReadRegisterRequest_DataType() {}

// Read response messages
type ReadInputResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         bool                   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadInputResponse) Reset() {
	*x = ReadInputResponse{}
	mi := &file_saia_v1_saia_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadInputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadInputResponse) ProtoMessage() {}

func (x *ReadInputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadInputResponse.ProtoReflect.Descriptor instead.
func (*ReadInputResponse) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{6}
}

func (x *ReadInputResponse) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type ReadOutputResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         bool                   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadOutputResponse) Reset() {
	*x = ReadOutputResponse{}
	mi := &file_saia_v1_saia_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadOutputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadOutputResponse) ProtoMessage() {}

func (x *ReadOutputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadOutputResponse.ProtoReflect.Descriptor instead.
func (*ReadOutputResponse) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{7}
}

func (x *ReadOutputResponse) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type ReadFlagResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         bool                   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadFlagResponse) Reset() {
	*x = ReadFlagResponse{}
	mi := &file_saia_v1_saia_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadFlagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFlagResponse) ProtoMessage() {}

func (x *ReadFlagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFlagResponse.ProtoReflect.Descriptor instead.
func (*ReadFlagResponse) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{8}
}

func (x *ReadFlagResponse) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type ReadCounterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         int32                  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadCounterResponse) Reset() {
	*x = ReadCounterResponse{}
	mi := &file_saia_v1_saia_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadCounterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCounterResponse) ProtoMessage() {}

func (x *ReadCounterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCounterResponse.ProtoReflect.Descriptor instead.
func (*ReadCounterResponse) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{9}
}

func (x *ReadCounterResponse) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ReadTimerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         int32                  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadTimerResponse) Reset() {
	*x = ReadTimerResponse{}
	mi := &file_saia_v1_saia_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadTimerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadTimerResponse) ProtoMessage() {}

func (x *ReadTimerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadTimerResponse.ProtoReflect.Descriptor instead.
func (*ReadTimerResponse) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{10}
}

func (x *ReadTimerResponse) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ReadRegisterResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*ReadRegisterResponse_IntValue
	//	*ReadRegisterResponse_FloatValue
	Value         isReadRegisterResponse_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadRegisterResponse) Reset() {
	*x = ReadRegisterResponse{}
	mi := &file_saia_v1_saia_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRegisterResponse) ProtoMessage() {}

func (x *ReadRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRegisterResponse.ProtoReflect.Descriptor instead.
func (*ReadRegisterResponse) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{11}
}

func (x *ReadRegisterResponse) GetValue() isReadRegisterResponse_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ReadRegisterResponse) GetIntValue() int32 {
	if x != nil {
		if x, ok := x.Value.(*ReadRegisterResponse_IntValue); ok {
			return x.IntValue
		}
	}
	return 0
}

func (x *ReadRegisterResponse) GetFloatValue() float32 {
	if x != nil {
		if x, ok := x.Value.(*ReadRegisterResponse_FloatValue); ok {
			return x.FloatValue
		}
	}
	return 0
}

type isReadRegisterResponse_Value interface {
	isReadRegisterResponse_Value()
}

type ReadRegisterResponse_IntValue struct {
	IntValue int32 `protobuf:"varint,1,opt,name=int_value,json=intValue,proto3,oneof"`
}

type ReadRegisterResponse_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,2,opt,name=float_value,json=floatValue,proto3,oneof"`
}

func (*ReadRegisterResponse_IntValue) isReadRegisterResponse_Value() {}

func (*ReadRegisterResponse_FloatValue) isReadRegisterResponse_Value() {}

// Write operations
type WriteFlagRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       uint32                 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	Value         bool                   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteFlagRequest) Reset() {
	*x = WriteFlagRequest{}
	mi := &file_saia_v1_saia_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteFlagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFlagRequest) ProtoMessage() {}

func (x *WriteFlagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFlagRequest.ProtoReflect.Descriptor instead.
func (*WriteFlagRequest) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{12}
}

func (x *WriteFlagRequest) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *WriteFlagRequest) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type WriteFlagResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteFlagResponse) Reset() {
	*x = WriteFlagResponse{}
	mi := &file_saia_v1_saia_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteFlagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFlagResponse) ProtoMessage() {}

func (x *WriteFlagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFlagResponse.ProtoReflect.Descriptor instead.
func (*WriteFlagResponse) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{13}
}

type WriteRegisterRequest struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Address uint32                 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are valid to be assigned to Value:
	//
	//	*WriteRegisterRequest_IntValue
	//	*WriteRegisterRequest_FloatValue
	Value         isWriteRegisterRequest_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteRegisterRequest) Reset() {
	*x = WriteRegisterRequest{}
	mi := &file_saia_v1_saia_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRegisterRequest) ProtoMessage() {}

func (x *WriteRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRegisterRequest.ProtoReflect.Descriptor instead.
func (*WriteRegisterRequest) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{14}
}

func (x *WriteRegisterRequest) GetAddress() uint32 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *WriteRegisterRequest) GetValue() isWriteRegisterRequest_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *WriteRegisterRequest) GetIntValue() int32 {
	if x != nil {
		if x, ok := x.Value.(*WriteRegisterRequest_IntValue); ok {
			return x.IntValue
		}
	}
	return 0
}

func (x *WriteRegisterRequest) GetFloatValue() float32 {
	if x != nil {
		if x, ok := x.Value.(*WriteRegisterRequest_FloatValue); ok {
			return x.FloatValue
		}
	}
	return 0
}

type isWriteRegisterRequest_Value interface {
	isWriteRegisterRequest_Value()
}

type WriteRegisterRequest_IntValue struct {
	IntValue int32 `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type WriteRegisterRequest_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,3,opt,name=float_value,json=floatValue,proto3,oneof"`
}

func (*WriteRegisterRequest_IntValue) isWriteRegisterRequest_Value() {}

func (*WriteRegisterRequest_FloatValue) isWriteRegisterRequest_Value() {}

type WriteRegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteRegisterResponse) Reset() {
	*x = WriteRegisterResponse{}
	mi := &file_saia_v1_saia_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRegisterResponse) ProtoMessage() {}

func (x *WriteRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saia_v1_saia_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRegisterResponse.ProtoReflect.Descriptor instead.
func (*WriteRegisterResponse) Descriptor() ([]byte, []int) {
	return file_saia_v1_saia_proto_rawDescGZIP(), []int{15}
}

var File_saia_v1_saia_proto protoreflect.FileDescriptor

const file_saia_v1_saia_proto_rawDesc = "" +
	"\n" +
	"\x12saia/v1/saia.proto\x12\asaia.v1\",\n" +
	"\x10ReadInputRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\rR\aaddress\"-\n" +
	"\x11ReadOutputRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\rR\aaddress\"+\n" +
	"\x0fReadFlagRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\rR\aaddress\".\n" +
	"\x12ReadCounterRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\rR\aaddress\",\n" +
	"\x10ReadTimerRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\rR\aaddress\"r\n" +
	"\x13ReadRegisterRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\rR\aaddress\x12\x17\n" +
	"\x06as_int\x18\x02 \x01(\bH\x00R\x05asInt\x12\x1b\n" +
	"\bas_float\x18\x03 \x01(\bH\x00R\aasFloatB\v\n" +
	"\tdata_type\")\n" +
	"\x11ReadInputResponse\x12\x14\n" +
	"\x05value\x18\x01 \x01(\bR\x05value\"*\n" +
	"\x12ReadOutputResponse\x12\x14\n" +
	"\x05value\x18\x01 \x01(\bR\x05value\"(\n" +
	"\x10ReadFlagResponse\x12\x14\n" +
	"\x05value\x18\x01 \x01(\bR\x05value\"+\n" +
	"\x13ReadCounterResponse\x12\x14\n" +
	"\x05value\x18\x01 \x01(\x05R\x05value\")\n" +
	"\x11ReadTimerResponse\x12\x14\n" +
	"\x05value\x18\x01 \x01(\x05R\x05value\"a\n" +
	"\x14ReadRegisterResponse\x12\x1d\n" +
	"\tint_value\x18\x01 \x01(\x05H\x00R\bintValue\x12!\n" +
	"\vfloat_value\x18\x02 \x01(\x02H\x00R\n" +
	"floatValueB\a\n" +
	"\x05value\"B\n" +
	"\x10WriteFlagRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\rR\aaddress\x12\x14\n" +
	"\x05value\x18\x02 \x01(\bR\x05value\"\x13\n" +
	"\x11WriteFlagResponse\"{\n" +
	"\x14WriteRegisterRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\rR\aaddress\x12\x1d\n" +
	"\tint_value\x18\x02 \x01(\x05H\x00R\bintValue\x12!\n" +
	"\vfloat_value\x18\x03 \x01(\x02H\x00R\n" +
	"floatValueB\a\n" +
	"\x05value\"\x17\n" +
	"\x15WriteRegisterResponse2\xcb\x04\n" +
	"\x0eSaiaPcdService\x12B\n" +
	"\tReadInput\x12\x19.saia.v1.ReadInputRequest\x1a\x1a.saia.v1.ReadInputResponse\x12E\n" +
	"\n" +
	"ReadOutput\x12\x1a.saia.v1.ReadOutputRequest\x1a\x1b.saia.v1.ReadOutputResponse\x12?\n" +
	"\bReadFlag\x12\x18.saia.v1.ReadFlagRequest\x1a\x19.saia.v1.ReadFlagResponse\x12H\n" +
	"\vReadCounter\x12\x1b.saia.v1.ReadCounterRequest\x1a\x1c.saia.v1.ReadCounterResponse\x12B\n" +
	"\tReadTimer\x12\x19.saia.v1.ReadTimerRequest\x1a\x1a.saia.v1.ReadTimerResponse\x12K\n" +
	"\fReadRegister\x12\x1c.saia.v1.ReadRegisterRequest\x1a\x1d.saia.v1.ReadRegisterResponse\x12B\n" +
	"\tWriteFlag\x12\x19.saia.v1.WriteFlagRequest\x1a\x1a.saia.v1.WriteFlagResponse\x12N\n" +
	"\rWriteRegister\x12\x1d.saia.v1.WriteRegisterRequest\x1a\x1e.saia.v1.WriteRegisterResponseB\x95\x01\n" +
	"\vcom.saia.v1B\tSaiaProtoP\x01Z>github.com/clementnuss/saia-grpc-service/gen/go/saia/v1;saiav1\xa2\x02\x03SXX\xaa\x02\aSaia.V1\xca\x02\aSaia\\V1\xe2\x02\x13Saia\\V1\\GPBMetadata\xea\x02\bSaia::V1b\x06proto3"

var (
	file_saia_v1_saia_proto_rawDescOnce sync.Once
	file_saia_v1_saia_proto_rawDescData []byte
)

func file_saia_v1_saia_proto_rawDescGZIP() []byte {
	file_saia_v1_saia_proto_rawDescOnce.Do(func() {
		file_saia_v1_saia_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_saia_v1_saia_proto_rawDesc), len(file_saia_v1_saia_proto_rawDesc)))
	})
	return file_saia_v1_saia_proto_rawDescData
}

var file_saia_v1_saia_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_saia_v1_saia_proto_goTypes = []any{
	(*ReadInputRequest)(nil),      // 0: saia.v1.ReadInputRequest
	(*ReadOutputRequest)(nil),     // 1: saia.v1.ReadOutputRequest
	(*ReadFlagRequest)(nil),       // 2: saia.v1.ReadFlagRequest
	(*ReadCounterRequest)(nil),    // 3: saia.v1.ReadCounterRequest
	(*ReadTimerRequest)(nil),      // 4: saia.v1.ReadTimerRequest
	(*ReadRegisterRequest)(nil),   // 5: saia.v1.ReadRegisterRequest
	(*ReadInputResponse)(nil),     // 6: saia.v1.ReadInputResponse
	(*ReadOutputResponse)(nil),    // 7: saia.v1.ReadOutputResponse
	(*ReadFlagResponse)(nil),      // 8: saia.v1.ReadFlagResponse
	(*ReadCounterResponse)(nil),   // 9: saia.v1.ReadCounterResponse
	(*ReadTimerResponse)(nil),     // 10: saia.v1.ReadTimerResponse
	(*ReadRegisterResponse)(nil),  // 11: saia.v1.ReadRegisterResponse
	(*WriteFlagRequest)(nil),      // 12: saia.v1.WriteFlagRequest
	(*WriteFlagResponse)(nil),     // 13: saia.v1.WriteFlagResponse
	(*WriteRegisterRequest)(nil),  // 14: saia.v1.WriteRegisterRequest
	(*WriteRegisterResponse)(nil), // 15: saia.v1.WriteRegisterResponse
}
var file_saia_v1_saia_proto_depIdxs = []int32{
	0,  // 0: saia.v1.SaiaPcdService.ReadInput:input_type -> saia.v1.ReadInputRequest
	1,  // 1: saia.v1.SaiaPcdService.ReadOutput:input_type -> saia.v1.ReadOutputRequest
	2,  // 2: saia.v1.SaiaPcdService.ReadFlag:input_type -> saia.v1.ReadFlagRequest
	3,  // 3: saia.v1.SaiaPcdService.ReadCounter:input_type -> saia.v1.ReadCounterRequest
	4,  // 4: saia.v1.SaiaPcdService.ReadTimer:input_type -> saia.v1.ReadTimerRequest
	5,  // 5: saia.v1.SaiaPcdService.ReadRegister:input_type -> saia.v1.ReadRegisterRequest
	12, // 6: saia.v1.SaiaPcdService.WriteFlag:input_type -> saia.v1.WriteFlagRequest
	14, // 7: saia.v1.SaiaPcdService.WriteRegister:input_type -> saia.v1.WriteRegisterRequest
	6,  // 8: saia.v1.SaiaPcdService.ReadInput:output_type -> saia.v1.ReadInputResponse
	7,  // 9: saia.v1.SaiaPcdService.ReadOutput:output_type -> saia.v1.ReadOutputResponse
	8,  // 10: saia.v1.SaiaPcdService.ReadFlag:output_type -> saia.v1.ReadFlagResponse
	9,  // 11: saia.v1.SaiaPcdService.ReadCounter:output_type -> saia.v1.ReadCounterResponse
	10, // 12: saia.v1.SaiaPcdService.ReadTimer:output_type -> saia.v1.ReadTimerResponse
	11, // 13: saia.v1.SaiaPcdService.ReadRegister:output_type -> saia.v1.ReadRegisterResponse
	13, // 14: saia.v1.SaiaPcdService.WriteFlag:output_type -> saia.v1.WriteFlagResponse
	15, // 15: saia.v1.SaiaPcdService.WriteRegister:output_type -> saia.v1.WriteRegisterResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_saia_v1_saia_proto_init() }
func file_saia_v1_saia_proto_init() {
	if File_saia_v1_saia_proto != nil {
		return
	}
	file_saia_v1_saia_proto_msgTypes[5].OneofWrappers = []any{
		(*ReadRegisterRequest_AsInt)(nil),
		(*ReadRegisterRequest_AsFloat)(nil),
	}
	file_saia_v1_saia_proto_msgTypes[11].OneofWrappers = []any{
		(*ReadRegisterResponse_IntValue)(nil),
		(*ReadRegisterResponse_FloatValue)(nil),
	}
	file_saia_v1_saia_proto_msgTypes[14].OneofWrappers = []any{
		(*WriteRegisterRequest_IntValue)(nil),
		(*WriteRegisterRequest_FloatValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_saia_v1_saia_proto_rawDesc), len(file_saia_v1_saia_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_saia_v1_saia_proto_goTypes,
		DependencyIndexes: file_saia_v1_saia_proto_depIdxs,
		MessageInfos:      file_saia_v1_saia_proto_msgTypes,
	}.Build()
	File_saia_v1_saia_proto = out.File
	file_saia_v1_saia_proto_goTypes = nil
	file_saia_v1_saia_proto_depIdxs = nil
}
