package di

import (
	"github.com/go-redis/redis/v8"
	"github.com/raedmajeed/booking-service"
	"github.com/raedmajeed/booking-service/config"
	api "github.com/raedmajeed/booking-service/pkg/api"
	"github.com/raedmajeed/booking-service/pkg/api/handlers"
	"github.com/raedmajeed/booking-service/pkg/db"
	ab "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/repository"
	"github.com/raedmajeed/booking-service/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func InitApi(cfg *easyFlight_booking_service.ConfigParams, redis *redis.Client, twilio *config.TwilioVerify, client ab.AdminServiceClient) {

	singalChan := make(chan os.Signal, 1)
	DB, _ := db.NewDBConnect(cfg)
	kfWriter := config.NewKafkaWriterConnect()
	kfReader := config.NewKafkaReaderConnect()
	repo := repository.NewBookingRepository(DB)
	svc := service.NewBookingService(repo, redis, cfg, kfWriter, kfReader, twilio, client)
	hdlr := handlers.NewBookingHandler(svc)
	go api.NewServer(cfg, hdlr, svc)

	signal.Notify(singalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sign := <-singalChan
	log.Printf("program stopper %v", sign)

	err := kfReader.SearchReader.Close()
	if err != nil {
		log.Println("error closing SearchReader")
	}
	err = kfReader.SearchSelectReader.Close()
	if err != nil {
		log.Println("error closing SearchSelectReader")
	}
}
