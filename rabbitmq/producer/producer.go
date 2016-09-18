package main

import (
	"github.com/streadway/amqp"
	"fmt"
	"os"
	"log"
	"time"
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

	body := "Hello World!!!"

	ticker := time.Tick(time.Millisecond)

	n := 0

	for range ticker {
		err := ch.Publish("", "test", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(fmt.Sprintf("Message number %v: %s", n, body)),
		})
		if err != nil {
			log.Printf("Error when publishing: %v", err)
		}
		fmt.Printf("Sent message number %v\n", n)
		n++
	}
}
