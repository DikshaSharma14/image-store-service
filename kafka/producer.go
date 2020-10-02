package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"image-store-service/util"
	"log"
	"os"
	"time"
)

var producer sarama.SyncProducer

func init() {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	config := sarama.NewConfig()

	config.ClientID = "image-store"
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	prd, err := sarama.NewSyncProducer([]string{util.Config.KafkaUrl}, config)

	if err != nil {
		fmt.Println(err)
	}

	producer = prd
}

func Publish(message string) {
	msg := &sarama.ProducerMessage{
		Topic:     util.Config.KafkaTopic,
		Value:     sarama.StringEncoder(message),
		Timestamp: time.Now(),
	}
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}
}
