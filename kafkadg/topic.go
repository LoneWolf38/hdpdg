package main

type KafkaTopic struct {
	TopicName string
}

func NewKafkaTopic(config Config) *KafkaTopic {
	return &KafkaTopic{
		TopicName: config.TopicName,
	}
}

func (k *KafkaTopic) CreateTopic() {
}
