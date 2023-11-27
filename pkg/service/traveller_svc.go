package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/utils"
	"log"
)

func (svc *BookingServiceStruct) AddTraveller(ctx context.Context, request *pb.TravellerRequest) (*pb.TravellerResponse, error) {
	token := request.Token
	if token == "" {
		return nil, errors.New("search token missing")
	}
	_, err := utils.ValidateSearchToken(token, *svc.cfg)
	if err != nil {
		log.Printf("there is error in AddTraveller() method, err: %v", err.Error())
		return nil, err
	}
	user, err := svc.repo.FindUserByEmail(request.Email)

	bookingRefID := generateBookingReference()
	var travellers []dom.Traveller

	for _, travellerDetail := range request.TravellerDetails {
		traveller := dom.Traveller{
			Name:   travellerDetail.Name,
			Age:    travellerDetail.Age,
			Gender: travellerDetail.Gender,
			UserId: user.ID,
		}
		travellers = append(travellers, traveller)
	}
	booking := dom.Booking{
		BookingReference: bookingRefID,
		BookingStatus:    "PENDING",
		UserId:           user.ID,
		Bookings:         travellers,
	}
	if err := svc.repo.CreateBookedTravellers(&booking); err != nil {
		log.Printf("failed to create bookedTravellers: %v", err)
		return nil, err
	}
	return &pb.TravellerResponse{
		BookingReference: bookingRefID,
	}, nil
}

func generateBookingReference() string {
	ref := uuid.New()
	return ref.String()
}
