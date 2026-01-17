package logger

import (
	"log/slog"
	"server-a/server/kafka/producer"
)

func SetLogger(kafkaProducer *producer.KafkaProducer) {
	h := &Handler{kafkaProducer}
	l := slog.New(h)
	slog.SetDefault(l)
}
