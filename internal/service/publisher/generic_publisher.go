package publisher

import ("errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

type GenericPublisher struct {
	QueueName string
}
func (sqn *GenericPublisher) SetQueueName(queueName string)  {
	sqn.QueueName = queueName	
}

func (gp *GenericPublisher) Publish(message []byte) error {
	// Implementação específica para publicar mensagem na fila
	if gp.QueueName == "" {
		return errors.New("QueueName não pode ser vazio")
	}


	return nil
}
