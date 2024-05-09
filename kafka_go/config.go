package main

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Comma Separate hostname:port string
	BrokerList string
	addrList   []string
	TopicName  string

	Async             bool
	AckMethod         string
	AutoTopicCreation bool
}

func readFromConfig(path string) (config Config, err error) {
	if _, err = os.Stat(path); err != nil {
		return config, err
	}
	config = Config{}
	fdata, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(fdata, &config)
	if err != nil {
		return config, err
	}

	config.addrList = append(config.addrList, strings.Split(config.BrokerList, ",")...)
	return config, nil
}
