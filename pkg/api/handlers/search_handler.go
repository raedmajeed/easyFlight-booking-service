package handlers

import (
	"context"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
)

func (handler *BookingHandler) RegisterSearchFlight(ctx context.Context, request *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	//fmt.Println(request)
	//var wg sync.WaitGroup
	//wg.Add(1)
	//flightResponse := make(chan *pb.SearchFlightResponse)
	handler.svc.SearchFlight(ctx, request)
	//wg.Wait()
	//fmt.Println()
	return nil, nil
}
