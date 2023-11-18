package config

import (
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	SearchWriter *kafka.Writer
}

func NewKafkaWriterConnect() *KafkaWriter {
	searchWriter := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "search-flight-request",
		Async:        true,
		RequiredAcks: 0,
	}
	return &KafkaWriter{
		SearchWriter: searchWriter,
	}
}
