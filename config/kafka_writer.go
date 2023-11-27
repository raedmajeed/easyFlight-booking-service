package config

import (
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	SearchWriter       *kafka.Writer
	SearchSelectWriter *kafka.Writer
}

func NewKafkaWriterConnect() *KafkaWriter {
	searchWriter := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "search-flight-request",
	}
	searchSelectWriter := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "search-flight-request-4",
	}
	return &KafkaWriter{
		SearchWriter:       searchWriter,
		SearchSelectWriter: searchSelectWriter,
	}
}
