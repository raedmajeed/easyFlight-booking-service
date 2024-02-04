package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/raedmajeed/booking-service"
	"github.com/raedmajeed/booking-service/config"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	inter "github.com/raedmajeed/booking-service/pkg/repository/interfaces"
	"github.com/raedmajeed/booking-service/pkg/service/interfaces"
)

type BookingServiceStruct struct {
	repo   inter.BookingRepository
	redis  *redis.Client
	cfg    *easyFlight_booking_service.ConfigParams
	kf     *config.KafkaWriter
	kf2    *config.KafkaReader2
	twilio *config.TwilioVerify
	client pb.AdminServiceClient
}

func NewBookingService(repo inter.BookingRepository, redis *redis.Client,
	cfg *easyFlight_booking_service.ConfigParams, kf *config.KafkaWriter, kf2 *config.KafkaReader2,
	twilio *config.TwilioVerify, client pb.AdminServiceClient) interfaces.BookingService {
	return &BookingServiceStruct{
		repo:   repo,
		redis:  redis,
		cfg:    cfg,
		kf:     kf,
		kf2:    kf2,
		twilio: twilio,
		client: client,
	}
}
