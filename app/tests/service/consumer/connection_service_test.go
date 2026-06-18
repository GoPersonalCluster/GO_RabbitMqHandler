package service_test

import (
	"go_rabbitmqhandler/internal/consumer"
	"go_rabbitmqhandler/internal/service"
	"testing"
)

func TestConnectionService(t *testing.T) {
	service := service.RabbitMQConfigComposite{}
	service.ConfigureConnection()

}

type TestConcreteFactory struct {
}

func (cS *TestConcreteFactory) CreateStrategy(event *consumer.IntegrationEvent) (
	consumer.StrategyHandler, error) {
	return TestConcreteStrategy{}, nil
}

type TestConcreteStrategy struct {
}

func (s *TestConcreteStrategy) Start() ([]byte, error) {
	print("Concrete strategy test run")
	return nil, nil
}
