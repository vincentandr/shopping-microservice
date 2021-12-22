// gRPC Catalog Client in cart

package clients

import (
	"context"
	"fmt"

	pb "github.com/vincentandr/shopping-microservice/src/catalog/catalogpb"
	"google.golang.org/grpc"
)

const (
	catalogRpcPort = ":50051"
)

var (
	catalogClientConn *grpc.ClientConn
	catalogClient pb.CatalogServiceClient
)

func NewCatalogClient() error{
	catalogClientConn, err := grpc.Dial("localhost" + catalogRpcPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("failed to connect to RPC client: %v", err)
	}

	catalogClient = pb.NewCatalogServiceClient(catalogClientConn)

	return nil
}

func DisconnectCatalogClient() error{
	err := catalogClientConn.Close()

	if err != nil{
		return fmt.Errorf("failed to disconnect from RPC client: %v", err)
	}

	return err
}

func GetProductsByIds(ctx context.Context, productIds []int) (*pb.GetProductsByIdsResponse , error){
	var convProductIds []int32

	for _, val := range productIds {
		convProductIds = append(convProductIds, int32(val))
	}

	products, err := catalogClient.GetProductsByIds(ctx, &pb.GetProductsByIdsRequest{ProductIds: convProductIds})
	if err != nil{
		return nil, err
	}

	return products, nil

}