package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/utils"
	"log"
	"math/rand"
)

func (svc *BookingServiceStruct) AddTraveller(ctx context.Context, request *pb.TravellerRequest) (*pb.TravellerResponse, error) {
	token := request.Token
	if token == "" {
		return nil, errors.New("search token missing")
	}
	_, err := utils.ValidateSearchToken(token, *svc.cfg)
	if err != nil {
		return nil, fmt.Errorf("\"there is error in AddTraveller() method, err: %v", err.Error())
	}

	user, err := svc.repo.FindUserByEmail(request.Email)
	var cf pb.SearchSelectResponse
	val := svc.redis.Get(ctx, token+"1")
	if err := json.Unmarshal([]byte(val.Val()), &cf); err != nil {
		return nil, err
	}
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

	var dep string
	var arr string
	if len(cf.DirectFlight.FlightPath.FlightSegment) > 0 {
		dep = cf.DirectFlight.FlightPath.FlightSegment[0].DepartureAirport
		arr = cf.DirectFlight.FlightPath.FlightSegment[0].ArrivalAirport
	}

	booking := dom.Booking{
		BookingReference: bookingRefID,
		BookingStatus:    "PENDING",
		UserId:           user.ID,
		Bookings:         travellers,
		PNR:              generatePNR(bookingRefID),
		DepartureAirport: dep,
		ArrivalAirport:   arr,
		Email:            request.Email,
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

func generatePNR(ref string) string {
	s1 := ref[:2]
	s2 := fmt.Sprintf("%v", rand.Int())
	s3 := string(ref[3])
	s4 := fmt.Sprintf("PR")
	return s4 + s1 + s2 + s3
}
