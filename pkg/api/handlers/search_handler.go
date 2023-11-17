package handlers

import (
	"context"
	"fmt"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
)

func (h *BookingHandler) RegisterSearchFlight(ctx context.Context, request *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	response, _ := h.svc.SearchFlight(request)
	fmt.Println(response)
	return nil, nil
}
