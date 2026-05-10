package models

type RabbitMQHandlerConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Queue    string
}
