package rabbitmq

import (
	"log"

	"github.com/beeploop/sylvie/internal/config"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	Msgs <-chan amqp.Delivery
}

func Init(cfg *config.Config) *RabbitMQ {
	rabbit := &RabbitMQ{}

	conn, err := amqp.Dial(cfg.RabbitConnectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
	rabbit.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	rabbit.ch = ch

	if _, err := ch.QueueDeclare(
		cfg.QueueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Fatal(err.Error())
	}

	msgs, err := ch.Consume(
		cfg.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	rabbit.Msgs = msgs

	return rabbit
}

func (r *RabbitMQ) Close() {
	r.ch.Close()
	r.conn.Close()
}
