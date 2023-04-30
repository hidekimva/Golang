package services

import (
	"log"

	"github.com/streadway/amqp"
)

func ConsumeMessageQueue(channel *amqp.Channel, queueRName string) []byte {
	// Consome mensagens da fila
	for {
		msgs, err := channel.Consume(
			queueRName,
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
			return msg.Body
		}

		return nil
	}

}
