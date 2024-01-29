package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/utils"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
	"time"
)

// SearchFlight method is used to search flights in the routes
// as per user defined, also it checks if the flight is direct or indirect
// and generates a SearchToken which is required for subsequent requests
func (svc *BookingServiceStruct) SearchFlight(ctx context.Context, request *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	newContext, cancel := context.WithTimeout(ctx, time.Second*10000)
	defer cancel()

	economy := cabinClass(request.Type)
	returnType := returnStatus(request.Type)
	// don't need this
	resp, err := svc.client.RegisterSearchFlight(newContext, &pb.SearchFlightRequestAdmin{
		DepartureAirport:    request.FromAirport,
		ArrivalAirport:      request.ToAirport,
		DepartureDate:       request.DepartDate,
		ReturnDepartureDate: request.ReturnDate,
		ReturnFlight:        returnType,
		MaxStops:            request.MaxStops,
		Economy:             economy,
	})
	if err != nil {
		return nil, err
	}

	//byteSearchData, err := marshalSearch(request, returnType, economy)
	//err = writingToKafka(newContext, byteSearchData, svc)
	//message, err := readingFromKafka(newContext, svc)

	//var paths dom.KafkaPaths
	//err = json.Unmarshal(message.Value, &paths)
	//if err != nil {
	//	return nil, fmt.Errorf("error unmarshaling SearchFlight(), err: %v", err.Error())
	//}

	adults, _ := strconv.Atoi(request.Adults)
	children, _ := strconv.Atoi(request.Children)
	info := dom.AdditionalInfo{
		AdultsCount:   adults,
		ChildrenCount: children,
		PassengerType: request.PassengerType,
		Economy:       economy,
	}

	if len(resp.DirectPath) == 0 {
		return nil, fmt.Errorf("error fetching flight details from admin service")
	}

	token, err := utils.GenerateSearchToken(&info, svc.cfg)
	ToFlight := ConvertToSearchResponse(resp.DirectPath, economy)
	returnFlights := ConvertToSearchResponse(resp.ReturnPath, economy)

	searchFlightResponse := &pb.SearchFlightResponse{
		TotalDirectFlights: int32(len(ToFlight)),
		TotalReturnFlights: int32(len(returnFlights)),
		ToFlights:          ToFlight,
		ReturnFlights:      returnFlights,
		SearchToken:        token,
	}

	marshal, err := json.Marshal(&searchFlightResponse)

	err = AddToRedis(svc.redis, token, marshal, ctx)
	if err != nil {
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

func ConvertToSearchResponse(paths []*pb.Path, economy bool) []*pb.SearchFlightDetail {
	var flightDetails []*pb.FlightDetail
	var flightPaths []*pb.SearchFlightDetail

	for _, path := range paths {
		flightDetails = []*pb.FlightDetail{}
		for _, f := range path.Flights {
			fare := f.EconomyFare
			if !economy {
				fare = f.BusinessFare
			}
			flightDetail := pb.FlightDetail{
				FlightNumber:     f.FlightNumber,
				Airline:          f.Airline,
				DepartureAirport: f.DepartureAirport,
				DepartureDate:    f.DepartureDate,
				DepartureTime:    f.DepartureTime,
				ArrivalAirport:   f.ArrivalAirport,
				ArrivalDate:      f.ArrivalDate,
				ArrivalTime:      f.ArrivalTime,
				FlightFare:       float32(fare),
				FlightChartId:    f.FlightChartId,
			}
			flightDetails = append(flightDetails, &flightDetail)
		}
		flightPaths = append(flightPaths, &pb.SearchFlightDetail{
			PathId:        int32(path.PathId),
			NumberOfStops: int32(path.NumberOfStops - 1),
			FlightSegment: flightDetails,
		})
	}
	return flightPaths
}

func cabinClass(economyCheck string) bool {
	economy := true
	if economyCheck == "1" {
		economy = false
	}
	return economy
}

func returnStatus(returnVal string) bool {
	returnType := false
	if returnVal != "" {
		returnType = true
	}
	return returnType
}

func readingFromKafka(ctx context.Context, svc *BookingServiceStruct) (kafka.Message, error) {
	log.Println("waiting for search details from Airline Service")
	var message kafka.Message

	message = svc.kf2.SearchReaderMethod(ctx)
	if message.Value == nil {
		return kafka.Message{}, errors.New("message read from kafka is empty readingFromKafka()")
	}
	return message, nil
}

func writingToKafka(ctx context.Context, byteSearchData []byte, svc *BookingServiceStruct) error {
	log.Printf("writing message to kakfa -> message: %v", string(byteSearchData))
	err := svc.kf.SearchWriter.WriteMessages(ctx,
		kafka.Message{
			Value: byteSearchData,
		})
	if err != nil {
		return fmt.Errorf("error writing to kafka in writingToKafka() err: %v", err.Error())
	}
	return err
}

func marshalSearch(request *pb.SearchFlightRequest, returnType bool, economy bool) ([]byte, error) {
	searchDetails := dom.SearchDetails{
		DepartureAirport:    request.FromAirport,
		ArrivalAirport:      request.ToAirport,
		DepartureDate:       request.DepartDate,
		ReturnDepartureDate: request.ReturnDate,
		ReturnFlight:        returnType,
		MaxStops:            request.MaxStops,
		Economy:             economy,
	}

	// marshaling and sending the search details to admin-service
	byteSearchData, err := json.Marshal(&searchDetails)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json in marshalSearch(), err: %v", err.Error())
	}
	return byteSearchData, nil
}
