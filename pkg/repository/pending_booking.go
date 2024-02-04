package repository

import (
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	"time"
)

func (repo *BookingRepositoryStruct) FindPendingBooking() ([]*dom.Booking, error) {
	var data []*dom.Booking
	threeDaysAgo := time.Now().AddDate(0, 0, -3)
	err := repo.DB.Where("booking_status = 'PENDING' AND created_at > ?", threeDaysAgo).Find(&data).Error
	return data, err
}
