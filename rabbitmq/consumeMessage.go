package rabbitmq

import (
	"log"

	"github.com/hidekimva/golang/rabbitmq/services"
	"github.com/streadway/amqp"
)

func ConsumeMessageQueue(user string, password string, url string, queueName string) amqp.Delivery {
	var queueRname = ""
	conn := services.Connection(user, password, url)
	defer conn.Close()

	channel := services.CreateChannel(conn, false, queueRname, queueName)
	defer channel.Close()

	// Consome mensagens da fila
	for {
		msgs, err := channel.Consume(
			queueName,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("Erro ao consumir fila: %s", err)
		}

		// Loop infinito para processar as mensagens recebidas
		for msg := range msgs {
			return msg
		}

	}

}
