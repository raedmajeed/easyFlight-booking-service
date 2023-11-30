package handlers

import (
	"context"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
)

func (handler *BookingHandler) RegisterConfirmBooking(ctx context.Context, request *pb.ConfirmBookingRequest) (*pb.ConfirmBookingResponse, error) {
	response, err := handler.svc.ConfirmBooking(ctx, request)
	return response, err
}

func (handler *BookingHandler) RegisterOnlinePayment(ctx context.Context, request *pb.OnlinePaymentRequest) (*pb.OnlinePaymentResponse, error) {
	response, err := handler.svc.OnlinePayment(ctx, request)
	return response, err
}

func (handler *BookingHandler) ResisterPaymentConfirmed(ctx context.Context, request *pb.PaymentConfirmedRequest) (*pb.PaymentConfirmedResponse, error) {
	response, err := handler.svc.PaymentConfirmed(ctx, request)
	return response, err
}
