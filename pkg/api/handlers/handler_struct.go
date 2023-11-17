package handlers

import (
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/service/interfaces"
)

type BookingHandler struct {
	svc interfaces.BookingService
	pb.BookingServer
}

func NewBookingHandler(svc interfaces.BookingService) *BookingHandler {
	return &BookingHandler{
		svc: svc,
	}
}