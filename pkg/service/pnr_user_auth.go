package service

import (
	"context"
	"errors"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/utils"
	"gorm.io/gorm"
	"log"
)

func (svc *BookingServiceStruct) ConfirmedUserLogin(ctx context.Context, p *pb.PNRLoginRequest) (*pb.PNRLoginResponse, error) {
	_, err := svc.repo.FindBookingByPNR(p.PNR)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return &pb.PNRLoginResponse{}, err
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return &pb.PNRLoginResponse{}, err
		}
	}

	token, err := utils.GeneratePNRToken(p.PNR, "PNR-USER", svc.cfg)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return &pb.PNRLoginResponse{}, err
	}

	return &pb.PNRLoginResponse{
		PNR:   p.PNR,
		Email: p.Email,
		Token: token,
	}, nil
}
