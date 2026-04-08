package queue

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type publisherImpl struct {
	Channel   *amqp.Channel
	QueueName string
}

func NewPublisher(ch *amqp.Channel, queueName string) *publisherImpl {
	return &publisherImpl{
		Channel:   ch,
		QueueName: queueName,
	}
}

func (p *publisherImpl) Publish(job Job) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	p.Channel.Publish(
		"",
		p.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)

	return nil
}
