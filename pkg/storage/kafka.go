
package storage

import (
	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func InitKafka(brokers []string, topic string) {
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func GetKafkaWriter() *kafka.Writer {
	return kafkaWriter
}

func CloseKafka() error {
	return kafkaWriter.Close()
}