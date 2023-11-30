package repository

import (
	"errors"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	"gorm.io/gorm"
	"log"
)

func (repo *BookingRepositoryStruct) FindUserByEmail(email string) (*dom.UserData, error) {
	var airline dom.UserData
	result := repo.DB.Where("email = ?", email).First(&airline)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of user %v", email)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &airline, nil
}

func (repo *BookingRepositoryStruct) CreateUser(airline *dom.UserData) (*dom.UserData, error) {
	result := repo.DB.Create(airline)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Println("duplicate key found")
			return nil, result.Error
		} else {
			return nil, result.Error
		}
	}
	return airline, nil
}

func (repo *BookingRepositoryStruct) FindBookingByPNR(pnr string) (*dom.Booking, error) {
	var booking dom.Booking
	result := repo.DB.Where("PNR = ?", pnr).First(&booking)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of user %v", pnr)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &booking, nil
}
