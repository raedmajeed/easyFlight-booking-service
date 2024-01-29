package pkg

import (
	"github.com/raedmajeed/booking-service/config"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewGrpcClient(cfg *config.ConfigParams) (*pb.AdminServiceClient, error) {
	log.Printf("dialing admin port: %v", cfg.ADMINBOOKINGPORT)
	client, err := grpc.Dial(cfg.ADMINBOOKINGPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error starting grpc server")
	}
	adminClient := pb.NewAdminServiceClient(client)
	return &adminClient, nil
}
