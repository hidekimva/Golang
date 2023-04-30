package rabbitmq

import (
	"log"

	"github.com/hidekimva/golang/rabbitmq/services"
	"github.com/streadway/amqp"
)

func SendResponseMessage(user string, password string, url string, body []byte, msg amqp.Delivery) {

	conn := services.Connection(user, password, url)
	defer conn.Close()

	channel := services.CreateChannel(conn, false, "", "")
	defer channel.Close()

	// Enviar a resposta para a fila de retorno
	err := channel.Publish(
		"",
		msg.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			Body:          body,
			CorrelationId: msg.CorrelationId,
		},
	)
	if err != nil {
		log.Fatalf("Falha ao enviar resposta: %v", err)
	}
}
