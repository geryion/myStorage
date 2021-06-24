package rabbitmq

import "github.com/streadway/amqp"

type RabbitMQ struct {
	Channel 	*amqp.Channel
	Name		string
	Exchange 	string
}

