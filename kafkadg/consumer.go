package main

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	ConsumerGroupID string
	TopicName       string
	BrokerList      []string
	kafkaConsumer   *kafka.Reader
}

func NewKafkaConsumer(config Config) *KafkaConsumer {
	kafkaConsumer := kafka.NewReader(kafka.ReaderConfig{
		GroupID:  config.ConsumerGroupID,
		Topic:    config.TopicName,
		Brokers:  config.addrList,
		MaxWait:  100 * time.Millisecond,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return &KafkaConsumer{
		ConsumerGroupID: config.ConsumerGroupID,
		TopicName:       config.TopicName,
		BrokerList:      config.addrList,
		kafkaConsumer:   kafkaConsumer,
	}
}

func (k *KafkaConsumer) Consume() {
}
