package handlers

import (
	"context"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"log"
)

func (handler *BookingHandler) RegisterSelectSeat(ctx context.Context, p *pb.SeatSelectRequest) (*pb.SeatSelectResponse, error) {
	response, err := handler.svc.SelectSeat(ctx, p)
	if err != nil {
		log.Println("error registering seat to the database RegisterSelectSeat() - bookingService")
		return nil, err
	}
	return response, nil
}
