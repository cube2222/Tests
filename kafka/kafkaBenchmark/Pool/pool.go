package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"time"
	"sync"
)

func main() {
	producerPool := sync.Pool{New: func() interface{} {
		producer, _ :=  sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
		return producer
	}}

	fmt.Println("Starting sending")
	wg := sync.WaitGroup{}
	wg.Add(10)
	start := time.Now()
	for i := 0; i < 10; i++ {
		go func(num int) {
			producer := producerPool.Get()
			defer wg.Done()
			for i := 0 + num; i < 100000 + num; i++ {
				msg := sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder(fmt.Sprint(i))}
				producer.(sarama.AsyncProducer).Input() <- &msg
			}
		}(i)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		producer := producerPool.Get()
		producer.(sarama.AsyncProducer).Close()
	}
	end := time.Now().Sub(start)
	fmt.Printf("%v\n", end.Seconds())
}
