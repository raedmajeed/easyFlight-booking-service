package di

import (
	"github.com/go-redis/redis/v8"
	"github.com/raedmajeed/booking-service/config"
	api "github.com/raedmajeed/booking-service/pkg/api"
	pkg "github.com/raedmajeed/booking-service/pkg/api"
	"github.com/raedmajeed/booking-service/pkg/api/handlers"
	"github.com/raedmajeed/booking-service/pkg/db"
	"github.com/raedmajeed/booking-service/pkg/repository"
	"github.com/raedmajeed/booking-service/pkg/service"
)

func InitApi(cfg *config.ConfigParams, redis *redis.Client, kf *config.KafkaConfigStruct) (*pkg.Server, error) {

	// db connection
	DB, err := db.NewDBConnect(cfg)
	if err != nil {
		return nil, err
	}
	repo := repository.NewBookingRepository(DB)
	svc := service.NewBookingService(repo, redis, cfg, kf)
	hdlr := handlers.NewBookingHandler(svc)
	server, err := api.NewServer(cfg, hdlr, kf)
	if err != nil {
		return nil, err
	}
	return server, nil
}
