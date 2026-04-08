package queue

type Publisher interface {
	Publish(Job) error
}
