package interfaces

import (
	"context"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
)

// dom "github.com/raedmajeed/booking-service/pkg/DOM"
// pb "github.com/raedmajeed/booking-service/pkg/pb"

type BookingService interface {
	SearchFlight(ctx context.Context, request *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error)
	SearchSelect(ctx context.Context, request *pb.SearchSelectRequest) (*pb.SearchSelectResponse, error)
	AddTraveller(ctx context.Context, request *pb.TravellerRequest) (*pb.TravellerResponse, error)

	UserLogin(p *pb.LoginRequest) (string, error)
	RegisterUserSvc(p *pb.UserRequest) (*dom.UserData, error)
	VerifyUserRequest(p *pb.OTPRequest) (*dom.UserData, error)

	ConfirmedUserLogin(ctx context.Context, p *pb.PNRLoginRequest) (*pb.PNRLoginResponse, error)
	SelectSeat(context.Context, *pb.SeatSelectRequest) (*pb.SeatSelectResponse, error)

	ConfirmBooking(context.Context, *pb.ConfirmBookingRequest) (*pb.ConfirmBookingResponse, error)
	OnlinePayment(context.Context, *pb.OnlinePaymentRequest) (*pb.OnlinePaymentResponse, error)
	PaymentConfirmed(context.Context, *pb.PaymentConfirmedRequest) (*pb.PaymentConfirmedResponse, error)
	CheckAndSendBookingReminder()
}
