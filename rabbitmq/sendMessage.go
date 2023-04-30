package rabbitmq

import "github.com/hidekimva/golang/rabbitmq/services"

func SendMessage(user string, password string, url string, queueReturn bool, queueName string, queueRName string, msg interface{}) {
	conn := services.Connection(user, password, url)
	defer conn.Close()

	channel := services.CreateChannel(conn, queueReturn, queueRName, queueName)
	defer channel.Close()

	services.PublicSendQueue(channel, msg, queueReturn, queueRName, queueName)

}
