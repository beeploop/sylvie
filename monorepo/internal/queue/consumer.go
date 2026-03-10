package queue

type Consumer interface {
	Consume(func(Job) error) error
}
