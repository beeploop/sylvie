package rabbitmq

import (
	"log"

	"github.com/beeploop/sylvie/internal/config"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn          *amqp.Connection
	consumerChan  *amqp.Channel
	publisherChan *amqp.Channel
}

func Init(cfg *config.Config) *RabbitMQ {
	rabbit := &RabbitMQ{}

	conn, err := amqp.Dial(cfg.RabbitConnectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
	rabbit.conn = conn

	// Channel for consumer
	consumerChan, err := conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	rabbit.consumerChan = consumerChan

	// Channel for publisher
	publisherChan, err := conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	rabbit.publisherChan = publisherChan

	return rabbit
}

func (r *RabbitMQ) ConnectToTranscodingQueue(queueName string) (<-chan amqp.Delivery, error) {
	if _, err := r.consumerChan.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	msgs, err := r.consumerChan.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *RabbitMQ) ConnectToPublishQueue(queueName string) error {
	if _, err := r.publisherChan.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) Publish(queueName string, data []byte) error {
	return r.publisherChan.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
}

func (r *RabbitMQ) Close() {
	r.consumerChan.Close()
	r.publisherChan.Close()
	r.conn.Close()
}
