package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	runConsumer()
}

func runConsumer() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Got message with key %v, with offset %v and data: %s\n", msg.Key, msg.Offset, msg.Value)
		case err := <-partitionConsumer.Errors():
			fmt.Printf("Error: %v", err)
		}
	}
}
