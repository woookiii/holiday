package logger

import (
	"context"
	"encoding/json"
	"log/slog"
	"server-a/src/kafka/producer"
)

type Handler struct {
	kafkaProducer *producer.KafkaProducer
}

func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (h *Handler) Handle(_ context.Context, record slog.Record) error {
	msg, err := json.Marshal(record)
	if err != nil {
		slog.Error("fail to marshal log record to byte",
			"err", err,
		)
		return err
	}

	return h.kafkaProducer.PushMessage("log", msg)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return h
}
