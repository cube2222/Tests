package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"bufio"
	"os"
)

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	consoleReader := bufio.NewReader(os.Stdin)

	for text := ""; err == nil; text, err = consoleReader.ReadString('\n') {
		_, _, err := producer.SendMessage(&sarama.ProducerMessage{Topic:"test", Value: sarama.StringEncoder(text)})
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}

	fmt.Printf("%v\n", err)
}
