package services

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func Connection(user string, password string, url string) *amqp.Connection {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s/", user, password, url)

	conn, err := amqp.Dial(connectionString)

	if err != nil {
		log.Fatalln("Failed to connect to RabbitMQ: ", err.Error())
	}

	log.Println("Success connecting to RabbitMQ")

	return conn
}
