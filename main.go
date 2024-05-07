package main

import (
	"fmt"
	"os"

	dg "github.com/LoneWolf38/hdpdg/datagen"

	"github.com/integrii/flaggy"
	"github.com/segmentio/kafka-go"
)

var (
	configPath  string = "./config.yml"
	iotDataPath        = "./datagen/data/iot_sample.csv"
)

func init() {
	flaggy.SetName("hdpdg")
	flaggy.SetDescription("Hadoop Services Datagenerator")
	flaggy.SetVersion("0.0.1")
	flaggy.String(&configPath, "c", "config", "path to config file")
	flaggy.Parse()
}

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
