package rabbitMq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	port      int
	host      string
	queueName []string
	username  string
	password  string
	messages  <-chan amqp.Delivery
	errors    []error
	channel   *amqp.Channel
}
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
func WithQueues(queueName []string) Option {
	return func(s *RabbitMQConfig) {
		s.queueName = queueName
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
func NewConnection(opts ...Option) *RabbitMQConfig {
	s := &RabbitMQConfig{
		host:      "localhost",              // Default value
		port:      5672,                     // Default value
		queueName: []string{"defaultqueue"}, // Default value
		username:  "admin",                  // Default value
		password:  "admin",                  // Default value],
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func ConfigureConnection() Option {
	return func(rmc *RabbitMQConfig) {
		conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
		rmc.failOnError(err, "Erro ao conectar no RabbitMQ")
		defer conn.Close()

		// // 📡 Canal
		ch, err := conn.Channel()
		rmc.failOnError(err, "Erro ao abrir canal")
		defer ch.Close()
		rmc.channel = ch
	}
}

func ConfigureQueue() Option {
	f := func(rmc *RabbitMQConfig) {
		for _, qn := range rmc.queueName {
			q, err := rmc.channel.QueueDeclare(
				qn,    // nome
				true,  // durável
				false, // auto-delete
				false, // exclusiva
				false, // no-wait
				nil,   // args
			)
			rmc.failOnError(err, "Erro ao declarar fila")

			// 👂 Consumir mensagens
			msgs, err := rmc.channel.Consume(
				q.Name,
				"",    // consumer
				false, // auto-ack (false = manual)
				false, // exclusive
				false, // no-local
				false, // no-wait
				nil,   // args
			)
			rmc.failOnError(err, "Erro ao registrar consumer")
			rmc.channel = msgs
		}
		// // 📬 Declarar fila

	}
	return f
}

func (cho *RabbitMQConfig) configureHost() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	cho.failOnError(err, "Erro ao conectar no RabbitMQ")
	defer conn.Close()

	// // 📡 Canal
	ch, err := conn.Channel()
	cho.failOnError(err, "Erro ao abrir canal")
	defer ch.Close()

	// // 📬 Declarar fila
	q, err := ch.QueueDeclare(
		"LLM_QUEUE", // nome
		true,        // durável
		false,       // auto-delete
		false,       // exclusiva
		false,       // no-wait
		nil,         // args
	)
	cho.failOnError(err, "Erro ao declarar fila")

	// 👂 Consumir mensagens
	msgs, err := ch.Consume(
		q.Name,
		"",    // consumer
		false, // auto-ack (false = manual)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	cho.failOnError(err, "Erro ao registrar consumer")
	return msgs, nil
}
func (FoE *RabbitMQConfig) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
