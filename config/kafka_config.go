package config

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaConfigStruct struct {
	SearchWriter *kafka.Writer
	SearchReader *kafka.Reader
}

func NewKafkaConfig() *KafkaConfigStruct {
	searchWriter := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "search-flight-request",
	}
	searchReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "search-flight-response",
		GroupID: "search-response-1",
	})
	return &KafkaConfigStruct{
		SearchWriter: searchWriter,
		SearchReader: searchReader,
	}
}

func (kf *KafkaConfigStruct) SearchReaderMethod(ctx context.Context) {
	log.Println("kafka topic search-flight-response listening")
	messageChan := make(chan kafka.Message)
	for {
		message, _ := kf.SearchReader.FetchMessage(ctx)
		select {
		case <-ctx.Done():
			log.Println("context cancelled returning")
			return
		case messageChan <- message:
			log.Println("message", string(message.Value))
		default:
		}
	}
}
