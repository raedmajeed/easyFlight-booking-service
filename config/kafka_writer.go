package config

import (
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	SearchWriter       *kafka.Writer
	SearchSelectWriter *kafka.Writer
	EmailWriter        *kafka.Writer
}

func NewKafkaWriterConnect() *KafkaWriter {
	emailWriter := kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "email-service",
		AllowAutoTopicCreation: true,
	}
	searchWriter := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "search-flight-bs-6",
		AllowAutoTopicCreation: true,
	}
	searchSelectWriter := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "search-select-flight-bs-6",
		AllowAutoTopicCreation: true,
	}
	return &KafkaWriter{
		EmailWriter:        &emailWriter,
		SearchWriter:       searchWriter,
		SearchSelectWriter: searchSelectWriter,
	}
}
