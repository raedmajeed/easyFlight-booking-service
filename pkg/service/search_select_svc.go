package service

import (
	"context"
	"encoding/json"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/utils"
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
	claims, err := utils.ValidateSearchToken(token, *svc.cfg)
	if err != nil {
		log.Println("search token is expired or null, please check")
		return nil, err
	}
	directPathID := request.DirectPathId
	returnPathID := request.ReturnPathId

	resp, err := svc.client.RegisterSelectFlight(ctx, &pb.SelectFlightAdmin{
		Token:        token,
		DirectPathId: directPathID,
		ReturnPathId: returnPathID,
		Adults:       int32(claims.Adults),
		Children:     int32(claims.Children),
		Economy:      claims.Economy,
	})

	response := ConvertToResponse(resp, claims.Economy)
	marshal, err := json.Marshal(resp)
	svc.redis.Set(ctx, token+"1", marshal, time.Minute*10)
	return response, nil
}

func ConvertToResponse(cf *pb.CompleteFlightDetails, economy bool) *pb.SearchSelectResponse {
	directFlight := cf.DirectFlight
	returnFlight := cf.ReturnFlight
	var fd1 []*pb.FlightDetail
	var fd2 []*pb.FlightDetail

	for _, f := range directFlight.FlightPath.FlightSegment {
		fd1 = append(fd1, &pb.FlightDetail{
			FlightNumber:     f.FlightNumber,
			Airline:          f.Airline,
			DepartureAirport: f.DepartureAirport,
			DepartureDate:    f.DepartureDate,
			DepartureTime:    f.DepartureTime,
			ArrivalAirport:   f.ArrivalAirport,
			ArrivalDate:      f.ArrivalDate,
			ArrivalTime:      f.ArrivalTime,
			FlightFare:       f.FlightFare,
			FlightChartId:    f.FlightChartId,
		})
	}

	cd := directFlight.Cancellation
	bd := directFlight.Baggage
	bg1 := &pb.BaggageBooking{
		CabinAllowedBreadth: int32(bd.CabinAllowedBreadth),
		CabinAllowedLength:  int32(bd.CabinAllowedLength),
		CabinAllowedWeight:  int32(bd.CabinAllowedWeight),
		CabinAllowedHeight:  int32(bd.CabinAllowedHeight),
		HandAllowedLength:   int32(bd.HandAllowedLength),
		HandAllowedBreadth:  int32(bd.HandAllowedBreadth),
		HandAllowedWeight:   int32(bd.HandAllowedWeight),
		HandAllowedHeight:   int32(bd.HandAllowedHeight),
		FeeForExtraKgCabin:  int32(bd.FeeExtraPerKgCabin),
		FeeForExtraKgHand:   int32(bd.FeeExtraPerKgHand),
	}
	c1 := &pb.CancellationBooking{
		CancellationDeadlineBefore: int32(cd.CancellationDeadlineBefore),
		CancellationPercentage:     int32(cd.CancellationPercentage),
		Refundable:                 cd.Refundable,
	}
	sd1 := &pb.SearchFlightDetail{
		PathId:        int32(directFlight.FlightPath.PathId),
		NumberOfStops: int32(directFlight.FlightPath.NumberOfStops),
		FlightSegment: fd1,
	}
	var fare float32
	if len(directFlight.FlightPath.FlightSegment) > 0 {
		fare = directFlight.FlightPath.FlightSegment[0].FlightFare
	}
	df := &pb.Facilities{
		Cancellation: c1,
		Baggage:      bg1,
		FlightPath:   sd1,
		Fare:         fare,
	}

	for _, f := range returnFlight.FlightPath.FlightSegment {
		fd2 = append(fd2, &pb.FlightDetail{
			FlightNumber:     f.FlightNumber,
			Airline:          f.Airline,
			DepartureAirport: f.DepartureAirport,
			DepartureDate:    f.DepartureDate,
			DepartureTime:    f.DepartureTime,
			ArrivalAirport:   f.ArrivalAirport,
			ArrivalDate:      f.ArrivalDate,
			ArrivalTime:      f.ArrivalTime,
			FlightFare:       f.FlightFare,
			FlightChartId:    f.FlightChartId,
		})
	}
	ca := returnFlight.Cancellation
	ba := returnFlight.Baggage
	bg2 := &pb.BaggageBooking{
		CabinAllowedLength:  int32(ba.CabinAllowedLength),
		CabinAllowedBreadth: int32(ba.CabinAllowedBreadth),
		CabinAllowedWeight:  int32(ba.CabinAllowedWeight),
		CabinAllowedHeight:  int32(ba.CabinAllowedHeight),
		HandAllowedLength:   int32(ba.HandAllowedLength),
		HandAllowedBreadth:  int32(ba.HandAllowedBreadth),
		HandAllowedWeight:   int32(ba.HandAllowedWeight),
		HandAllowedHeight:   int32(ba.HandAllowedHeight),
		FeeForExtraKgCabin:  int32(ba.FeeExtraPerKgCabin),
		FeeForExtraKgHand:   int32(ba.FeeExtraPerKgHand),
	}
	c2 := &pb.CancellationBooking{
		CancellationDeadlineBefore: int32(ca.CancellationDeadlineBefore),
		CancellationPercentage:     int32(ca.CancellationPercentage),
		Refundable:                 ca.Refundable,
	}
	sd2 := &pb.SearchFlightDetail{
		PathId:        int32(returnFlight.FlightPath.PathId),
		NumberOfStops: int32(returnFlight.FlightPath.NumberOfStops),
		FlightSegment: fd2,
	}
	if len(directFlight.FlightPath.FlightSegment) > 0 {
		fare = directFlight.FlightPath.FlightSegment[0].FlightFare
	}

	rf := &pb.Facilities{
		Cancellation: c2,
		Baggage:      bg2,
		FlightPath:   sd2,
		Fare:         fare,
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
