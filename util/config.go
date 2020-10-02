package util

import (
	"github.com/magiconair/properties"
	"log"
)

type Configuration struct {
	KafkaUrl   string `properties:"kafkaUrl"`
	KafkaTopic string `properties:"kafkaTopic"`
}

var Config Configuration

func init() {
	prop := properties.MustLoadFile("config.properties", properties.UTF8)
	if err := prop.Decode(&Config); err != nil {
		log.Fatal(err)
	}
}
