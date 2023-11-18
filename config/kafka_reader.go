package config

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type KafkaReader2 struct {
	SearchReader *kafka.Reader
}

func NewKafkaReaderConnect() *KafkaReader2 {
	searchReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "search-flight-response",
		GroupID:  "search-response-1",
		MaxBytes: 10e1,
	})
	return &KafkaReader2{
		SearchReader: searchReader,
	}
}

func (kf *KafkaReader2) SearchReaderMethod(ctx context.Context, group *sync.WaitGroup, messageChan chan kafka.Message) {
	//defer group.Done()
	log.Println("kafka topic search-flight-response listening")
	//go func() {
	for {
		message, _ := kf.SearchReader.FetchMessage(ctx)
		fmt.Println(string(message.Value))
		select {
		case <-ctx.Done():
			log.Println("context cancelled returning")
			return
		default:
			err := kf.SearchReader.CommitMessages(ctx, message)
			if err != nil {
				return
			}
			log.Println("message++++", string(message.Value))
			break
		}
		break
	}
	fmt.Println("yahooo")
	//}()
	//return messageChan
}
