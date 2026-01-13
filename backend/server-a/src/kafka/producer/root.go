package producer

import (
	"log/slog"
	"os"
	"server-a/config"
	"time"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.AsyncProducer
}

func NewKafkaProducer(cfg *config.Config) *KafkaProducer {
	producer, err := createProducer(cfg)
	if err != nil {
		slog.Error("fail to create producer",
			"err", err,
		)
		panic(err)
	}

	kp := KafkaProducer{producer}

	return &kp
}

func createProducer(config *config.Config) (sarama.AsyncProducer, error) {
	cfg := sarama.NewConfig()
	cfg.ClientID = config.Kafka.ProducerClientId
	cfg.Net.SASL.Enable = true
	cfg.Net.SASL.Version = 1
	cfg.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	//cfg.Net.SASL.User = <api-key>
	//cfg.Net.SASL.Password = <secret>
	cfg.Net.TLS.Enable = true
	cfg.Net.SASL.Handshake = true

	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Compression = sarama.CompressionZSTD
	cfg.Producer.RequiredAcks = sarama.WaitForLocal
	cfg.Producer.Retry.Max = 30
	cfg.Producer.Retry.Backoff = time.Millisecond * 10
	//cfg.Producer.Idempotent = true
	//cfg.Producer.RequiredAcks = sarama.WaitForAll
	//cfg.Net.MaxOpenRequests = 1

	return sarama.NewAsyncProducer([]string{os.Getenv("KAFKA_URL")}, cfg)
}

func (kp *KafkaProducer) PushMessage(topic string, message []byte) error {

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	kp.producer.Input() <- &msg

	select {
	case succeedMsg := <-kp.producer.Successes():
		slog.Info("Success to produce message",
			"partition", succeedMsg.Partition)
		return nil
	case err := <-kp.producer.Errors():
		slog.Error("Failed to produce message",
			"err", err)
		return err
	}
}

func (kp *KafkaProducer) Close() error {
	return kp.producer.Close()
}
