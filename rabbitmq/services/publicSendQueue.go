package services

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func PublicSendQueue(channel *amqp.Channel, msg interface{}, queueReturn bool, queueRName string, queueName string) {
	// Codifica body para o envio
	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln("Error encoding json: ", err.Error())
	}

	if queueReturn == true {
		// Publica na fila de envio
		err = channel.Publish(
			"",
			queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
				ReplyTo:     queueRName,
			},
		)
		if err != nil {
			log.Fatalln("Failed to send to queue: " + err.Error())
		} else {
			log.Println("Sent in queue successfully: ", body)
		}
	} else {
		// Publica na fila de envio
		err = channel.Publish(
			"",
			queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			log.Fatalln("Failed to send to queue: " + err.Error())
		} else {
			log.Println("Sent in queue successfully: ", body)
		}
	}

}
