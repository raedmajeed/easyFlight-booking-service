package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type ConfigParams struct {
	DBHost       string `mapstructure:"DBHOST"`
	DBName       string `mapstructure:"DBNAME"`
	DBUser       string `mapstructure:"DBUSER"`
	DBPort       string `mapstructure:"DBPORT"`
	DBPassword   string `mapstructure:"DBPASSWORD"`
	PORT         string `mapstructure:"PORT"`
	BSERVICEPORT string `mapstructure:"BSERVICEPORT"`
	REDISHOST    string `mapstructure:"REDISHOST"`
	SECRETKEY    string `mapstructure:"SECRETKEY"`
	SERVICETOKEN string `mapstructure:"SERVICETOKEN"`
	TOKEN        string `mapstructure:"TOKEN"`
	SID          string `mapstructure:"SID"`
}

var envs = []string{
	"DBHOST", "DBNAME", "DBSUER", "DBPORT", "DBPASSWORD", "PORT", "ADMINPORT", "REDISHOST", "SECRETKEY", "SERVICETOKEN", "SID", "TOKEN",
}

func Configuration() (*ConfigParams, error, *redis.Client) {
	var cfg ConfigParams
	viper.SetConfigFile("../../.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Unable to load env values, err: %v", err.Error())
		return &ConfigParams{}, err, nil
	}

	for _, e := range envs {
		if err := viper.BindEnv(e); err != nil {
			return &cfg, err, nil
		}
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to unmarshal values, err: %v", err.Error())
	}

	if err := validator.New().Struct(&cfg); err != nil {
		return &cfg, err, nil
	}

	redis := connectToRedis(&cfg)
	return &cfg, err, redis
}

func connectToRedis(cfg *ConfigParams) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.REDISHOST,
		Password: "",
		DB:       2,
	})
	return client
}
