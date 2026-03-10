package application

import (
	"sylvie/internal/queue"

	"github.com/streadway/amqp"
)

type Application struct {
	RabbitConnection *amqp.Connection
	RabbitChannel    *amqp.Channel
	Publisher        queue.Publisher
}
