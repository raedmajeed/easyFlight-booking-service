package handlers

import (
	"context"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"log"
)

func (handler *BookingHandler) RegisterSearchFlight(ctx context.Context, request *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	//fmt.Println(request)
	//var wg sync.WaitGroup
	//wg.Add(1)
	//flightResponse := make(chan *pb.SearchFlightResponse)
	searchResponse, err := handler.svc.SearchFlight(ctx, request)
	if err != nil {
		log.Println("unable to get the search details")
		return nil, err
	}
	//wg.Wait()
	//fmt.Println()
	return searchResponse, nil
}

func (handler *BookingHandler) RegisterSearchSelect(ctx context.Context, request *pb.SearchSelectRequest) (*pb.SearchSelectResponse, error) {
	response, err := handler.svc.SearchSelect(ctx, request)
	if err != nil {
		log.Println("unable to get the search details")
		return nil, err
	}
	return response, nil
}
func (handler *BookingHandler) RegisterTravellerDetails(ctx context.Context, request *pb.TravellerRequest) (*pb.TravellerResponse, error) {
	response, err := handler.svc.AddTraveller(ctx, request)
	if err != nil {
		log.Println("unable to get the search details")
		return nil, err
	}
	return response, nil
}
