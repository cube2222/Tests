package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"log"
	"time"
)

func main() {
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	//var successes int64 = 0

	fmt.Println("Starting sending")

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		msg := sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder(fmt.Sprint(i))}
		producer.Input() <- &msg

	}
	producer.Close()
	end := time.Now().Sub(start)
	fmt.Printf("%v\n", end.Seconds())
}
