package main

import (
	"context"
	"errors"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	TopicName     string
	BrokerList    []string
	kafkaProducer *kafka.Writer
}

func NewKafkaProducer(config Config) *KafkaProducer {
	return &KafkaProducer{
		TopicName:  config.TopicName,
		BrokerList: config.addrList,
		kafkaProducer: &kafka.Writer{
			Addr:                   kafka.TCP(config.addrList...),
			Topic:                  config.TopicName,
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: config.AutoTopicCreation,
		},
	}
}

func (k *KafkaProducer) Produce(messages []kafka.Message) {
	var err error
	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()

		// attempt to create topic prior to publishing the message
		err = k.kafkaProducer.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.Error("unexpected error %v", err)
		}
		break
	}

	if err := k.kafkaProducer.Close(); err != nil {
		log.Error("failed to close writer:", err)
	}
}
