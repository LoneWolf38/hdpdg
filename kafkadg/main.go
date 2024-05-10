package main

import (
	"fmt"
	"os"

	dg "kafkadg/datagen"

	"github.com/LoneWol38/kafka-go"
)

var (
	configPath  string = "./config.yml"
	iotDataPath        = "./datagen/data/iot_sample.csv"
)

func main() {
	InitLogger()
	ch := make(chan string, 10000)
	dg.ReadCSV(iotDataPath, ch)
	cfg, err := readFromConfig(configPath)
	if err != nil {
		log.Error("failed to read config: ", err)
		os.Exit(1)
	}
	w := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.addrList...),
		Topic:                  cfg.TopicName,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: cfg.AutoTopicCreation,
		Logger:                 kafka.LoggerFunc(kafkaLogger),
	}
	fmt.Println(w)
	// writeToKafka(config)
}
