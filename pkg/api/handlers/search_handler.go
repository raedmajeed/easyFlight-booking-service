package handlers

import (
	"context"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
)

func (handler *BookingHandler) RegisterSearchFlight(ctx context.Context, request *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	searchResponse, err := handler.svc.SearchFlight(ctx, request)
	if err != nil {
		return nil, err
	}
	return searchResponse, nil
}

func (handler *BookingHandler) RegisterSearchSelect(ctx context.Context, request *pb.SearchSelectRequest) (*pb.SearchSelectResponse, error) {
	response, err := handler.svc.SearchSelect(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (handler *BookingHandler) RegisterTravellerDetails(ctx context.Context, request *pb.TravellerRequest) (*pb.TravellerResponse, error) {
	response, err := handler.svc.AddTraveller(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
