package main

import (
	"github.com/streadway/amqp"
	"fmt"
	"os"
	"log"
)

func main() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%v/", os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	err = ch.ExchangeDeclare("myeventexchange", "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, key := range os.Args[2:] {
		err = ch.QueueBind(q.Name, key, "myeventexchange", false, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	msgChan, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgChan {
		go func(msg amqp.Delivery) {
			fmt.Printf("Received message: %s\n", msg.Body)
			err := msg.Ack(false)
			if err != nil {
				fmt.Printf("Error when acknowledging: %v", err)
			}
		}(msg)
	}
}
