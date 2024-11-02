// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: api/basket/v1/basket.proto

package basketv1

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

type Basket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items []*Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Basket) Reset() {
	*x = Basket{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Basket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Basket) ProtoMessage() {}

func (x *Basket) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Basket.ProtoReflect.Descriptor instead.
func (*Basket) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{0}
}

func (x *Basket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Basket) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId      string  `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ProductId    string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	StoreName    string  `protobuf:"bytes,3,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	ProductName  string  `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductPrice float64 `protobuf:"fixed64,5,opt,name=product_price,json=productPrice,proto3" json:"product_price,omitempty"`
	Quantity     int32   `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Item) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *Item) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Item) GetProductPrice() float64 {
	if x != nil {
		return x.ProductPrice
	}
	return 0
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type StartBasketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *StartBasketRequest) Reset() {
	*x = StartBasketRequest{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartBasketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBasketRequest) ProtoMessage() {}

func (x *StartBasketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBasketRequest.ProtoReflect.Descriptor instead.
func (*StartBasketRequest) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{2}
}

func (x *StartBasketRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type StartBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartBasketResponse) Reset() {
	*x = StartBasketResponse{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBasketResponse) ProtoMessage() {}

func (x *StartBasketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBasketResponse.ProtoReflect.Descriptor instead.
func (*StartBasketResponse) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{3}
}

func (x *StartBasketResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelBasketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelBasketRequest) Reset() {
	*x = CancelBasketRequest{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelBasketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBasketRequest) ProtoMessage() {}

func (x *CancelBasketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBasketRequest.ProtoReflect.Descriptor instead.
func (*CancelBasketRequest) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{4}
}

func (x *CancelBasketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelBasketResponse) Reset() {
	*x = CancelBasketResponse{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBasketResponse) ProtoMessage() {}

func (x *CancelBasketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBasketResponse.ProtoReflect.Descriptor instead.
func (*CancelBasketResponse) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{5}
}

type CheckoutBasketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PaymentId string `protobuf:"bytes,2,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
}

