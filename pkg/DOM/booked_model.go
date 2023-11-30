package DOM

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	PNR              string `gorm:"unique"`
	Email            string
	Economy          bool
	PaymentId        string `gorm:"unique"`
	BookingReference string `gorm:"unique;not null"`
	BookingStatus    string `gorm:"default:PENDING"`
	DepartureAirport string
	ArrivalAirport   string
	FlightChartIDs   []byte
	TotalFare        string
	CancelledStatus  string      `gorm:"default:false"`
	UserId           uint        `gorm:"foreignKey:"`
	User             UserData    `gorm:"foreignKey:UserId"`
	Bookings         []Traveller `gorm:"many2many:traveller_booking;"`
}

type TravellerBooking struct {
	TravellerId uint
	BookingId   uint
}
