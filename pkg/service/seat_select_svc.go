package service

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"log"
)

type JsonRequest struct {
	Seats []string `json:"seats"`
}

func (svc *BookingServiceStruct) SelectSeat(ctx context.Context, p *pb.SeatSelectRequest) (*pb.SeatSelectResponse, error) {
	// fetch the booking table using pnr number
	booking, err := svc.repo.FindBookingByPNR(p.PNR)
	if err != nil {
		log.Println("error fetching booking from table, err: ", err.Error())
		return nil, err
	}

	fmt.Println(booking.CancelledStatus, "booking status")
	if booking.BookingStatus != "CONFIRMED" {
		log.Println("user not confirmed")
		return nil, err
	}

	if booking.CancelledStatus != "false" {
		log.Println("ticket cancelled, cannot add seat now")
		return nil, err
	}

	bookingDetails, err := svc.repo.FindNumberOfBookedUsers(ctx, p.PNR)
	travellerLists, err := svc.repo.FindNumberOfTravellers(bookingDetails.ID)
	if err != nil {
		log.Println("error finding traveller list")
		return nil, err
	}

	travellerDetail, err := svc.repo.FindTravellerById(travellerLists[0].TravellerId)

	if len(p.SeatArray) != len(travellerLists) {
		log.Println("the number of seats selected is greater than number of travellers or less")
		return nil, errors.New("the number of seats selected is greater than number of travellers or les")
	}

	log.Println("sending response to admin-airline service")
	response, err := svc.client.RegisterSelectSeat(ctx, &pb.SeatRequest{
		PNR:           p.PNR,
		SeatArray:     p.SeatArray,
		FlightChartId: p.FlightChartId,
		Email:         booking.Email,
		Economy:       travellerDetail.Economy,
	})
	if err != nil {
		log.Println("error registering seats to the database SelectSeat() - booking-service")
		return nil, err
	}
	return &pb.SeatSelectResponse{
		PNR:     response.PNR,
		SeatNos: response.SeatNos,
	}, nil
}
