package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"time"
	"sync"
	"log"
)

func main() {
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	//var successes int64 = 0

	fmt.Println("Starting sending")
	wg := sync.WaitGroup{}
	wg.Add(10)
	start := time.Now()
	for i := 0; i < 10; i++ {
		go func(num int, producer sarama.AsyncProducer) {
			defer wg.Done()
			for i := 0 + num; i < 100000 + num; i++ {
				msg := sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder(fmt.Sprint(i))}
				producer.Input() <- &msg
			}
		}(i, producer)
	}
	wg.Wait()
	producer.Close()
	end := time.Now().Sub(start)
	fmt.Printf("%v\n", end.Seconds())
}
