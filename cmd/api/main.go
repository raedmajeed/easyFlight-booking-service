package main

import (
	"github.com/raedmajeed/booking-service/config"
	"github.com/raedmajeed/booking-service/pkg/di"
	"log"
)

func main() {
	cfg, err, redis := config.Configuration()
	kafka := config.NewKafkaConfig()
	if err != nil {
		log.Printf("unable to load env values, err: %v", err.Error())
		return
	}
	server, err := di.InitApi(cfg, redis, kafka)
	if err != nil {
		log.Fatalf("Server not starter due to error: %v", err.Error())
		return
	}
	if err = server.ServerStart(); err != nil {
		log.Fatalf("Server not starter due to error: %v", err.Error())
	}
}
