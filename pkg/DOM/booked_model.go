package DOM

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	PNR                  string `gorm:"unique"`
	Email                string
	Economy              bool
	PaymentId            string `gorm:""`
	BookingReference     string `gorm:"uniques;not null"`
	BookingStatus        string `gorm:"default:PENDING"`
	DepartureAirport     string
	ArrivalAirport       string
	DirectFlightChartIds []byte `gorm:"type:longtext"`
	ReturnFlightChartIds []byte `gorm:"type:longtext"`
	TotalFare            string
	CancelledStatus      string      `gorm:"default:false"`
	UserId               uint        `gorm:"foreignKey:"`
	User                 UserData    `gorm:"foreignKey:UserId"`
	Bookings             []Traveller `gorm:"many2many:traveller_booking;"`
}

type TravellerBooking struct {
	TravellerId uint
	BookingId   uint
}
