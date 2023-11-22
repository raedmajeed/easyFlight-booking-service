package repository

import (
	// "github.com/raedmajeed/booking-service/pkg/repository/interfaces"

	interfaces "github.com/raedmajeed/booking-service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type BookingRepositoryStruct struct {
	DB *gorm.DB
}

func NewBookingRepository(db *gorm.DB) interfaces.BookingRepository {
	return &BookingRepositoryStruct{
		DB: db,
	}
}
