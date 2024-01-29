package config

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaReader2 struct {
	SearchReader       *kafka.Reader
	SearchSelectReader *kafka.Reader
}

func NewKafkaReaderConnect() *KafkaReader2 {
	searchReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "search-flight-response-3",
		GroupID:  "search-response-3",
		MaxBytes: 10e1,
	})
	searchSelectReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "selected-flight-response-2",
		GroupID: "flight-response-2",
	})

	return &KafkaReader2{
		SearchReader:       searchReader,
		SearchSelectReader: searchSelectReader,
	}
}

func (kf *KafkaReader2) SearchReaderMethod(ctx context.Context) kafka.Message {
	log.Println("kafka topic search-flight-response listening")
	ch := make(chan kafka.Message)
	returnMsg := make(chan kafka.Message)
	message, _ := kf.SearchReader.ReadMessage(ctx)
	ch <- message
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("context cancelled returning")
				returnMsg <- kafka.Message{}
			case msg := <-ch:
				returnMsg <- msg
			}
		}
	}()
	return <-returnMsg
}

func (kf *KafkaReader2) SearchSelectReaderMethod(ctx context.Context) kafka.Message {
	//defer group.Done()
	log.Println("kafka topic SEARCH-FLIGHT-RESPONSE listening")
	//go func() {
	var message kafka.Message
	for {
		message, _ = kf.SearchSelectReader.ReadMessage(ctx)
		select {
		case <-ctx.Done():
			log.Println("context cancelled returning")
			return kafka.Message{}
		default:
			log.Println("go the message yesss select!!!")
			return message
			//err := kf.SearchSelectReader.CommitMessages(ctx, message)
			//if err != nil {
			//	return kafka.Message{}
			//}
			//if message.Value != nil {
			//	log.Println("received message search reader method")
			//	return message
			//}
		}
	}
	//}()
	//return messageChan
}
