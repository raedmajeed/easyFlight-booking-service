package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/raedmajeed/booking-service/config"
	inter "github.com/raedmajeed/booking-service/pkg/repository/interfaces"
	"github.com/raedmajeed/booking-service/pkg/service/interfaces"
)

type BookingServiceStruct struct {
	repo  inter.BookingRepostory
	redis *redis.Client
	cfg   *config.ConfigParams
	kf    *config.KafkaConfigStruct
}

func NewBookingService(repo inter.BookingRepostory, redis *redis.Client, cfg *config.ConfigParams, kf *config.KafkaConfigStruct) interfaces.BookingService {
	return &BookingServiceStruct{
		repo:  repo,
		redis: redis,
		cfg:   cfg,
		kf:    kf,
	}
}
