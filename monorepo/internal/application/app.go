package application

import (
	"sylvie/internal/queue"
	"sylvie/internal/router/controllers"

	"github.com/streadway/amqp"
)

type Application struct {
	RabbitConnection *amqp.Connection
	RabbitChannel    *amqp.Channel
	Publisher        queue.Publisher

	UploadController controllers.UploadController
}
