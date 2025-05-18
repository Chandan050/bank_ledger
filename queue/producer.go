package queue

import (
	"banking_ledger/internal/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

func StartProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 3

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error while intializing the kafka producer")
	}
	return producer
}

func SendMesage(topic string, tx *models.Transaction) error {
	data, err := json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("Error while marshaling the data %v", err)
	}
	message := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	partiation, offset, err := producer.SendMessage(&message)
	if err != nil {
		return err
	}

	log.Printf("partiation %d and offset %d for sent message %d", partiation, offset, tx.AccountID)

	return nil
}
