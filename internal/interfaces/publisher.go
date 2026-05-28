package interfaces

import (
	"go_rabbitmqhandler/internal/model"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	Publish(message []byte) error
	GetIdentity() model.PublisherIdentity
	SetChannel(channel *amqp.Channel)
}
