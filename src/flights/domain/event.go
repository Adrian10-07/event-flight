package domain

type EventPublisher interface {
	Publish(eventName string, data interface{}) error
}
