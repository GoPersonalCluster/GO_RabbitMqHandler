package interfaces

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	Publish(message []byte) error
	SetChannel(channel *amqp.Channel)
}
