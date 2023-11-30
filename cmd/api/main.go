package main

import (
	"github.com/raedmajeed/booking-service/config"
	client "github.com/raedmajeed/booking-service/pkg/api"
	"github.com/raedmajeed/booking-service/pkg/di"
	"log"
)

func main() {
	cfg, err, redis := config.Configuration()
	twilio := config.SetupTwilio(cfg)
	c, err := client.NewGrpcClient(cfg)
	if err != nil {
		log.Printf("unable to load env values, err: %v", err.Error())
		return
	}
	di.InitApi(cfg, redis, twilio, *c)
}
