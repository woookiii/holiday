package logger

import (
	"log/slog"
	"server-a/src/kafka/producer"
)

func NewLogger(kafkaProducer *producer.KafkaProducer) {
	h := &Handler{kafkaProducer}
	l := slog.New(h)
	slog.SetDefault(l)
}
