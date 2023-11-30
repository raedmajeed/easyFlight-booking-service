package db

import (
	"fmt"
	"github.com/raedmajeed/booking-service/pkg/DOM"
	"log"

	"github.com/raedmajeed/booking-service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnect(cfg *config.ConfigParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection to DB %s Failed, Error: %s", cfg.DBName, err)
		return nil, err
	}

	// MIGRATING DB
	err = database.AutoMigrate(&DOM.UserData{}, &DOM.Booking{})

	if err != nil {
		log.Printf("unable to migrate db, err: %v", err)
		return nil, err
	}

	return database, nil
}
