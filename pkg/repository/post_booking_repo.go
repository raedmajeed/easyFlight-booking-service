package repository

import (
	"context"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
)

func (repo *BookingRepositoryStruct) FindNumberOfBookedUsers(ctx context.Context, pnr string) (*dom.Booking, error) {
	var data dom.Booking
	result := repo.DB.Preload("traveller_booking").Where("pnr = ?", pnr).First(&data).Error
	return &data, result
}

func (repo *BookingRepositoryStruct) FindNumberOfTravellers(id uint) ([]*dom.TravellerBooking, error) {
	var data []*dom.TravellerBooking
	result := repo.DB.Raw("SELECT * FROM traveller_booking where booking_id = ?", id).Scan(&data).Error
	return data, result
}

func (repo *BookingRepositoryStruct) FindTravellerById(id uint) (*dom.Traveller, error) {
	var data dom.Traveller
	result := repo.DB.Where("id = ?", id).First(&data).Error
	return &data, result
}

func (repo *BookingRepositoryStruct) UpdateTravellerSeat(id uint, seat string) error {
	return repo.DB.Model(&dom.Traveller{}).Where("id = ?", id).Update("seat_no", seat).Error
}
