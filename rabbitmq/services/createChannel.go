package services

import (
	"log"

	"github.com/streadway/amqp"
)

func CreateChannel(conn *amqp.Connection, queueR bool, queueRName string, queueName string) *amqp.Channel {
	// Cria um canal para se comunicar com o RabbitMQ
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalln("Error creating channel with rabbitmq: ", err.Error())
	}

	// Declara fila de retorno
	if queueR == true {
		_, err := channel.QueueDeclare(
			queueRName,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalln("Failed to declare return queue: ", err.Error())
		}

		// Declara fila de envio
		_, err2 := channel.QueueDeclare(
			queueName,
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			log.Fatalln("Failed to declare send queue: ", err2.Error())
		}

	} else {

		// Declara fila de envio
		_, err := channel.QueueDeclare(
			queueName,
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			log.Fatalln("Failed to declare send queue: ", err.Error())
		}
	}

	return channel

}
