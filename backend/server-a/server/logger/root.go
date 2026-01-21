package logger

import (
	"log"
	"log/slog"
	"os"
	"server-a/server/kafka/producer"
)

func SetLogger(kafkaProducer *producer.KafkaProducer) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	log.Print("success to set logger")
}
