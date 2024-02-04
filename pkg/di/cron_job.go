package di

import (
	"github.com/raedmajeed/booking-service/pkg/service/interfaces"
	"time"
)

func CronJob(svc interfaces.BookingService) {
	go func() {
		for {
			select {
			case <-time.After(time.Hour * 24):
				svc.CheckAndSendBookingReminder()
			}
		}
	}()
}
