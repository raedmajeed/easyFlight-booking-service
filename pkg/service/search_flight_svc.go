package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/utils"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
	"time"
)

func (svc *BookingServiceStruct) SearchFlight(ctx context.Context, request *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	//defer group.Done()/
	//messageChan := make(chan kafka.Message)
	//svc.kf2.SearchReaderMethod(ctx, group, messageChan)
	economy := true
	if request.Type == "1" {
		economy = false
	}
	returnType := false
	if request.ReturnDate != "" {
		returnType = true
	}
	searchDetails := dom.SearchDetails{
		DepartureAirport:    request.FromAirport,
		ArrivalAirport:      request.ToAirport,
		DepartureDate:       request.DepartDate,
		ReturnDepartureDate: request.ReturnDate,
		ReturnFlight:        returnType,
		MaxStops:            request.MaxStops,
		Economy:             economy,
	}

	byteSearchData, err := json.Marshal(&searchDetails)
	if err != nil {
		log.Println("error marshaling data")
		return nil, err
	}

	var message kafka.Message
	err = svc.kf.SearchWriter.WriteMessages(ctx,
		kafka.Message{
			Value: byteSearchData,
		})
	if err != nil {
		log.Println("error writing to kafka: ", err.Error())
		return nil, err
	}

	message = svc.kf2.SearchReaderMethod(ctx)
	if message.Value == nil {
		log.Println("nothing in kafka message")
		return nil, err
	}

	var paths dom.KafkaPaths
	err = json.Unmarshal(message.Value, &paths)
	if err != nil {
		return nil, err
	}

	adults, _ := strconv.Atoi(request.Adults)
	children, _ := strconv.Atoi(request.Children)
	info := dom.AdditionalInfo{
		AdultsCount:   adults,
		ChildrenCount: children,
		PassengerType: request.PassengerType,
		Economy:       economy,
	}

	token, err := utils.GenerateSearchToken(&info, svc.cfg)
	ToFlight := ConvertToSearchResponse(paths.DirectPath)
	returnFlights := ConvertToSearchResponse(paths.ReturnPath)
	searchFlightResponse := &pb.SearchFlightResponse{
		TotalDirectFlights: int32(len(ToFlight)),
		TotalReturnFlights: int32(len(returnFlights)),
		ToFlights:          ToFlight,
		ReturnFlights:      returnFlights,
		SearchToken:        token,
	}

	err = AddToRedis(svc.redis, token, message.Value, ctx)
	if err != nil {
		log.Println("unable to push values to redis err:", err.Error())
		return nil, err
	}
	return searchFlightResponse, nil
}

func AddToRedis(redis *redis.Client, token string, message []byte, ctx context.Context) error {
	status := redis.Set(ctx, token, message, time.Minute*10)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func ConvertToSearchResponse(paths []dom.Path) []*pb.SearchFlightDetails {
	var flightDetails []*pb.FlightDetails
	var flightPaths []*pb.SearchFlightDetails

	for _, path := range paths {
		flightDetails = []*pb.FlightDetails{}
		for _, f := range path.Flights {
			flightDetail := pb.FlightDetails{
				FlightNumber:     f.FlightNumber,
				Airline:          f.Airline,
				DepartureAirport: f.DepartureAirport,
				DepartureDate:    f.DepartureDate,
				DepartureTime:    f.DepartureTime,
				ArrivalAirport:   f.ArrivalAirport,
				ArrivalDate:      f.ArrivalDate,
				ArrivalTime:      f.ArrivalTime,
			}
			flightDetails = append(flightDetails, &flightDetail)
		}
		flightPaths = append(flightPaths, &pb.SearchFlightDetails{
			PathId:        int32(path.PathId),
			NumberOfStops: int32(path.NumberOfStops - 1),
			FlightSegment: flightDetails,
		})
	}
	return flightPaths
}
