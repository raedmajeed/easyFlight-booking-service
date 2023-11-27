package repository

import dom "github.com/raedmajeed/booking-service/pkg/DOM"

func (repo *BookingRepositoryStruct) CreateBookedTravellers(booking *dom.Booking) error {
	return repo.DB.Create(booking).Error
}
