package easyFlight_booking_service

import (
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

type ConfigParams struct {
	DBHost            string `mapstructure:"DBHOST"`
	DBName            string `mapstructure:"DBNAME"`
	DBUser            string `mapstructure:"DBUSER"`
	DBPort            string `mapstructure:"DBPORT"`
	DBPassword        string `mapstructure:"DBPASSWORD"`
	PORT              string `mapstructure:"PORT"`
	BSERVICEPORT      string `mapstructure:"BSERVICEPORT"`
	REDISHOST         string `mapstructure:"REDISHOST"`
	SECRETKEY         string `mapstructure:"SECRETKEY"`
	SERVICETOKEN      string `mapstructure:"SERVICETOKEN"`
	TOKEN             string `mapstructure:"TOKEN"`
	SID               string `mapstructure:"SID"`
	ADMINBOOKINGPORT  string `mapstructure:"ADMINBOOKINGPORT"`
	RAZORPAYKEYID     string `mapstructure:"RAZORPAYKEYID"`
	RAZORPAYSECRETKEY string `mapstructure:"RAZORPAYSECRETKEY"`
}

func Configuration() (*ConfigParams, error, *redis.Client) {
	cfg := ConfigParams{}
	//if err := godotenv.Load("../../.env"); err != nil {
	//	os.Exit(1)
	//}

	cfg.DBName = os.Getenv("DBNAME")
	cfg.DBHost = os.Getenv("DBHOST")
	cfg.DBUser = os.Getenv("DBUSER")
	cfg.DBPort = os.Getenv("DBPORT")
	cfg.DBPassword = os.Getenv("DBPASSWORD")
	cfg.PORT = os.Getenv("PORT")
	cfg.BSERVICEPORT = os.Getenv("BSERVICEPORT")
	cfg.REDISHOST = os.Getenv("REDISHOST")
	cfg.SECRETKEY = os.Getenv("SECRETKEY")
	cfg.SERVICETOKEN = os.Getenv("SERVICETOKEN")
	cfg.TOKEN = os.Getenv("TOKEN")
	cfg.SID = os.Getenv("SID")
	cfg.ADMINBOOKINGPORT = os.Getenv("ADMINBOOKINGPORT")
	cfg.RAZORPAYKEYID = os.Getenv("RAZORPAYKEYID")
	cfg.RAZORPAYSECRETKEY = os.Getenv("RAZORPAYSECRETKEY")

	log.Println("version 10 -> docker", cfg)

	redisClient := connectToRedis(&cfg)
	return &cfg, nil, redisClient
}

func connectToRedis(cfg *ConfigParams) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.REDISHOST,
		Password: "",
		DB:       2,
	})
	return client
}
