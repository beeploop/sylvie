package queue

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type consumerImpl struct {
	Channel   *amqp.Channel
	QueueName string
}

func NewConsumer(ch *amqp.Channel, queueName string) *consumerImpl {
	return &consumerImpl{
		Channel:   ch,
		QueueName: queueName,
	}
}

func (c *consumerImpl) Consume(handler func(Job) error) error {
	msgs, err := c.Channel.Consume(
		c.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {

		var job Job
		if err := json.Unmarshal(msg.Body, &job); err != nil {
			msg.Nack(false, false)
			continue
		}

		if err := handler(job); err != nil {
			msg.Nack(false, true)
			continue
		}

		msg.Ack(false)
	}

	return nil
}
