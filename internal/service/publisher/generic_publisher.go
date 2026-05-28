package publisher

import (
	"errors"
	"go_rabbitmqhandler/internal/model"

	amqp "github.com/rabbitmq/amqp091-go"
)

type GenericPublisher struct {
	QueueName string
	channel   *amqp.Channel
}

func (sqn *GenericPublisher) GetIdentity() model.PublisherIdentity {
	return model.PublisherIdentity{}
}
func (sqn *GenericPublisher) SetChannel(channel *amqp.Channel) {
	// Implementação específica para configurar o canal do publisher
	// Exemplo: sqn.channel = channel
}

func (gp *GenericPublisher) Publish(message []byte) error {
	// Implementação específica para publicar mensagem na fila
	if gp.QueueName == "" {
		return errors.New("QueueName não pode ser vazio")
	}

	err := gp.channel.Publish(gp.QueueName, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        message,
	})

	if err != nil {
		return err
	}

	return nil
}
