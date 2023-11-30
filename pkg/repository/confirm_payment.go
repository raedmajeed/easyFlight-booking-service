package repository

import dom "github.com/raedmajeed/booking-service/pkg/DOM"

func (repo *BookingRepositoryStruct) FindBooking(email, bookingReference string) (dom.Booking, error) {
	var data dom.Booking
	err := repo.DB.Where("email = ? and booking_reference = ?", email, bookingReference).First(&data).Error
	return data, err
}

func (repo *BookingRepositoryStruct) UpdateBookings(booking dom.Booking) error {
	return repo.DB.Model(&dom.Booking{}).Where("booking_reference = ?", booking.BookingReference).Updates(booking).Error
}
