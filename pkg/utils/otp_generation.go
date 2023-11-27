package utils

import (
	"math/rand"
	"time"
)

func GenerateOTP() int {
	currentTime := time.Now()
	minutes := currentTime.Minute()
	seconds := currentTime.Second()

	// Generate a random seed for the remaining four digits
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomPart := r.Intn(10000) // Generates a random 4-digit number

	// Combine minutes, seconds, and the random part
	otp := minutes*10000 + seconds*100 + randomPart

	return otp
}
