package handlers

import (
	"context"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"log"
)

func (handler *BookingHandler) RegisterPNRLogin(ctx context.Context, p *pb.PNRLoginRequest) (*pb.PNRLoginResponse, error) {
	response, err := handler.svc.ConfirmedUserLogin(ctx, p)
	if err != nil {
		log.Println("unable to get the search details")
		return nil, err
	}
	return response, nil
}
