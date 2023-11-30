package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/raedmajeed/booking-service/pkg/DOM"
	"github.com/raedmajeed/booking-service/pkg/utils"
	"github.com/segmentio/kafka-go"
)

func (svc *BookingServiceStruct) CheckAndSendBookingReminder() {
	response, err := svc.repo.FindPendingBooking()
	if err != nil {
		marshal := utils.ErrorEmail(err, "CheckAndSendBookingReminder()", "service", "booking-service")
		writeToKafka(svc, marshal)
	}
	for _, r := range response {
		response, _ := svc.repo.FindBooking(r.Email, r.BookingReference)
		if err != nil {
			marshal := utils.ErrorEmail(err, "CheckAndSendBookingReminder()", "service", "booking-service")
			writeToKafka(svc, marshal)
		}
		email := &DOM.EmailMessage{
			Email:   r.Email,
			Subject: "COMPLETE YOUR FLIGHT BOOKING NOW",
			Content: fmt.Sprintf("complete you booking now from %v to %v", response.DepartureAirport, response.ArrivalAirport),
		}
		marshal, _ := json.Marshal(email)
		writeToKafka(svc, marshal)
	}
}

func writeToKafka(svc *BookingServiceStruct, message []byte) {
	_ = svc.kf.EmailWriter.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
}
