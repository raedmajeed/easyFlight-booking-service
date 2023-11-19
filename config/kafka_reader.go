package config

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaReader2 struct {
	SearchReader *kafka.Reader
}

func NewKafkaReaderConnect() *KafkaReader2 {
	searchReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "search-flight-response-1",
		GroupID:  "search-response-1",
		MaxBytes: 10e1,
	})
	return &KafkaReader2{
		SearchReader: searchReader,
	}
}

func (kf *KafkaReader2) SearchReaderMethod(ctx context.Context) kafka.Message {
	//defer group.Done()
	log.Println("kafka topic search-flight-response listening")
	//go func() {
	var message kafka.Message
	for {
		message, _ = kf.SearchReader.FetchMessage(ctx)
		select {
		case <-ctx.Done():
			log.Println("context cancelled returning")
			return kafka.Message{}
		default:
			err := kf.SearchReader.CommitMessages(ctx, message)
			if err != nil {
				return kafka.Message{}
			}
			//break
		}
		break
	}
	return message
	//}()
	//return messageChan
}
