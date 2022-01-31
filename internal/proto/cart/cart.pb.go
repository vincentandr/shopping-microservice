// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: cart.proto

package cartpb

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

type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CheckoutRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CheckoutResponse) Reset() {
	*x = CheckoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutResponse) ProtoMessage() {}

func (x *CheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutResponse.ProtoReflect.Descriptor instead.
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{1}
}

func (x *CheckoutResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetCartItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetCartItemsRequest) Reset() {
	*x = GetCartItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartItemsRequest) ProtoMessage() {}

func (x *GetCartItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartItemsRequest.ProtoReflect.Descriptor instead.
func (*GetCartItemsRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{2}
}

func (x *GetCartItemsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string  `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price     float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Qty       int32   `protobuf:"varint,4,opt,name=qty,proto3" json:"qty,omitempty"`
	Stock     int32   `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Desc      string  `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc,omitempty"`
	Image     string  `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ItemResponse) Reset() {
	*x = ItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemResponse) ProtoMessage() {}

func (x *ItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemResponse.ProtoReflect.Descriptor instead.
func (*ItemResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{3}
}

func (x *ItemResponse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ItemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ItemResponse) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *ItemResponse) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *ItemResponse) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ItemResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type ItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ItemResponse `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Subtotal float32         `protobuf:"fixed32,2,opt,name=subtotal,proto3" json:"subtotal,omitempty"`
}

func (x *ItemsResponse) Reset() {
	*x = ItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemsResponse) ProtoMessage() {}

func (x *ItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemsResponse.ProtoReflect.Descriptor instead.
func (*ItemsResponse) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{4}
}

func (x *ItemsResponse) GetProducts() []*ItemResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ItemsResponse) GetSubtotal() float32 {
	if x != nil {
		return x.Subtotal
	}
	return 0
}

type AddOrUpdateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	NewQty    int32  `protobuf:"varint,3,opt,name=new_qty,json=newQty,proto3" json:"new_qty,omitempty"`
}

func (x *AddOrUpdateCartRequest) Reset() {
	*x = AddOrUpdateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrUpdateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrUpdateCartRequest) ProtoMessage() {}

func (x *AddOrUpdateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrUpdateCartRequest.ProtoReflect.Descriptor instead.
func (*AddOrUpdateCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{5}
}

func (x *AddOrUpdateCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddOrUpdateCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddOrUpdateCartRequest) GetNewQty() int32 {
	if x != nil {
		return x.NewQty
	}
	return 0
}

type RemoveItemFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *RemoveItemFromCartRequest) Reset() {
	*x = RemoveItemFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveItemFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemFromCartRequest) ProtoMessage() {}

func (x *RemoveItemFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveItemFromCartRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveItemFromCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveItemFromCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type RemoveAllCartItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RemoveAllCartItemsRequest) Reset() {
	*x = RemoveAllCartItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAllCartItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAllCartItemsRequest) ProtoMessage() {}

func (x *RemoveAllCartItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAllCartItemsRequest.ProtoReflect.Descriptor instead.
func (*RemoveAllCartItemsRequest) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveAllCartItemsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_cart_proto protoreflect.FileDescriptor

var file_cart_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x61,
	0x72, 0x74, 0x70, 0x62, 0x22, 0x2a, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x2d, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xa9, 0x01, 0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x5d, 0x0a, 0x0d, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x69, 0x0a, 0x16, 0x41, 0x64,
	0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x6e, 0x65, 0x77, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x65, 0x77, 0x51, 0x74, 0x79, 0x22, 0x53, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x19, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x32, 0x9d, 0x03, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x49, 0x0a, 0x11, 0x47, 0x72, 0x70, 0x63, 0x5f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x14, 0x47,
	0x72, 0x70, 0x63, 0x5f, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x17,
	0x47, 0x72, 0x70, 0x63, 0x5f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x46, 0x72, 0x6f, 0x6d, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x17, 0x47, 0x72, 0x70, 0x63, 0x5f, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x21,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c,
	0x6c, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x72,
	0x70, 0x63, 0x5f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x6e, 0x64, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_proto_rawDescOnce sync.Once
	file_cart_proto_rawDescData = file_cart_proto_rawDesc
)

func file_cart_proto_rawDescGZIP() []byte {
	file_cart_proto_rawDescOnce.Do(func() {
		file_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_proto_rawDescData)
	})
	return file_cart_proto_rawDescData
}

var file_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cart_proto_goTypes = []interface{}{
	(*CheckoutRequest)(nil),           // 0: cartpb.CheckoutRequest
	(*CheckoutResponse)(nil),          // 1: cartpb.CheckoutResponse
	(*GetCartItemsRequest)(nil),       // 2: cartpb.GetCartItemsRequest
	(*ItemResponse)(nil),              // 3: cartpb.ItemResponse
	(*ItemsResponse)(nil),             // 4: cartpb.ItemsResponse
	(*AddOrUpdateCartRequest)(nil),    // 5: cartpb.AddOrUpdateCartRequest
	(*RemoveItemFromCartRequest)(nil), // 6: cartpb.RemoveItemFromCartRequest
	(*RemoveAllCartItemsRequest)(nil), // 7: cartpb.RemoveAllCartItemsRequest
}
var file_cart_proto_depIdxs = []int32{
	3, // 0: cartpb.ItemsResponse.products:type_name -> cartpb.ItemResponse
	2, // 1: cartpb.CartService.Grpc_GetCartItems:input_type -> cartpb.GetCartItemsRequest
	5, // 2: cartpb.CartService.Grpc_AddOrUpdateCart:input_type -> cartpb.AddOrUpdateCartRequest
	6, // 3: cartpb.CartService.Grpc_RemoveItemFromCart:input_type -> cartpb.RemoveItemFromCartRequest
	7, // 4: cartpb.CartService.Grpc_RemoveAllCartItems:input_type -> cartpb.RemoveAllCartItemsRequest
	0, // 5: cartpb.CartService.Grpc_Checkout:input_type -> cartpb.CheckoutRequest
	4, // 6: cartpb.CartService.Grpc_GetCartItems:output_type -> cartpb.ItemsResponse
	4, // 7: cartpb.CartService.Grpc_AddOrUpdateCart:output_type -> cartpb.ItemsResponse
	4, // 8: cartpb.CartService.Grpc_RemoveItemFromCart:output_type -> cartpb.ItemsResponse
	4, // 9: cartpb.CartService.Grpc_RemoveAllCartItems:output_type -> cartpb.ItemsResponse
	1, // 10: cartpb.CartService.Grpc_Checkout:output_type -> cartpb.CheckoutResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cart_proto_init() }
func file_cart_proto_init() {
	if File_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutRequest); i {
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
		file_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutResponse); i {
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
		file_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartItemsRequest); i {
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
		file_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemResponse); i {
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
		file_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemsResponse); i {
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
		file_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrUpdateCartRequest); i {
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
		file_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveItemFromCartRequest); i {
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
		file_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAllCartItemsRequest); i {
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
			RawDescriptor: file_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_proto_goTypes,
		DependencyIndexes: file_cart_proto_depIdxs,
		MessageInfos:      file_cart_proto_msgTypes,
	}.Build()
	File_cart_proto = out.File
	file_cart_proto_rawDesc = nil
	file_cart_proto_goTypes = nil
	file_cart_proto_depIdxs = nil
}
