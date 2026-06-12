package publisher

import (
	amqp "github.com/streadway/amqp"
)

type LogPublisher struct {
	queueName string
	channel   *amqp.Channel
}

func (sqn *LogPublisher) SetChannel(channel *amqp.Channel) {
	// Implementação específica para configurar o canal do publisher
	sqn.channel = channel
}

func (gp *LogPublisher) Publish(message []byte) error {
	// Implementação específica para publicar mensagem na fila

	err := gp.channel.Publish(
		"LogQueue",
		"",
		false,
		false,
		GetAmqPublishingOptions(message))

	if err != nil {
		return err
	}

	return nil
}
