// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/notification/v1/notification.proto

package notificationv1

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

type NotifyOrderCreatedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *NotifyOrderCreatedRequest) Reset() {
	*x = NotifyOrderCreatedRequest{}
	mi := &file_api_notification_v1_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyOrderCreatedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderCreatedRequest) ProtoMessage() {}

func (x *NotifyOrderCreatedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_v1_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderCreatedRequest.ProtoReflect.Descriptor instead.
func (*NotifyOrderCreatedRequest) Descriptor() ([]byte, []int) {
	return file_api_notification_v1_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyOrderCreatedRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *NotifyOrderCreatedRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type NotifyOrderCreatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyOrderCreatedResponse) Reset() {
	*x = NotifyOrderCreatedResponse{}
	mi := &file_api_notification_v1_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyOrderCreatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderCreatedResponse) ProtoMessage() {}

func (x *NotifyOrderCreatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_v1_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderCreatedResponse.ProtoReflect.Descriptor instead.
func (*NotifyOrderCreatedResponse) Descriptor() ([]byte, []int) {
	return file_api_notification_v1_notification_proto_rawDescGZIP(), []int{1}
}

type NotifyOrderCanceledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *NotifyOrderCanceledRequest) Reset() {
	*x = NotifyOrderCanceledRequest{}
	mi := &file_api_notification_v1_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyOrderCanceledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderCanceledRequest) ProtoMessage() {}

func (x *NotifyOrderCanceledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_v1_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderCanceledRequest.ProtoReflect.Descriptor instead.
func (*NotifyOrderCanceledRequest) Descriptor() ([]byte, []int) {
	return file_api_notification_v1_notification_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyOrderCanceledRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *NotifyOrderCanceledRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type NotifyOrderCanceledResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyOrderCanceledResponse) Reset() {
	*x = NotifyOrderCanceledResponse{}
	mi := &file_api_notification_v1_notification_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyOrderCanceledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderCanceledResponse) ProtoMessage() {}

func (x *NotifyOrderCanceledResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_v1_notification_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderCanceledResponse.ProtoReflect.Descriptor instead.
func (*NotifyOrderCanceledResponse) Descriptor() ([]byte, []int) {
	return file_api_notification_v1_notification_proto_rawDescGZIP(), []int{3}
}

type NotifyOrderReadyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *NotifyOrderReadyRequest) Reset() {
	*x = NotifyOrderReadyRequest{}
	mi := &file_api_notification_v1_notification_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyOrderReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderReadyRequest) ProtoMessage() {}

func (x *NotifyOrderReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_v1_notification_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderReadyRequest.ProtoReflect.Descriptor instead.
func (*NotifyOrderReadyRequest) Descriptor() ([]byte, []int) {
	return file_api_notification_v1_notification_proto_rawDescGZIP(), []int{4}
}

func (x *NotifyOrderReadyRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *NotifyOrderReadyRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type NotifyOrderReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyOrderReadyResponse) Reset() {
	*x = NotifyOrderReadyResponse{}
	mi := &file_api_notification_v1_notification_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyOrderReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOrderReadyResponse) ProtoMessage() {}

func (x *NotifyOrderReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_notification_v1_notification_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOrderReadyResponse.ProtoReflect.Descriptor instead.
func (*NotifyOrderReadyResponse) Descriptor() ([]byte, []int) {
	return file_api_notification_v1_notification_proto_rawDescGZIP(), []int{5}
}

var File_api_notification_v1_notification_proto protoreflect.FileDescriptor

var file_api_notification_v1_notification_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x57, 0x0a,
	0x19, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x1a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1d,
	0x0a, 0x1b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x0a,
	0x17, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xfe, 0x02, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x12, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x7a, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71,
	0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xfb, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x68, 0x65, 0x6e, 0x67, 0x78, 0x75, 0x66, 0x65, 0x6e, 0x67, 0x31, 0x39, 0x39, 0x34, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x2d, 0x61, 0x72, 0x63, 0x68,
	0x2d, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x4e, 0x58, 0xaa, 0x02, 0x13, 0x41, 0x70, 0x69, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x13, 0x41, 0x70, 0x69, 0x5c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x41, 0x70, 0x69, 0x5c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_notification_v1_notification_proto_rawDescOnce sync.Once
	file_api_notification_v1_notification_proto_rawDescData = file_api_notification_v1_notification_proto_rawDesc
)

func file_api_notification_v1_notification_proto_rawDescGZIP() []byte {
	file_api_notification_v1_notification_proto_rawDescOnce.Do(func() {
		file_api_notification_v1_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_notification_v1_notification_proto_rawDescData)
	})
	return file_api_notification_v1_notification_proto_rawDescData
}

var file_api_notification_v1_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_notification_v1_notification_proto_goTypes = []any{
	(*NotifyOrderCreatedRequest)(nil),   // 0: api.notification.v1.NotifyOrderCreatedRequest
	(*NotifyOrderCreatedResponse)(nil),  // 1: api.notification.v1.NotifyOrderCreatedResponse
	(*NotifyOrderCanceledRequest)(nil),  // 2: api.notification.v1.NotifyOrderCanceledRequest
	(*NotifyOrderCanceledResponse)(nil), // 3: api.notification.v1.NotifyOrderCanceledResponse
	(*NotifyOrderReadyRequest)(nil),     // 4: api.notification.v1.NotifyOrderReadyRequest
	(*NotifyOrderReadyResponse)(nil),    // 5: api.notification.v1.NotifyOrderReadyResponse
}
var file_api_notification_v1_notification_proto_depIdxs = []int32{
	0, // 0: api.notification.v1.NotificationsService.NotifyOrderCreated:input_type -> api.notification.v1.NotifyOrderCreatedRequest
	2, // 1: api.notification.v1.NotificationsService.NotifyOrderCanceled:input_type -> api.notification.v1.NotifyOrderCanceledRequest
	4, // 2: api.notification.v1.NotificationsService.NotifyOrderReady:input_type -> api.notification.v1.NotifyOrderReadyRequest
	1, // 3: api.notification.v1.NotificationsService.NotifyOrderCreated:output_type -> api.notification.v1.NotifyOrderCreatedResponse
	3, // 4: api.notification.v1.NotificationsService.NotifyOrderCanceled:output_type -> api.notification.v1.NotifyOrderCanceledResponse
	5, // 5: api.notification.v1.NotificationsService.NotifyOrderReady:output_type -> api.notification.v1.NotifyOrderReadyResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_notification_v1_notification_proto_init() }
func file_api_notification_v1_notification_proto_init() {
	if File_api_notification_v1_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_notification_v1_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_notification_v1_notification_proto_goTypes,
		DependencyIndexes: file_api_notification_v1_notification_proto_depIdxs,
		MessageInfos:      file_api_notification_v1_notification_proto_msgTypes,
	}.Build()
	File_api_notification_v1_notification_proto = out.File
	file_api_notification_v1_notification_proto_rawDesc = nil
	file_api_notification_v1_notification_proto_goTypes = nil
	file_api_notification_v1_notification_proto_depIdxs = nil
}
