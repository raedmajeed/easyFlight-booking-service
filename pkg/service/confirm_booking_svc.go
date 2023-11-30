package service

import (
	"context"
	"encoding/json"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"log"
	"time"
)

func (svc *BookingServiceStruct) ConfirmBooking(ctx context.Context, p *pb.ConfirmBookingRequest) (*pb.ConfirmBookingResponse, error) {
	token := p.Token
	email := p.Email
	bookingReference := p.BookingReference
	data, err := svc.repo.FindBooking(email, bookingReference)
	if err != nil {
		log.Println("booking not happened, add travellers")
		return nil, err
	}

	if data.BookingStatus == "CONFIRMED" {
		log.Println("booking already confirmed")
		return nil, err
	}

	val := svc.redis.Get(ctx, token+"1")
	var req dom.CompleteFlightFacilities
	if err := json.Unmarshal([]byte(val.Val()), &req); err != nil {
		log.Println("unable to unmarshal redis ConfirmBooking()")
		return nil, err
	}

	directFlight := req.DirectFlight.FlightPath.Flights
	returnFlight := req.ReturnFlight.FlightPath.Flights
	directStops := req.DirectFlight.FlightPath.NumberOfStops
	returnStops := req.ReturnFlight.FlightPath.NumberOfStops
	adults := req.NumberOfAdults
	//children := req.NumberOfChildren
	cabinClass := req.CabinClass
	directCancellation := req.DirectFlight.Cancellation
	returnCancellation := req.ReturnFlight.Cancellation
	directBaggage := req.DirectFlight.Baggage
	returnBaggage := req.ReturnFlight.Baggage
	economy := req.CabinClass

	var directFlightList []*pb.FlightDetails
	for _, flight := range directFlight {
		fare := flight.EconomyFare
		if economy == "BUSINESS" {
			fare = flight.BusinessFare
		}
		directFlightList = append(directFlightList, &pb.FlightDetails{
			FlightNumber:     flight.FlightNumber,
			Airline:          flight.Airline,
			DepartureAirport: flight.DepartureAirport,
			DepartureDate:    flight.DepartureDate,
			DepartureTime:    flight.DepartureTime,
			ArrivalAirport:   flight.ArrivalAirport,
			ArrivalDate:      flight.ArrivalDate,
			ArrivalTime:      flight.ArrivalTime,
			FlightFare:       float32(fare),
		})
	}

	directCancel := &pb.Cancellation{
		CancellationPercentage:     int32(directCancellation.CancellationPercentage),
		CancellationDeadlineBefore: int32(directCancellation.CancellationDeadlineBefore),
		Refundable:                 directCancellation.Refundable,
	}

	directBag := &pb.Baggage{
		CabinAllowedBreadth: int32(directBaggage.CabinAllowedBreadth),
		CabinAllowedHeight:  int32(directBaggage.CabinAllowedHeight),
		CabinAllowedWeight:  int32(directBaggage.CabinAllowedWeight),
		CabinAllowedLength:  int32(directBaggage.CabinAllowedLength),
		HandAllowedHeight:   int32(directBaggage.HandAllowedHeight),
		HandAllowedWeight:   int32(directBaggage.HandAllowedWeight),
		HandAllowedBreadth:  int32(directBaggage.HandAllowedBreadth),
		HandAllowedLength:   int32(directBaggage.HandAllowedLength),
		FeeForExtraKgHand:   int32(directBaggage.FeeExtraPerKGHand),
		FeeForExtraKgCabin:  int32(directBaggage.FeeExtraPerKGCabin),
		Restrictions:        directBaggage.Restrictions,
	}
	direct := &pb.DirectReturnFlight{
		FlightDetails: directFlightList,
		NumberOfStops: int32(directStops),
		Cancellation:  directCancel,
		Baggage:       directBag,
	}

	var returnFlightList []*pb.FlightDetails
	for _, flight := range returnFlight {
		fare := flight.EconomyFare
		if economy == "BUSINESS" {
			fare = flight.BusinessFare
		}
		returnFlightList = append(returnFlightList, &pb.FlightDetails{
			FlightNumber:     flight.FlightNumber,
			Airline:          flight.Airline,
			DepartureAirport: flight.DepartureAirport,
			DepartureDate:    flight.DepartureDate,
			DepartureTime:    flight.DepartureTime,
			ArrivalAirport:   flight.ArrivalAirport,
			ArrivalDate:      flight.ArrivalDate,
			ArrivalTime:      flight.ArrivalTime,
			FlightFare:       float32(fare),
		})
	}

	returnCancel := &pb.Cancellation{
		CancellationPercentage:     int32(returnCancellation.CancellationPercentage),
		CancellationDeadlineBefore: int32(returnCancellation.CancellationDeadlineBefore),
		Refundable:                 returnCancellation.Refundable,
	}

	returnBag := &pb.Baggage{
		CabinAllowedBreadth: int32(returnBaggage.CabinAllowedBreadth),
		CabinAllowedHeight:  int32(returnBaggage.CabinAllowedHeight),
		CabinAllowedWeight:  int32(returnBaggage.CabinAllowedWeight),
		CabinAllowedLength:  int32(returnBaggage.CabinAllowedLength),
		HandAllowedHeight:   int32(returnBaggage.HandAllowedHeight),
		HandAllowedWeight:   int32(returnBaggage.HandAllowedWeight),
		HandAllowedBreadth:  int32(returnBaggage.HandAllowedBreadth),
		HandAllowedLength:   int32(returnBaggage.HandAllowedLength),
		FeeForExtraKgHand:   int32(returnBaggage.FeeExtraPerKGHand),
		FeeForExtraKgCabin:  int32(returnBaggage.FeeExtraPerKGCabin),
		Restrictions:        returnBaggage.Restrictions,
	}
	returnFlights := &pb.DirectReturnFlight{
		FlightDetails: returnFlightList,
		NumberOfStops: int32(returnStops),
		Cancellation:  returnCancel,
		Baggage:       returnBag,
	}

	marshal, err := json.Marshal(req)
	if err != nil {
		log.Println("unable to marshal request")
		return nil, err
	}

	svc.redis.Set(ctx, bookingReference, marshal, time.Minute*10)
	return &pb.ConfirmBookingResponse{
		DirectFlight:     direct,
		ReturnFlight:     returnFlights,
		NumberOfChildren: int32(adults),
		CabinClass:       cabinClass,
		BookingReference: bookingReference,
	}, nil

}
