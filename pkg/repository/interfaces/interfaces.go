package interfaces

import dom "github.com/raedmajeed/booking-service/pkg/DOM"

//dom "github.com/raedmajeed/booking-service/pkg/DOM"
// pb "github.com/raedmajeed/booking-service/pkg/pb"

type BookingRepository interface {
	FindUserByEmail(email string) (*dom.UserData, error)
	CreateUser(airline *dom.UserData) (*dom.UserData, error)
	CreateBookedTravellers(booking *dom.Booking) error
	//CreateBookedTravellers(bookedTravellers *dom.BookedTravellers) error
}
