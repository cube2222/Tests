package main

import (
	"github.com/Shopify/sarama"
	"log"
	"fmt"
	"bufio"
	"os"
)

func main() {
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	textSender := make(chan string)

	go func(producer *sarama.AsyncProducer) {
		for {
			select {
			case textToSend := <- textSender:
				msg := sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder(textToSend)}
				(*producer).Input() <- &msg
			case err := <- (*producer).Errors():
				fmt.Errorf("%v", err)
			}
		}
	}(&producer)

	consoleReader := bufio.NewReader(os.Stdin)

	for text := ""; err == nil; text, err = consoleReader.ReadString('\n') {
		textSender <- text
	}

}
