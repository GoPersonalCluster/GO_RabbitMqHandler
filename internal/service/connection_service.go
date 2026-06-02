package service

import (
	"fmt"
	"go_rabbitmqhandler/internal/interfaces"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConfigComposite[T any] struct {
	channel    ChannelConfig[T]
	connection *amqp.Connection
}
type ChannelConfig[T any] struct {
	publishers []interfaces.Publisher
	consumers  []interfaces.Consumer[T]
	channel    *amqp.Channel
	name       string
}

func FindOrElse[T any](
	items []T,
	predicate func(T) bool,
	orElse func() T,
) T {
	for _, item := range items {
		if predicate(item) {
			return item
		}
	}

	return orElse()
}

func (rmc *RabbitMQConfigComposite[T]) AddConsumer(channelName string,
	queueName string,
	abstractFactory interfaces.AbstractFactoryHandler,
	consumer interfaces.Consumer[T]) {

	rmc.channel.consumers = append(rmc.channel.consumers, consumer)
}

func (rmc *RabbitMQConfigComposite[T]) AddPublisher(publisher interfaces.Publisher) {
	rmc.channel.publishers = append(rmc.channel.publishers, publisher)
}

func (rmc *RabbitMQConfigComposite[T]) ConfigureConnection(host string, port int, un string, pwd string) {
	conn, err := amqp.Dial(fmt.Sprintf(`amqp://%s:%s@%s:%d/`,
		un,
		pwd,
		host,
		port))
	rmc.failOnError(err, "Erro ao conectar no RabbitMQ")
	defer conn.Close()

	// // 📡 Canal
	ch, err := conn.Channel()
	rmc.channel.channel = ch

	rmc.failOnError(err, "Erro ao abrir canal")
	defer rmc.CloseConnection()
}

func (rmc *RabbitMQConfigComposite[T]) CloseConnection() {
	rmc.connection.Close()
}

func (FoE *RabbitMQConfigComposite[T]) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		//FoE.errors = append(FoE.errors, err)
	}
}
