package publisher

import amqp "github.com/streadway/amqp"

type PublisherInterface interface {
	Publish(message []byte) error
	SetChannel(channel *amqp.Channel, queueName string)
}

func GetAmqPublishingOptions(body []byte) amqp.Publishing {
	return amqp.Publishing{
		ContentType: "application/octet-stream",
		Body:        body,
	}
}
