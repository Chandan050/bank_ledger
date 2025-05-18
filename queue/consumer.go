package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"banking_ledger/internal/models"
	"banking_ledger/internal/service"

	"github.com/IBM/sarama"
)

// ConsumeTransactions consumes messages from Kafka and processes transactions
func ConsumeTransactions(ctx context.Context, srv service.Service, topic string, partition int) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error while intializing the consumer %v", err)
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Error in consuming partition ", topic)
	}

	defer partitionConsumer.Close()

	for {
		log.Println("consuming the data")
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Println("got message", string(msg.Value))
			var tx models.Transaction
			err := json.Unmarshal(msg.Value, &tx)
			if err != nil {
				log.Printf("Failed to unmarshal transaction: %v", err)
				continue
			}

			_, err = srv.ProcessTransaction(ctx, &tx)
			if err != nil {
				log.Printf("Failed to process transaction: %v", err)
			} else {
				log.Printf("Processed transaction for account %s", tx.AccountID)
			}

		case err := <-partitionConsumer.Errors():
			log.Printf("Consumer error: %v", err)

			// case <-ctx.Done():
			// 	log.Println("Consumer shutting down")
			// 	return
		}
	}
}
