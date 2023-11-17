package interfaces

import pb "github.com/raedmajeed/booking-service/pkg/pb"

// dom "github.com/raedmajeed/booking-service/pkg/DOM"
// pb "github.com/raedmajeed/booking-service/pkg/pb"

type BookingService interface {
	SearchFlight(request *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error)
}
