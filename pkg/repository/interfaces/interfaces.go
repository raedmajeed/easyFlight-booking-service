package interfaces

import (
	"context"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
)

//dom "github.com/raedmajeed/booking-service/pkg/DOM"
// pb "github.com/raedmajeed/booking-service/pkg/pb"

type BookingRepository interface {
	FindUserByEmail(email string) (*dom.UserData, error)
	CreateUser(airline *dom.UserData) (*dom.UserData, error)
	CreateBookedTravellers(booking *dom.Booking) error
	//CreateBookedTravellers(bookedTravellers *dom.BookedTravellers) error

	FindBookingByPNR(string) (*dom.Booking, error)
	FindNumberOfBookedUsers(ctx context.Context, string2 string) (*dom.Booking, error)

	FindNumberOfTravellers(uint2 uint) ([]*dom.TravellerBooking, error)
	FindTravellerById(uint2 uint) (*dom.Traveller, error)
	FindBooking(email, bookingReference string) (dom.Booking, error)

	UpdateBookings(booking dom.Booking) error
	FindPendingBooking() ([]*dom.Booking, error)
}
