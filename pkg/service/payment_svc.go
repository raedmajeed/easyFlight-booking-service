package service

import (
	"context"
	"encoding/json"
	"fmt"
	dom "github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	razorpay "github.com/razorpay/razorpay-go"
	"log"
)

func (svc *BookingServiceStruct) OnlinePayment(ctx context.Context, request *pb.OnlinePaymentRequest) (*pb.OnlinePaymentResponse, error) {
	var flightDetails pb.ConfirmBookingResponse
	email := request.Email
	bookingReference := request.BookingReference
	val := svc.redis.Get(ctx, bookingReference).Val()
	err := json.Unmarshal([]byte(val), &flightDetails)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json err: %v", err.Error())
	}

	bookingDetails, err := svc.repo.FindBooking(email, bookingReference)
	if err != nil {
		return nil, fmt.Errorf("error finding booking details for user %v", email)
	}

	client := razorpay.NewClient(svc.cfg.RAZORPAYKEYID, svc.cfg.RAZORPAYSECRETKEY)
	directAmount := flightDetails.DirectFlightFare
	returnAmount := flightDetails.ReturnFlightFare
	fare := int(directAmount) + int(returnAmount)
	amountInPaise := int(fare) * 100

	data := map[string]interface{}{
		"amount":   amountInPaise,
		"currency": "INR",
		"receipt":  "bookingReference",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		log.Println("error creating order, err: ", err.Error())
		return nil, err
	}

	orderId := body["id"].(string)
	//type RazorpayDetails struct {
	//	UserID           uint
	//	TotalFare        int
	//	BookingReference string
	//	Email            string
	//	OrderID          string
	//}

	return &pb.OnlinePaymentResponse{
		UserId:           int32(bookingDetails.UserId),
		TotalFare:        int32(fare),
		BookingReference: bookingReference,
		Email:            email,
		OrderId:          orderId,
	}, nil
}

func (svc *BookingServiceStruct) PaymentConfirmed(ctx context.Context, request *pb.PaymentConfirmedRequest) (*pb.PaymentConfirmedResponse, error) {
	bookingReference := request.BookingReference
	email := request.Email
	bookingDetails, err := svc.repo.FindBooking(email, bookingReference)
	if err != nil {
		return nil, fmt.Errorf("error finding booking details for user %v", email)
	}

	val := svc.redis.Get(ctx, request.Token+"1")
	var req dom.CompleteFlightFacilities
	if err := json.Unmarshal([]byte(val.Val()), &req); err != nil {
		log.Println("unable to unmarshal redis ConfirmBooking()")
		return nil, err
	}

	economy := true
	if req.CabinClass != "ECONOMY" {
		economy = false
	}

	directFlightCharts := req.DirectFlight.FlightPath.Flights
	returnFlights := req.ReturnFlight.FlightPath.Flights
	var directFlights []int32
	for _, f := range directFlightCharts {
		directFlights = append(directFlights, int32(f.FlightChartID))
	}
	var returnFlight []int32
	for _, f := range returnFlights {
		returnFlight = append(returnFlight, int32(f.FlightChartID))
	}

	totalTravellers := req.NumberOfChildren + req.NumberOfAdults
	_, err = svc.client.AddConfirmedSeats(ctx, &pb.ConfirmedSeatRequest{
		Economy:               economy,
		FlightChartIdDirect:   directFlights,
		FlightChartIdIndirect: returnFlight,
		Travellers:            int32(totalTravellers),
	})

	if err != nil {
		return nil, err
	}

	bookingDetails.BookingStatus = "CONFIRMED"
	bookingDetails.PaymentId = request.PaymentId
	//bookingDetails.FlightChartIDs =
	err = svc.repo.UpdateBookings(bookingDetails)
	if err != nil {
		return nil, err
	}
	return &pb.PaymentConfirmedResponse{
		PaymentId: request.PaymentId,
		BookingId: bookingReference,
		PNR:       bookingDetails.PNR,
	}, err
}
