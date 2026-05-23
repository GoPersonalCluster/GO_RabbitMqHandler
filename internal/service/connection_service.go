package service

import (
	"go_rabbitmqhandler/internal/interfaces"
	"log"
	"slices"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	port       int
	host       string
	consumers  []consumers
	username   string
	password   string
	errors     []error
	channel    []*amqp.Channel
	publishers []publishers
	connection *amqp.Connection
}

// type queueConfig struct {
// 	name            string
// 	abstractFactory interfaces.AbstractFactoryHandler
// }

type Option func(*RabbitMQConfig)

func WithHost(host string) Option {
	return func(s *RabbitMQConfig) {
		s.host = host
	}
}
func WithPort(port int) Option {
	return func(s *RabbitMQConfig) {
		s.port = port
	}
}
func AddConsumer(queueName string, abstractFactory interfaces.AbstractFactoryHandler) Option {
	return func(s *RabbitMQConfig) {
		s.queueConfig = append(s.queueConfig,
			queueConfig{
				name:            queueName,
				abstractFactory: abstractFactory,
			},
		)
	}
}
func Username(username string) Option {
	return func(s *RabbitMQConfig) {
		s.username = username
	}
}
func Password(password string) Option {
	return func(s *RabbitMQConfig) {
		s.password = password
	}
}
func (c *RabbitMQConfig) Configure(opts ...Option) (*RabbitMQConfig, []error) {
	s := &RabbitMQConfig{
		host:        "localhost",                           // Default value
		port:        5672,                                  // Default value
		queueConfig: []queueConfig{{name: "defaultqueue"}}, // Default value
		username:    "admin",                               // Default value
		password:    "admin",                               // Default value],
	}

	for _, opt := range opts {
		opt(s)
	}
	return s, c.errors
}

func ConfigureConnection() Option {
	return func(rmc *RabbitMQConfig) {

		conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
		if err != nil {
			rmc.errors = append(rmc.errors, err)
			rmc.failOnError(err, "Erro ao conectar no RabbitMQ")
		}
		rmc.connection = conn

		defer conn.Close()
		// // 📡 Canal
		ch, err := conn.Channel()
		rmc.failOnError(err, "Erro ao abrir canal")
		defer ch.Close()
		rmc.channel = ch
	}
}
func (rmc *RabbitMQConfig) CloseConnection() {
	defer rmc.connection.Close()
}
func (rmc *RabbitMQConfig) CloseChannel(channelName string) {
	index := slices.IndexFunc(&rmc.consumers , func(s string) bool {
		return n > 15
	})

}

func (FoE *RabbitMQConfig) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		FoE.errors = append(FoE.errors, err)
	}
}
