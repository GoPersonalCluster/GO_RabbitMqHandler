package publisher

import (
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type FilterPublisher struct {
	queueName string
}

func (sqn *FilterPublisher) SetChannel(channel *amqp.Channel) {
	// Implementação específica para configurar o canal do publisher
	// Exemplo: sqn.channel = channel
}

func (gp *FilterPublisher) Publish(message []byte, channel *amqp.Channel) error {
	// Implementação específica para publicar mensagem na fila

	err := channel.Publish(
		"FilterQueue",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})

	if err != nil {
		return err
	}

	return nil
}
