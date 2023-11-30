package DOM

import "gorm.io/gorm"

type UserData struct {
	gorm.Model
	Email    string
	Phone    string
	Password string
	Name     string
}

type Traveller struct {
	gorm.Model
	Name          string
	Age           string
	Gender        string
	SeatNo        string
	VegMealOption bool
	CheckedIn     bool
	UserId        uint
	User          UserData  `gorm:"foreignKey:UserId"`
	Booking       []Booking `gorm:"many2many:traveller_booking;"`
	Economy       bool      `gorm:"default:false"`
}

type ConfirmedUserLogin struct {
	Email string
	PNR   string
}

type PNRLoginResponse struct {
	Email string
	PNR   string
	Token string
}
