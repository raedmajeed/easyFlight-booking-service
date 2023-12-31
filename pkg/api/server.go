package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/booking-service/config"
	"github.com/raedmajeed/booking-service/pkg/api/handlers"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/service/interfaces"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	E   *gin.Engine
	cfg *config.ConfigParams
	svc interfaces.BookingService
}

func NewServer(cfg *config.ConfigParams, handler *handlers.BookingHandler, svc interfaces.BookingService) (*Server, error) {
	//newCont, cancel := context.WithCancel(context.Background())
	//defer cancel()

	//kf := config.NewKafkaReaderConnect(svc)
	//go kf.SearchReaderMethod(newCont)
	err := NewGrpcServer(cfg, handler)

	if err != nil {
		log.Println("error connecting to gRPC server")
		return nil, err
	}
	r := gin.Default()
	return &Server{
		E:   r,
		cfg: cfg,
		svc: svc,
	}, nil
}

func NewGrpcServer(cfg *config.ConfigParams, handler *handlers.BookingHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", cfg.BSERVICEPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error Connecting to gRPC server")
		return err
	}
	grp := grpc.NewServer()
	pb.RegisterBookingServer(grp, handler)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}

	log.Printf("listening on gRPC server %v", cfg.BSERVICEPORT)
	err = grp.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}

func (s *Server) ServerStart() error {
	err := s.E.Run(":" + s.cfg.PORT)
	if err != nil {
		log.Println("error starting server")
		return err
	}
	log.Println("Server started")
	return nil
}
