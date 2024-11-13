// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/customer/v1/customer.proto

package customerv1

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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SmsNumber string `protobuf:"bytes,3,opt,name=sms_number,json=smsNumber,proto3" json:"sms_number,omitempty"`
	Enabled   bool   `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetSmsNumber() string {
	if x != nil {
		return x.SmsNumber
	}
	return ""
}

func (x *Customer) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type RegisterCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SmsNumber string `protobuf:"bytes,2,opt,name=sms_number,json=smsNumber,proto3" json:"sms_number,omitempty"`
}

func (x *RegisterCustomerRequest) Reset() {
	*x = RegisterCustomerRequest{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCustomerRequest) ProtoMessage() {}

func (x *RegisterCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCustomerRequest.ProtoReflect.Descriptor instead.
func (*RegisterCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterCustomerRequest) GetSmsNumber() string {
	if x != nil {
		return x.SmsNumber
	}
	return ""
}

type RegisterCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RegisterCustomerResponse) Reset() {
	*x = RegisterCustomerResponse{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCustomerResponse) ProtoMessage() {}

func (x *RegisterCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCustomerResponse.ProtoReflect.Descriptor instead.
func (*RegisterCustomerResponse) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterCustomerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EnableCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EnableCustomerRequest) Reset() {
	*x = EnableCustomerRequest{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnableCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableCustomerRequest) ProtoMessage() {}

func (x *EnableCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableCustomerRequest.ProtoReflect.Descriptor instead.
func (*EnableCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{3}
}

func (x *EnableCustomerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EnableCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnableCustomerResponse) Reset() {
	*x = EnableCustomerResponse{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnableCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableCustomerResponse) ProtoMessage() {}

func (x *EnableCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableCustomerResponse.ProtoReflect.Descriptor instead.
func (*EnableCustomerResponse) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{4}
}

type DisableCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DisableCustomerRequest) Reset() {
	*x = DisableCustomerRequest{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DisableCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableCustomerRequest) ProtoMessage() {}

func (x *DisableCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableCustomerRequest.ProtoReflect.Descriptor instead.
func (*DisableCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{5}
}

func (x *DisableCustomerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DisableCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisableCustomerResponse) Reset() {
	*x = DisableCustomerResponse{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DisableCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableCustomerResponse) ProtoMessage() {}

func (x *DisableCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableCustomerResponse.ProtoReflect.Descriptor instead.
func (*DisableCustomerResponse) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{6}
}

type ChangeSmsNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SmsNumber string `protobuf:"bytes,2,opt,name=sms_number,json=smsNumber,proto3" json:"sms_number,omitempty"`
}

func (x *ChangeSmsNumberRequest) Reset() {
	*x = ChangeSmsNumberRequest{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeSmsNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSmsNumberRequest) ProtoMessage() {}

func (x *ChangeSmsNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSmsNumberRequest.ProtoReflect.Descriptor instead.
func (*ChangeSmsNumberRequest) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{7}
}

func (x *ChangeSmsNumberRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangeSmsNumberRequest) GetSmsNumber() string {
	if x != nil {
		return x.SmsNumber
	}
	return ""
}

type ChangeSmsNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeSmsNumberResponse) Reset() {
	*x = ChangeSmsNumberResponse{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeSmsNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSmsNumberResponse) ProtoMessage() {}

func (x *ChangeSmsNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSmsNumberResponse.ProtoReflect.Descriptor instead.
func (*ChangeSmsNumberResponse) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{8}
}

type AuthorizeCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorizeCustomerRequest) Reset() {
	*x = AuthorizeCustomerRequest{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizeCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeCustomerRequest) ProtoMessage() {}

func (x *AuthorizeCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeCustomerRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{9}
}

func (x *AuthorizeCustomerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AuthorizeCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorizeCustomerResponse) Reset() {
	*x = AuthorizeCustomerResponse{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizeCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeCustomerResponse) ProtoMessage() {}

func (x *AuthorizeCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeCustomerResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeCustomerResponse) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{10}
}

type GetCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{11}
}

func (x *GetCustomerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *GetCustomerResponse) Reset() {
	*x = GetCustomerResponse{}
	mi := &file_api_customer_v1_customer_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerResponse) ProtoMessage() {}

func (x *GetCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_customer_v1_customer_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerResponse) Descriptor() ([]byte, []int) {
	return file_api_customer_v1_customer_proto_rawDescGZIP(), []int{12}
}

func (x *GetCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

var File_api_customer_v1_customer_proto protoreflect.FileDescriptor

var file_api_customer_v1_customer_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x22, 0x67, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x4c, 0x0a, 0x17, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x73,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a,
	0x16, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x0a, 0x16,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6d, 0x73, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6d, 0x73, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2a, 0x0a, 0x18, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x19,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x4c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x32, 0xfc, 0x04,
	0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x69, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a,
	0x0e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x66, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0f, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6c, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xde, 0x01, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x68, 0x65, 0x6e, 0x67, 0x78, 0x75, 0x66, 0x65, 0x6e, 0x67, 0x31, 0x39, 0x39,
	0x34, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x2d, 0x61,
	0x72, 0x63, 0x68, 0x2d, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x41, 0x43, 0x58, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x41, 0x70,
	0x69, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_customer_v1_customer_proto_rawDescOnce sync.Once
	file_api_customer_v1_customer_proto_rawDescData = file_api_customer_v1_customer_proto_rawDesc
)

func file_api_customer_v1_customer_proto_rawDescGZIP() []byte {
	file_api_customer_v1_customer_proto_rawDescOnce.Do(func() {
		file_api_customer_v1_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_customer_v1_customer_proto_rawDescData)
	})
	return file_api_customer_v1_customer_proto_rawDescData
}

var file_api_customer_v1_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_customer_v1_customer_proto_goTypes = []any{
	(*Customer)(nil),                  // 0: api.customer.v1.Customer
	(*RegisterCustomerRequest)(nil),   // 1: api.customer.v1.RegisterCustomerRequest
	(*RegisterCustomerResponse)(nil),  // 2: api.customer.v1.RegisterCustomerResponse
	(*EnableCustomerRequest)(nil),     // 3: api.customer.v1.EnableCustomerRequest
	(*EnableCustomerResponse)(nil),    // 4: api.customer.v1.EnableCustomerResponse
	(*DisableCustomerRequest)(nil),    // 5: api.customer.v1.DisableCustomerRequest
	(*DisableCustomerResponse)(nil),   // 6: api.customer.v1.DisableCustomerResponse
	(*ChangeSmsNumberRequest)(nil),    // 7: api.customer.v1.ChangeSmsNumberRequest
	(*ChangeSmsNumberResponse)(nil),   // 8: api.customer.v1.ChangeSmsNumberResponse
	(*AuthorizeCustomerRequest)(nil),  // 9: api.customer.v1.AuthorizeCustomerRequest
	(*AuthorizeCustomerResponse)(nil), // 10: api.customer.v1.AuthorizeCustomerResponse
	(*GetCustomerRequest)(nil),        // 11: api.customer.v1.GetCustomerRequest
	(*GetCustomerResponse)(nil),       // 12: api.customer.v1.GetCustomerResponse
}
var file_api_customer_v1_customer_proto_depIdxs = []int32{
	0,  // 0: api.customer.v1.GetCustomerResponse.customer:type_name -> api.customer.v1.Customer
	1,  // 1: api.customer.v1.CustomersService.RegisterCustomer:input_type -> api.customer.v1.RegisterCustomerRequest
	3,  // 2: api.customer.v1.CustomersService.EnableCustomer:input_type -> api.customer.v1.EnableCustomerRequest
	5,  // 3: api.customer.v1.CustomersService.DisableCustomer:input_type -> api.customer.v1.DisableCustomerRequest
	7,  // 4: api.customer.v1.CustomersService.ChangeSmsNumber:input_type -> api.customer.v1.ChangeSmsNumberRequest
	9,  // 5: api.customer.v1.CustomersService.AuthorizeCustomer:input_type -> api.customer.v1.AuthorizeCustomerRequest
	11, // 6: api.customer.v1.CustomersService.GetCustomer:input_type -> api.customer.v1.GetCustomerRequest
	2,  // 7: api.customer.v1.CustomersService.RegisterCustomer:output_type -> api.customer.v1.RegisterCustomerResponse
	4,  // 8: api.customer.v1.CustomersService.EnableCustomer:output_type -> api.customer.v1.EnableCustomerResponse
	6,  // 9: api.customer.v1.CustomersService.DisableCustomer:output_type -> api.customer.v1.DisableCustomerResponse
	8,  // 10: api.customer.v1.CustomersService.ChangeSmsNumber:output_type -> api.customer.v1.ChangeSmsNumberResponse
	10, // 11: api.customer.v1.CustomersService.AuthorizeCustomer:output_type -> api.customer.v1.AuthorizeCustomerResponse
	12, // 12: api.customer.v1.CustomersService.GetCustomer:output_type -> api.customer.v1.GetCustomerResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_customer_v1_customer_proto_init() }
func file_api_customer_v1_customer_proto_init() {
	if File_api_customer_v1_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_customer_v1_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_customer_v1_customer_proto_goTypes,
		DependencyIndexes: file_api_customer_v1_customer_proto_depIdxs,
		MessageInfos:      file_api_customer_v1_customer_proto_msgTypes,
	}.Build()
	File_api_customer_v1_customer_proto = out.File
	file_api_customer_v1_customer_proto_rawDesc = nil
	file_api_customer_v1_customer_proto_goTypes = nil
	file_api_customer_v1_customer_proto_depIdxs = nil
}
