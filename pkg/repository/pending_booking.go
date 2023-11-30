package repository

import dom "github.com/raedmajeed/booking-service/pkg/DOM"

func (repo *BookingRepositoryStruct) FindPendingBooking() ([]*dom.Booking, error) {
	var data []*dom.Booking
	err := repo.DB.Where("booking_status = PENDING").Find(&data).Error
	return data, err
}
