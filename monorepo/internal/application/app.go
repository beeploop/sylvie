package application

import (
	"sylvie/internal/http/controllers"
	"sylvie/internal/queue"

	"github.com/streadway/amqp"
)

type Application struct {
	RabbitConnection *amqp.Connection
	RabbitChannel    *amqp.Channel
	Publisher        queue.Publisher

	UploadController controllers.UploadController
}
