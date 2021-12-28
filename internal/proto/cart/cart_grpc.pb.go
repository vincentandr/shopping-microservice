// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cartpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	GetCartItems(ctx context.Context, in *GetCartItemsRequest, opts ...grpc.CallOption) (*ItemsResponse, error)
	AddOrUpdateCart(ctx context.Context, in *AddOrUpdateCartRequest, opts ...grpc.CallOption) (*ItemsResponse, error)
	RemoveItemFromCart(ctx context.Context, in *RemoveItemFromCartRequest, opts ...grpc.CallOption) (*ItemsResponse, error)
	RemoveAllCartItems(ctx context.Context, in *RemoveAllCartItemsRequest, opts ...grpc.CallOption) (*ItemsResponse, error)
	Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) GetCartItems(ctx context.Context, in *GetCartItemsRequest, opts ...grpc.CallOption) (*ItemsResponse, error) {
	out := new(ItemsResponse)
	err := c.cc.Invoke(ctx, "/cartpb.CartService/GetCartItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddOrUpdateCart(ctx context.Context, in *AddOrUpdateCartRequest, opts ...grpc.CallOption) (*ItemsResponse, error) {
	out := new(ItemsResponse)
	err := c.cc.Invoke(ctx, "/cartpb.CartService/AddOrUpdateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveItemFromCart(ctx context.Context, in *RemoveItemFromCartRequest, opts ...grpc.CallOption) (*ItemsResponse, error) {
	out := new(ItemsResponse)
	err := c.cc.Invoke(ctx, "/cartpb.CartService/RemoveItemFromCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveAllCartItems(ctx context.Context, in *RemoveAllCartItemsRequest, opts ...grpc.CallOption) (*ItemsResponse, error) {
	out := new(ItemsResponse)
	err := c.cc.Invoke(ctx, "/cartpb.CartService/RemoveAllCartItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error) {
	out := new(CheckoutResponse)
	err := c.cc.Invoke(ctx, "/cartpb.CartService/Checkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	GetCartItems(context.Context, *GetCartItemsRequest) (*ItemsResponse, error)
	AddOrUpdateCart(context.Context, *AddOrUpdateCartRequest) (*ItemsResponse, error)
	RemoveItemFromCart(context.Context, *RemoveItemFromCartRequest) (*ItemsResponse, error)
	RemoveAllCartItems(context.Context, *RemoveAllCartItemsRequest) (*ItemsResponse, error)
	Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) GetCartItems(context.Context, *GetCartItemsRequest) (*ItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartItems not implemented")
}
func (UnimplementedCartServiceServer) AddOrUpdateCart(context.Context, *AddOrUpdateCartRequest) (*ItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateCart not implemented")
}
func (UnimplementedCartServiceServer) RemoveItemFromCart(context.Context, *RemoveItemFromCartRequest) (*ItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveItemFromCart not implemented")
}
func (UnimplementedCartServiceServer) RemoveAllCartItems(context.Context, *RemoveAllCartItemsRequest) (*ItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAllCartItems not implemented")
}
func (UnimplementedCartServiceServer) Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_GetCartItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cartpb.CartService/GetCartItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartItems(ctx, req.(*GetCartItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddOrUpdateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddOrUpdateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cartpb.CartService/AddOrUpdateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddOrUpdateCart(ctx, req.(*AddOrUpdateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveItemFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveItemFromCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveItemFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cartpb.CartService/RemoveItemFromCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveItemFromCart(ctx, req.(*RemoveItemFromCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveAllCartItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAllCartItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveAllCartItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cartpb.CartService/RemoveAllCartItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveAllCartItems(ctx, req.(*RemoveAllCartItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cartpb.CartService/Checkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).Checkout(ctx, req.(*CheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cartpb.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCartItems",
			Handler:    _CartService_GetCartItems_Handler,
		},
		{
			MethodName: "AddOrUpdateCart",
			Handler:    _CartService_AddOrUpdateCart_Handler,
		},
		{
			MethodName: "RemoveItemFromCart",
			Handler:    _CartService_RemoveItemFromCart_Handler,
		},
		{
			MethodName: "RemoveAllCartItems",
			Handler:    _CartService_RemoveAllCartItems_Handler,
		},
		{
			MethodName: "Checkout",
			Handler:    _CartService_Checkout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart.proto",
}