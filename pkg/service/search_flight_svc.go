package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/raedmajeed/booking-service/pkg/DOM"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/segmentio/kafka-go"
	"log"
)

func (svc *BookingServiceStruct) SearchFlight(ctx context.Context, request *pb.SearchFlightRequest) {
	//defer group.Done()/
	//messageChan := make(chan kafka.Message)
	//svc.kf2.SearchReaderMethod(ctx, group, messageChan)
	economy := true
	if request.Type == "1" {
		economy = false
	}
	returnType := false
	if request.ReturnDate != "" {
		returnType = true
	}
	searchDetails := DOM.SearchDetails{
		DepartureAirport:    request.FromAirport,
		ArrivalAirport:      request.ToAirport,
		DepartureDate:       request.DepartDate,
		ReturnDepartureDate: request.ReturnDate,
		ReturnFlight:        returnType,
		MaxStops:            request.MaxStops,
		Economy:             economy,
	}

	byteSearchData, err := json.Marshal(&searchDetails)
	if err != nil {
		log.Println("error marshaling data")
		return
	}

	var message kafka.Message
	err = svc.kf.SearchWriter.WriteMessages(ctx,
		kafka.Message{
			Value: byteSearchData,
		})
	if err != nil {
		log.Println("error writing to kafka: ", err.Error())
		return
	}

	fmt.Println("BACK TO SERVICE")
	//time.Sleep(time.Second * 2000)
	message = svc.kf2.SearchReaderMethod(ctx)
	var paths DOM.KafkaPath
	err = json.Unmarshal(message.Value, &paths)
	if err != nil {
		return
	}
	fmt.Println(paths.DirectPath)
	fmt.Println()
	fmt.Println(paths.ReturnPath)

	//time.Sleep(time.Second * 2000)
	//group.Add(1)
	//select {
	//case <-ctx.Done():
	//	log.Println("context cancelled")
	//	return
	//case message = <-messageChan:
	//	break
	//}
	//fmt.Println(string(message.Value), "message received")
}
