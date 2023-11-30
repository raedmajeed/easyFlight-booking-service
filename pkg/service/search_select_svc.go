package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/utils"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type SelectRequest struct {
	Token        string
	DirectPathId string
	ReturnPathId string
	Adults       int
	Children     int
	Economy      bool
}

func (svc *BookingServiceStruct) SearchSelect(ctx context.Context, request *pb.SearchSelectRequest) (*pb.SearchSelectResponse, error) {
	token := request.Token
	var completeFacilities dom.CompleteFlightFacilities
	claims, err := utils.ValidateSearchToken(token, *svc.cfg)
	if err != nil {
		log.Println("search token is expired or null, please check")
		return nil, err
	}

	directPathID := request.DirectPathId
	returnPathID := request.ReturnPathId

	//type check struct
	selectReq := SelectRequest{
		Token:        token,
		DirectPathId: directPathID,
		ReturnPathId: returnPathID,
		Adults:       claims.Adults,
		Children:     claims.Children,
		Economy:      claims.Economy,
	}

	marshal, err := json.Marshal(selectReq)
	if err != nil {
		log.Println("error marshalling json SearchSelect() - search_select_svc 1")
		return nil, err
	}
	err = svc.kf.SearchSelectWriter.WriteMessages(context.Background(), kafka.Message{
		Value: marshal,
	})
	if err != nil {
		log.Println("failed to write message to kafka from booking service to admin service")
		return nil, err
	}
	fmt.Println("waiting for message from admin service")
	message := svc.kf2.SearchSelectReaderMethod(ctx)

	msg := message.Value
	err = json.Unmarshal(msg, &completeFacilities)
	if err != nil {
		log.Println("error unmarshalling json SearchSelect() - search_select_svc 2")
		return nil, err
	}

	if completeFacilities.DirectFlight.FlightPath.PathId < 0 {
		return nil, errors.New("unable to fetch the flight details")
	}

	response := ConvertToResponse(completeFacilities, claims.Economy)
	svc.redis.Set(ctx, token+"1", message.Value, time.Minute*10)
	return response, nil
}

func ConvertToResponse(cf dom.CompleteFlightFacilities, economy bool) *pb.SearchSelectResponse {
	directFlight := cf.DirectFlight
	returnFlight := cf.ReturnFlight
	var fd1 []*pb.FlightDetails
	var fd2 []*pb.FlightDetails

	for _, f := range directFlight.FlightPath.Flights {
		fare := f.EconomyFare
		if !economy {
			fare = f.BusinessFare
		}
		fd1 = append(fd1, &pb.FlightDetails{
			FlightNumber:     f.FlightNumber,
			Airline:          f.Airline,
			DepartureAirport: f.DepartureAirport,
			DepartureDate:    f.DepartureDate,
			DepartureTime:    f.DepartureTime,
			ArrivalAirport:   f.ArrivalAirport,
			ArrivalDate:      f.ArrivalDate,
			ArrivalTime:      f.ArrivalTime,
			FlightFare:       float32(fare),
		})
	}

	cd := directFlight.Cancellation
	bd := directFlight.Baggage
	bg1 := &pb.Baggage{
		CabinAllowedBreadth: int32(bd.CabinAllowedBreadth),
		CabinAllowedLength:  int32(bd.CabinAllowedLength),
		CabinAllowedWeight:  int32(bd.CabinAllowedWeight),
		CabinAllowedHeight:  int32(bd.CabinAllowedHeight),
		HandAllowedLength:   int32(bd.HandAllowedLength),
		HandAllowedBreadth:  int32(bd.HandAllowedBreadth),
		HandAllowedWeight:   int32(bd.HandAllowedWeight),
		HandAllowedHeight:   int32(bd.HandAllowedHeight),
		FeeForExtraKgCabin:  int32(bd.FeeExtraPerKGCabin),
		FeeForExtraKgHand:   int32(bd.FeeExtraPerKGHand),
	}
	c1 := &pb.Cancellation{
		CancellationDeadlineBefore: int32(cd.CancellationDeadlineBefore),
		CancellationPercentage:     int32(cd.CancellationPercentage),
		Refundable:                 cd.Refundable,
	}
	sd1 := &pb.SearchFlightDetails{
		PathId:        int32(directFlight.FlightPath.PathId),
		NumberOfStops: int32(directFlight.FlightPath.NumberOfStops),
		FlightSegment: fd1,
	}
	df := &pb.Facilities{
		Cancellation: c1,
		Baggage:      bg1,
		Path:         sd1,
	}

	// Return Flight
	for _, f := range returnFlight.FlightPath.Flights {
		fare := f.EconomyFare
		if !economy {
			fare = f.BusinessFare
		}
		fd2 = append(fd2, &pb.FlightDetails{
			FlightNumber:     f.FlightNumber,
			Airline:          f.Airline,
			DepartureAirport: f.DepartureAirport,
			DepartureDate:    f.DepartureDate,
			DepartureTime:    f.DepartureTime,
			ArrivalAirport:   f.ArrivalAirport,
			ArrivalDate:      f.ArrivalDate,
			ArrivalTime:      f.ArrivalTime,
			FlightFare:       float32(fare),
		})
	}
	ca := returnFlight.Cancellation
	ba := returnFlight.Baggage
	bg2 := &pb.Baggage{
		CabinAllowedLength:  int32(ba.CabinAllowedLength),
		CabinAllowedBreadth: int32(ba.CabinAllowedBreadth),
		CabinAllowedWeight:  int32(ba.CabinAllowedWeight),
		CabinAllowedHeight:  int32(ba.CabinAllowedHeight),
		HandAllowedLength:   int32(ba.HandAllowedLength),
		HandAllowedBreadth:  int32(ba.HandAllowedBreadth),
		HandAllowedWeight:   int32(ba.HandAllowedWeight),
		HandAllowedHeight:   int32(ba.HandAllowedHeight),
		FeeForExtraKgCabin:  int32(ba.FeeExtraPerKGCabin),
		FeeForExtraKgHand:   int32(ba.FeeExtraPerKGHand),
	}
	c2 := &pb.Cancellation{
		CancellationDeadlineBefore: int32(ca.CancellationDeadlineBefore),
		CancellationPercentage:     int32(ca.CancellationPercentage),
		Refundable:                 ca.Refundable,
	}
	sd2 := &pb.SearchFlightDetails{
		PathId:        int32(returnFlight.FlightPath.PathId),
		NumberOfStops: int32(returnFlight.FlightPath.NumberOfStops),
		FlightSegment: fd2,
	}
	rf := &pb.Facilities{
		Cancellation: c2,
		Baggage:      bg2,
		Path:         sd2,
	}

	result := &pb.SearchSelectResponse{
		DirectFlight:     df,
		ReturnFlight:     rf,
		NumberOfAdults:   int32(cf.NumberOfAdults),
		NumberOfChildren: int32(cf.NumberOfChildren),
		CabinClass:       cf.CabinClass,
	}
	return result
}
