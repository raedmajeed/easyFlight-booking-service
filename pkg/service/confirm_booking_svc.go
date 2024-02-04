package service

import (
	"context"
	"encoding/json"
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
	var req pb.SearchSelectResponse
	if err := json.Unmarshal([]byte(val.Val()), &req); err != nil {
		log.Println("unable to unmarshal redis ConfirmBooking()")
		return nil, err
	}

	directFlight := req.DirectFlight.FlightPath.FlightSegment
	returnFlight := req.ReturnFlight.FlightPath.FlightSegment
	directStops := req.DirectFlight.FlightPath.NumberOfStops
	returnStops := req.ReturnFlight.FlightPath.NumberOfStops

	adults := req.NumberOfAdults
	children := req.NumberOfChildren
	cabinClass := req.CabinClass

	directCancellation := req.DirectFlight.Cancellation
	returnCancellation := req.ReturnFlight.Cancellation
	directBaggage := req.DirectFlight.Baggage
	returnBaggage := req.ReturnFlight.Baggage

	var fare float32
	var directFlightList []*pb.FlightDetail
	for _, flight := range directFlight {
		directFlightList = append(directFlightList, &pb.FlightDetail{
			FlightNumber:     flight.FlightNumber,
			Airline:          flight.Airline,
			DepartureAirport: flight.DepartureAirport,
			DepartureDate:    flight.DepartureDate,
			DepartureTime:    flight.DepartureTime,
			ArrivalAirport:   flight.ArrivalAirport,
			ArrivalDate:      flight.ArrivalDate,
			ArrivalTime:      flight.ArrivalTime,
			FlightChartId:    flight.FlightChartId,
		})
		fare = flight.FlightFare
	}
	directFlightFare := flightFareCalc(fare, adults, children)
	directCancel := &pb.CancellationBooking{
		CancellationPercentage:     int32(directCancellation.CancellationPercentage),
		CancellationDeadlineBefore: int32(directCancellation.CancellationDeadlineBefore),
		Refundable:                 directCancellation.Refundable,
	}

	directBag := &pb.BaggageBooking{
		CabinAllowedBreadth: int32(directBaggage.CabinAllowedBreadth),
		CabinAllowedHeight:  int32(directBaggage.CabinAllowedHeight),
		CabinAllowedWeight:  int32(directBaggage.CabinAllowedWeight),
		CabinAllowedLength:  int32(directBaggage.CabinAllowedLength),
		HandAllowedHeight:   int32(directBaggage.HandAllowedHeight),
		HandAllowedWeight:   int32(directBaggage.HandAllowedWeight),
		HandAllowedBreadth:  int32(directBaggage.HandAllowedBreadth),
		HandAllowedLength:   int32(directBaggage.HandAllowedLength),
		FeeForExtraKgHand:   int32(directBaggage.FeeForExtraKgHand),
		FeeForExtraKgCabin:  int32(directBaggage.FeeForExtraKgCabin),
		Restrictions:        directBaggage.Restrictions,
	}
	direct := &pb.DirectReturnFlight{
		FlightDetails: directFlightList,
		NumberOfStops: int32(directStops),
		Cancellation:  directCancel,
		Baggage:       directBag,
	}

	// Return Flight
	var returnFlightList []*pb.FlightDetail
	for _, flight := range returnFlight {
		returnFlightList = append(returnFlightList, &pb.FlightDetail{
			FlightChartId:    flight.FlightChartId,
			FlightNumber:     flight.FlightNumber,
			Airline:          flight.Airline,
			DepartureAirport: flight.DepartureAirport,
			DepartureDate:    flight.DepartureDate,
			DepartureTime:    flight.DepartureTime,
			ArrivalAirport:   flight.ArrivalAirport,
			ArrivalDate:      flight.ArrivalDate,
			ArrivalTime:      flight.ArrivalTime,
		})
		fare = flight.FlightFare
	}
	returnFlightFare := flightFareCalc(fare, adults, children)
	returnCancel := &pb.CancellationBooking{
		CancellationPercentage:     int32(returnCancellation.CancellationPercentage),
		CancellationDeadlineBefore: int32(returnCancellation.CancellationDeadlineBefore),
		Refundable:                 returnCancellation.Refundable,
	}

	returnBag := &pb.BaggageBooking{
		CabinAllowedBreadth: int32(returnBaggage.CabinAllowedBreadth),
		CabinAllowedHeight:  int32(returnBaggage.CabinAllowedHeight),
		CabinAllowedWeight:  int32(returnBaggage.CabinAllowedWeight),
		CabinAllowedLength:  int32(returnBaggage.CabinAllowedLength),
		HandAllowedHeight:   int32(returnBaggage.HandAllowedHeight),
		HandAllowedWeight:   int32(returnBaggage.HandAllowedWeight),
		HandAllowedBreadth:  int32(returnBaggage.HandAllowedBreadth),
		HandAllowedLength:   int32(returnBaggage.HandAllowedLength),
		FeeForExtraKgHand:   int32(returnBaggage.FeeForExtraKgHand),
		FeeForExtraKgCabin:  int32(returnBaggage.FeeForExtraKgCabin),
		Restrictions:        returnBaggage.Restrictions,
	}
	returnFlights := &pb.DirectReturnFlight{
		FlightDetails: returnFlightList,
		NumberOfStops: int32(returnStops),
		Cancellation:  returnCancel,
		Baggage:       returnBag,
	}

	completeDetails := &pb.ConfirmBookingResponse{
		DirectFlight:     direct,
		ReturnFlight:     returnFlights,
		NumberOfChildren: int32(children),
		NumberOfAdults:   adults,
		CabinClass:       cabinClass,
		BookingReference: bookingReference,
		DirectFlightFare: directFlightFare,
		ReturnFlightFare: returnFlightFare,
	}
	marshal, err := json.Marshal(completeDetails)
	if err != nil {
		log.Println("unable to marshal request")
		return nil, err
	}

	svc.redis.Set(ctx, bookingReference, marshal, time.Minute*10)
	return completeDetails, nil
}

func flightFareCalc(fare float32, adults, children int32) float32 {
	return (fare * float32(adults)) + (fare * float32(adults) * 0.8)
}
