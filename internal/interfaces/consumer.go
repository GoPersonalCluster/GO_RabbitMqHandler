package interfaces

import (
	"github.com/streadway/amqp"
)

type Consumer[T any] interface {
	Consume() (T, error)
	SetQueue(queueName string)
}
