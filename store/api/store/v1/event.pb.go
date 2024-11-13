// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/store/v1/event.proto

package storev1

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

type StoreCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *StoreCreated) Reset() {
	*x = StoreCreated{}
	mi := &file_api_store_v1_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCreated) ProtoMessage() {}

func (x *StoreCreated) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCreated.ProtoReflect.Descriptor instead.
func (*StoreCreated) Descriptor() ([]byte, []int) {
	return file_api_store_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *StoreCreated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoreCreated) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoreCreated) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type StoreParticipationToggled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Participating bool   `protobuf:"varint,2,opt,name=participating,proto3" json:"participating,omitempty"`
}

func (x *StoreParticipationToggled) Reset() {
	*x = StoreParticipationToggled{}
	mi := &file_api_store_v1_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreParticipationToggled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreParticipationToggled) ProtoMessage() {}

func (x *StoreParticipationToggled) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreParticipationToggled.ProtoReflect.Descriptor instead.
func (*StoreParticipationToggled) Descriptor() ([]byte, []int) {
	return file_api_store_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *StoreParticipationToggled) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoreParticipationToggled) GetParticipating() bool {
	if x != nil {
		return x.Participating
	}
	return false
}

type StoreRebranded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StoreRebranded) Reset() {
	*x = StoreRebranded{}
	mi := &file_api_store_v1_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreRebranded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRebranded) ProtoMessage() {}

func (x *StoreRebranded) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_event_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreRebranded.ProtoReflect.Descriptor instead.
func (*StoreRebranded) Descriptor() ([]byte, []int) {
	return file_api_store_v1_event_proto_rawDescGZIP(), []int{2}
}

func (x *StoreRebranded) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoreRebranded) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProductAdded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StoreId     string  `protobuf:"bytes,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Sku         string  `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
	Price       float64 `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ProductAdded) Reset() {
	*x = ProductAdded{}
	mi := &file_api_store_v1_event_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductAdded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductAdded) ProtoMessage() {}

func (x *ProductAdded) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_event_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductAdded.ProtoReflect.Descriptor instead.
func (*ProductAdded) Descriptor() ([]byte, []int) {
	return file_api_store_v1_event_proto_rawDescGZIP(), []int{3}
}

func (x *ProductAdded) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductAdded) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *ProductAdded) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductAdded) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductAdded) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *ProductAdded) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ProductRebranded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ProductRebranded) Reset() {
	*x = ProductRebranded{}
	mi := &file_api_store_v1_event_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductRebranded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRebranded) ProtoMessage() {}

func (x *ProductRebranded) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_event_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRebranded.ProtoReflect.Descriptor instead.
func (*ProductRebranded) Descriptor() ([]byte, []int) {
	return file_api_store_v1_event_proto_rawDescGZIP(), []int{4}
}

func (x *ProductRebranded) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductRebranded) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductRebranded) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ProductPriceChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Delta float64 `protobuf:"fixed64,2,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *ProductPriceChanged) Reset() {
	*x = ProductPriceChanged{}
	mi := &file_api_store_v1_event_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductPriceChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPriceChanged) ProtoMessage() {}

func (x *ProductPriceChanged) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_event_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPriceChanged.ProtoReflect.Descriptor instead.
func (*ProductPriceChanged) Descriptor() ([]byte, []int) {
	return file_api_store_v1_event_proto_rawDescGZIP(), []int{5}
}

func (x *ProductPriceChanged) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductPriceChanged) GetDelta() float64 {
	if x != nil {
		return x.Delta
	}
	return 0
}

type ProductRemoved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductRemoved) Reset() {
	*x = ProductRemoved{}
	mi := &file_api_store_v1_event_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductRemoved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRemoved) ProtoMessage() {}

func (x *ProductRemoved) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_event_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRemoved.ProtoReflect.Descriptor instead.
func (*ProductRemoved) Descriptor() ([]byte, []int) {
	return file_api_store_v1_event_proto_rawDescGZIP(), []int{6}
}

func (x *ProductRemoved) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_store_v1_event_proto protoreflect.FileDescriptor

var file_api_store_v1_event_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x4e, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x19, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x67, 0x67, 0x6c, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x34, 0x0a, 0x0e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x97, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x64, 0x64,
	0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x58, 0x0a, 0x10, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x64, 0x65, 0x6c,
	0x74, 0x61, 0x22, 0x20, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x42, 0xc3, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x65, 0x6e, 0x67, 0x78, 0x75, 0x66, 0x65, 0x6e, 0x67, 0x31,
	0x39, 0x39, 0x34, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6e,
	0x2d, 0x61, 0x72, 0x63, 0x68, 0x2d, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x58,
	0xaa, 0x02, 0x0c, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0c, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x18, 0x41, 0x70, 0x69, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_store_v1_event_proto_rawDescOnce sync.Once
	file_api_store_v1_event_proto_rawDescData = file_api_store_v1_event_proto_rawDesc
)

func file_api_store_v1_event_proto_rawDescGZIP() []byte {
	file_api_store_v1_event_proto_rawDescOnce.Do(func() {
		file_api_store_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_store_v1_event_proto_rawDescData)
	})
	return file_api_store_v1_event_proto_rawDescData
}

var file_api_store_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_store_v1_event_proto_goTypes = []any{
	(*StoreCreated)(nil),              // 0: api.store.v1.StoreCreated
	(*StoreParticipationToggled)(nil), // 1: api.store.v1.StoreParticipationToggled
	(*StoreRebranded)(nil),            // 2: api.store.v1.StoreRebranded
	(*ProductAdded)(nil),              // 3: api.store.v1.ProductAdded
	(*ProductRebranded)(nil),          // 4: api.store.v1.ProductRebranded
	(*ProductPriceChanged)(nil),       // 5: api.store.v1.ProductPriceChanged
	(*ProductRemoved)(nil),            // 6: api.store.v1.ProductRemoved
}
var file_api_store_v1_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_store_v1_event_proto_init() }
func file_api_store_v1_event_proto_init() {
	if File_api_store_v1_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_store_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_store_v1_event_proto_goTypes,
		DependencyIndexes: file_api_store_v1_event_proto_depIdxs,
		MessageInfos:      file_api_store_v1_event_proto_msgTypes,
	}.Build()
	File_api_store_v1_event_proto = out.File
	file_api_store_v1_event_proto_rawDesc = nil
	file_api_store_v1_event_proto_goTypes = nil
	file_api_store_v1_event_proto_depIdxs = nil
}
