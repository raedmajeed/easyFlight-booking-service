package interfaces

import (
	"context"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
)

// dom "github.com/raedmajeed/booking-service/pkg/DOM"
// pb "github.com/raedmajeed/booking-service/pkg/pb"

type BookingService interface {
	SearchFlight(context.Context, *pb.SearchFlightRequest)
}