func (x *CheckoutBasketRequest) Reset() {
	*x = CheckoutBasketRequest{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckoutBasketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutBasketRequest) ProtoMessage() {}

func (x *CheckoutBasketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutBasketRequest.ProtoReflect.Descriptor instead.
func (*CheckoutBasketRequest) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{6}
}

func (x *CheckoutBasketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CheckoutBasketRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type CheckoutBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckoutBasketResponse) Reset() {
	*x = CheckoutBasketResponse{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckoutBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutBasketResponse) ProtoMessage() {}

func (x *CheckoutBasketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutBasketResponse.ProtoReflect.Descriptor instead.
func (*CheckoutBasketResponse) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{7}
}

type AddItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AddItemRequest) Reset() {
	*x = AddItemRequest{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemRequest) ProtoMessage() {}

func (x *AddItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemRequest.ProtoReflect.Descriptor instead.
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{8}
}

func (x *AddItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddItemResponse) Reset() {
	*x = AddItemResponse{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResponse) ProtoMessage() {}

func (x *AddItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemResponse.ProtoReflect.Descriptor instead.
func (*AddItemResponse) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{9}
}

type RemoveItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *RemoveItemRequest) Reset() {
	*x = RemoveItemRequest{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemRequest) ProtoMessage() {}

func (x *RemoveItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemRequest.ProtoReflect.Descriptor instead.
func (*RemoveItemRequest) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoveItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *RemoveItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type RemoveItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveItemResponse) Reset() {
	*x = RemoveItemResponse{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemResponse) ProtoMessage() {}

func (x *RemoveItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemResponse.ProtoReflect.Descriptor instead.
func (*RemoveItemResponse) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{11}
}

type GetBasketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBasketRequest) Reset() {
	*x = GetBasketRequest{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBasketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBasketRequest) ProtoMessage() {}

func (x *GetBasketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBasketRequest.ProtoReflect.Descriptor instead.
func (*GetBasketRequest) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{12}
}

func (x *GetBasketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Basket *Basket `protobuf:"bytes,1,opt,name=basket,proto3" json:"basket,omitempty"`
}

func (x *GetBasketResponse) Reset() {
	*x = GetBasketResponse{}
	mi := &file_api_basket_v1_basket_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBasketResponse) ProtoMessage() {}

func (x *GetBasketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_basket_v1_basket_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBasketResponse.ProtoReflect.Descriptor instead.
func (*GetBasketResponse) Descriptor() ([]byte, []int) {
	return file_api_basket_v1_basket_proto_rawDescGZIP(), []int{13}
}

func (x *GetBasketResponse) GetBasket() *Basket {
	if x != nil {
		return x.Basket
	}
	return nil
}

var File_api_basket_v1_basket_proto protoreflect.FileDescriptor

var file_api_basket_v1_basket_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x43, 0x0a, 0x06, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0xc3, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x35, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x25, 0x0a,
	0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x62, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x32, 0x96, 0x04, 0x0a, 0x0d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xcb, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68,
	0x65, 0x6e, 0x67, 0x78, 0x75, 0x66, 0x65, 0x6e, 0x67, 0x31, 0x39, 0x39, 0x34, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x2d, 0x61, 0x72, 0x63, 0x68, 0x2d,
	0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x42, 0x58, 0xaa, 0x02, 0x0d, 0x41,
	0x70, 0x69, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x41,
	0x70, 0x69, 0x5c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x41,
	0x70, 0x69, 0x5c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_basket_v1_basket_proto_rawDescOnce sync.Once
	file_api_basket_v1_basket_proto_rawDescData = file_api_basket_v1_basket_proto_rawDesc
)

func file_api_basket_v1_basket_proto_rawDescGZIP() []byte {
	file_api_basket_v1_basket_proto_rawDescOnce.Do(func() {
		file_api_basket_v1_basket_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_basket_v1_basket_proto_rawDescData)
	})
	return file_api_basket_v1_basket_proto_rawDescData
}

var file_api_basket_v1_basket_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_basket_v1_basket_proto_goTypes = []any{
	(*Basket)(nil),                 // 0: api.basket.v1.Basket
	(*Item)(nil),                   // 1: api.basket.v1.Item
	(*StartBasketRequest)(nil),     // 2: api.basket.v1.StartBasketRequest
	(*StartBasketResponse)(nil),    // 3: api.basket.v1.StartBasketResponse
	(*CancelBasketRequest)(nil),    // 4: api.basket.v1.CancelBasketRequest
	(*CancelBasketResponse)(nil),   // 5: api.basket.v1.CancelBasketResponse
	(*CheckoutBasketRequest)(nil),  // 6: api.basket.v1.CheckoutBasketRequest
	(*CheckoutBasketResponse)(nil), // 7: api.basket.v1.CheckoutBasketResponse
	(*AddItemRequest)(nil),         // 8: api.basket.v1.AddItemRequest
	(*AddItemResponse)(nil),        // 9: api.basket.v1.AddItemResponse
	(*RemoveItemRequest)(nil),      // 10: api.basket.v1.RemoveItemRequest
	(*RemoveItemResponse)(nil),     // 11: api.basket.v1.RemoveItemResponse
	(*GetBasketRequest)(nil),       // 12: api.basket.v1.GetBasketRequest
	(*GetBasketResponse)(nil),      // 13: api.basket.v1.GetBasketResponse
}
var file_api_basket_v1_basket_proto_depIdxs = []int32{
	1,  // 0: api.basket.v1.Basket.items:type_name -> api.basket.v1.Item
	0,  // 1: api.basket.v1.GetBasketResponse.basket:type_name -> api.basket.v1.Basket
	2,  // 2: api.basket.v1.BasketService.StartBasket:input_type -> api.basket.v1.StartBasketRequest
	4,  // 3: api.basket.v1.BasketService.CancelBasket:input_type -> api.basket.v1.CancelBasketRequest
	6,  // 4: api.basket.v1.BasketService.CheckoutBasket:input_type -> api.basket.v1.CheckoutBasketRequest
	8,  // 5: api.basket.v1.BasketService.AddItem:input_type -> api.basket.v1.AddItemRequest
	10, // 6: api.basket.v1.BasketService.RemoveItem:input_type -> api.basket.v1.RemoveItemRequest
	12, // 7: api.basket.v1.BasketService.GetBasket:input_type -> api.basket.v1.GetBasketRequest
	3,  // 8: api.basket.v1.BasketService.StartBasket:output_type -> api.basket.v1.StartBasketResponse
	5,  // 9: api.basket.v1.BasketService.CancelBasket:output_type -> api.basket.v1.CancelBasketResponse
	7,  // 10: api.basket.v1.BasketService.CheckoutBasket:output_type -> api.basket.v1.CheckoutBasketResponse
	9,  // 11: api.basket.v1.BasketService.AddItem:output_type -> api.basket.v1.AddItemResponse
	11, // 12: api.basket.v1.BasketService.RemoveItem:output_type -> api.basket.v1.RemoveItemResponse
	13, // 13: api.basket.v1.BasketService.GetBasket:output_type -> api.basket.v1.GetBasketResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_basket_v1_basket_proto_init() }
func file_api_basket_v1_basket_proto_init() {
	if File_api_basket_v1_basket_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_basket_v1_basket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_basket_v1_basket_proto_goTypes,
		DependencyIndexes: file_api_basket_v1_basket_proto_depIdxs,
		MessageInfos:      file_api_basket_v1_basket_proto_msgTypes,
	}.Build()
	File_api_basket_v1_basket_proto = out.File
	file_api_basket_v1_basket_proto_rawDesc = nil
	file_api_basket_v1_basket_proto_goTypes = nil
	file_api_basket_v1_basket_proto_depIdxs = nil
}
