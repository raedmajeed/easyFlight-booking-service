package di

import "github.com/raedmajeed/booking-service/pkg/service/interfaces"

func CronJob(svc interfaces.BookingService) {
	go svc.CheckAndSendBookingReminder()
}
