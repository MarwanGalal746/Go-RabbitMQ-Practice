package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to the RabbitMQ instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume("RabbitMQ practice",
		"",
		true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	go func() {
		for d := range msgs {
			fmt.Println("Received msg: ", d.Body)
		}
	}()
	fmt.Println("Successfully consumed msg to the queue")

}
