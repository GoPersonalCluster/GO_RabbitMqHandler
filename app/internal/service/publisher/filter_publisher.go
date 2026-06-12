package publisher

import (
	amqp "github.com/streadway/amqp"
)

type FilterPublisher struct {
	queueName string
	channel   *amqp.Channel
}

func (sqn *FilterPublisher) SetChannel(channel *amqp.Channel) {
	// Implementação específica para configurar o canal do publisher
	sqn.channel = channel
}

func (gp *FilterPublisher) Publish(message []byte) error {
	// Implementação específica para publicar mensagem na fila

	err := gp.channel.Publish(
		"FilterQueue",
		"",
		false,
		false,
		GetAmqPublishingOptions(message))

	if err != nil {
		return err
	}

	return nil
}
