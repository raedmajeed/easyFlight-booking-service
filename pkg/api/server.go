package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/booking-service"
	"github.com/raedmajeed/booking-service/pkg/api/handlers"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/service/interfaces"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	E      *gin.Engine
	cfg    *easyFlight_booking_service.ConfigParams
	svc    interfaces.BookingService
	client pb.AdminServiceClient
}

func NewServer(cfg *easyFlight_booking_service.ConfigParams, handler *handlers.BookingHandler, svc interfaces.BookingService) {
	err := NewGrpcServer(cfg, handler)
	if err != nil {
		log.Println("error connecting to gRPC server", err.Error())
	}
}

func NewGrpcServer(cfg *easyFlight_booking_service.ConfigParams, handler *handlers.BookingHandler) error {
	addr := fmt.Sprintf(":%s", cfg.BSERVICEPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error Connecting to gRPC server 1", err.Error())
		return err
	}
	grp := grpc.NewServer()
	pb.RegisterBookingServer(grp, handler)
	if err != nil {
		log.Println("error connecting to gRPC server 2", err.Error())
		return err
	}
	log.Printf("listening on gRPC server listening from API-service %v", cfg.BSERVICEPORT)
	err = grp.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server 3", err.Error())
		return err
	}
	return nil
}
