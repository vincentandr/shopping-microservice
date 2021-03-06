package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db "github.com/vincentandr/shopping-microservice/cart/internal/db"
	rmqCart "github.com/vincentandr/shopping-microservice/cart/internal/pubsub"
	pb "github.com/vincentandr/shopping-microservice/internal/proto/cart"
	catalogpb "github.com/vincentandr/shopping-microservice/internal/proto/catalog"
	paymentpb "github.com/vincentandr/shopping-microservice/internal/proto/payment"
)

const (
	CACHE_EXPIRE = 15 * time.Minute
)

type Server struct {
	pb.UnimplementedCartServiceServer
	CatalogClient catalogpb.CatalogServiceClient
	PaymentClient paymentpb.PaymentServiceClient
	Repo *db.Repository
	RmqConsumer *rmqCart.RbmqListener
}

func (s *Server) Grpc_GetCartItems(ctx context.Context, in *pb.GetCartItemsRequest) (*pb.ItemsResponse, error) {
	// Get product ids and its quantity in cart by userId
	res, err := s.Repo.GetCartItems(ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	// Return empty response if there is no items in cart
	if len(res) == 0 {
		return &pb.ItemsResponse{}, nil
	}

	// Get Product ID Keys from map
	ids := GetMapKeys(res)

	// RPC call catalog server to get cart products' names
	products, err := s.CatalogClient.Grpc_GetProductsByIds(ctx, &catalogpb.GetProductsByIdsRequest{ProductIds: ids})
	if err != nil {
		return nil, err
	}

	// Return response in format product id, product name, qty in cart, desc, and image
	items, err := AppendItemToResponse(products, res)
	if err != nil{
		return nil, err
	}

	return items, nil
}

func (s *Server) Grpc_AddOrUpdateCart(ctx context.Context, in *pb.AddOrUpdateCartRequest) (*pb.ItemsResponse, error) {
	res, err := s.Repo.AddOrUpdateCart(ctx, CACHE_EXPIRE, in.UserId, in.ProductId, int(in.NewQty))
	if err != nil{
		return nil, err
	}

	// Return empty response if there is no items in cart
	if len(res) == 0 {
		return &pb.ItemsResponse{}, nil
	}

	// Get Product ID Keys from map
	ids := GetMapKeys(res)

	// RPC call catalog server to get cart products' names
	products, err := s.CatalogClient.Grpc_GetProductsByIds(ctx, &catalogpb.GetProductsByIdsRequest{ProductIds: ids})
	if err != nil {
		return nil, err
	}

	// Return response in format product id, product name, qty, desc, and image in cart
	items, err := AppendItemToResponse(products, res)
	if err != nil{
		return nil, err
	}

    return items, nil
}

func (s *Server) Grpc_RemoveItemFromCart(ctx context.Context, in *pb.RemoveItemFromCartRequest) (*pb.ItemsResponse, error) {
	res, err := s.Repo.RemoveItemFromCart(ctx, in.UserId, in.ProductId)
	if err != nil{
		return nil, err
	}

	// Return empty response if there is no items in cart
	if len(res) == 0 {
		return &pb.ItemsResponse{}, nil
	}

    // Get Product ID Keys from map
	ids := GetMapKeys(res)

	// RPC call catalog server to get cart products' names
	products, err := s.CatalogClient.Grpc_GetProductsByIds(ctx, &catalogpb.GetProductsByIdsRequest{ProductIds: ids})
	if err != nil {
		return nil, err
	}

	// Return response in format product id, product name, qty, desc, and image in cart
	items, err := AppendItemToResponse(products, res)
	if err != nil{
		return nil, err
	}

    return items, nil
}

func (s *Server) Grpc_RemoveAllCartItems(ctx context.Context, in *pb.RemoveAllCartItemsRequest) (*pb.ItemsResponse, error) {
	_, err := s.Repo.RemoveAllCartItems(ctx, in.UserId)
	if err != nil{
		return nil, err
	}

	// Return empty response since there is no items in cart
	return &pb.ItemsResponse{}, nil
}

func (s *Server) Grpc_Checkout(ctx context.Context, in *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	// Get user id's cart items
	res, err := s.Grpc_GetCartItems(ctx, &pb.GetCartItemsRequest{UserId: in.UserId})
	if err != nil{
		return nil, err
	}

	// RPC call to payment checkout to create order and return order id
	itemsForOrder := make([]*paymentpb.ItemResponse, len(res.Products))
	var subtotal float32 = 0

	for i, item := range res.Products {
		itemsForOrder[i] = &paymentpb.ItemResponse{ProductId: item.ProductId, Name: item.Name, Price: item.Price, Qty: item.Qty, Desc: item.Desc, Image: item.Image}
		
		// Calc subtotal
		subtotal += item.Price * float32(item.Qty)
	}

	response, err := s.PaymentClient.Grpc_PaymentCheckout(ctx, &paymentpb.CheckoutRequest{UserId: in.UserId, Items: itemsForOrder, Subtotal: subtotal})
	if err != nil{
		return nil, err
	}

	return &pb.CheckoutResponse{OrderId: response.OrderId}, nil
}

func GetMapKeys(hm map[string]string) ([]string) {
	ids := make([]string, len(hm))
	i := 0
	for k := range hm {
		ids[i] = k
		i++
	}

	return ids
}

func AppendItemToResponse(catalogRes *catalogpb.GetProductsResponse, hm map[string]string) (*pb.ItemsResponse, error){
	items := pb.ItemsResponse{}
	var subtotal float32

	for _, prod := range catalogRes.Products {
		qty, err := strconv.Atoi(hm[prod.ProductId])
		if err != nil{
			return nil, fmt.Errorf("failed to convert qty from string to int: %v", err)
		}

		// Calc subtotal
		subtotal += prod.Price * float32(qty)	

		item := &pb.ItemResponse{ProductId: prod.ProductId, Name: prod.Name, Price: prod.Price, Qty: int32(qty), Stock: prod.Qty, Desc: prod.Desc, Image: prod.Image}

		items.Products = append(items.Products, item)
	}

	items.Subtotal = subtotal

	return &items, nil
}