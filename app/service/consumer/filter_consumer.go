package consumer

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/parser"
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/publisher"

	"github.com/streadway/amqp"
)

type FilterConsumer struct {
	config           ConsumerConfig
	delivery         <-chan amqp.Delivery
	genericPublisher publisher.PublisherInterface
	logPublisher     publisher.PublisherInterface
}

func (fC *FilterConsumer) SetConfiguration(config *ConsumerConfig) {
	println("set configuration")
	fC.config = *config
}
func (fC *FilterConsumer) configureConsumer(ch *amqp.Channel) error {
	println("iniciado configure consumer")
	q, err := ch.QueueDeclare(
		fC.config.QueueName,  // nome
		fC.config.Durable,    // durável
		fC.config.AutoDelete, // auto-delete
		fC.config.Exclusive,  // exclusiva
		fC.config.NoWait,     // no-wait
		fC.config.Args,       // args
	)
	println("declared queue ", fC.config.QueueName)
	if err != nil {
		return err
	}
	// 👂 Consumir mensagens
	msgs, err := ch.Consume(
		q.Name,
		fC.config.QueueName,  // nome
		fC.config.Durable,    // durável
		fC.config.AutoDelete, // auto-delete
		fC.config.Exclusive,  // exclusiva
		fC.config.NoWait,     // no-wait
		fC.config.Args,       // args
	)

	if err != nil {
		return err
	}
	fC.delivery = msgs
	fC.setGenericPublisher(ch)
	fC.setLogPublisher()

	return nil
}
func (fC *FilterConsumer) setLogPublisher() {
	fC.config.QueueName = "LogQueue"

}
func (fC *FilterConsumer) setGenericPublisher(ch *amqp.Channel) {
	publisher := publisher.GenericPublisher{}
	fC.genericPublisher = &publisher
}
func (fC *FilterConsumer) getStrategy(message IntegrationEvent) (StrategyHandler, error) {
	strategy, err := fC.config.AbstractFactory.CreateStrategy(&message)

	if err != nil {
		return nil, err
	}

	return strategy, nil
}

func (fC *FilterConsumer) Consume(ch *amqp.Channel) {
	println("consumer filter iniciado")
	fC.configureConsumer(ch)

	println("end consumer configuration")
	forever := make(chan bool)
	println("iniciado forever do consumer")
	for d := range fC.delivery {

		parser := parser.JsonParser[IntegrationEvent]{}
		i := parser.NewParser()
		message, err := i.Decode(d.Body)
		if err != nil {
			fC.publishErrorLog(err, ch, message)
			continue
		}
		strategy, err := fC.getStrategy(message)
		if err != nil {
			fC.publishErrorLog(err, ch, message)
			continue
		}

		response, err := strategy.Start()
		if err != nil {
			fC.publishErrorLog(err, ch, message)
			d.Ack(true)
			continue
		}
		nq, err := message.GetNextQueue()
		if nq == "" {
			continue
		}

		if fC.genericPublisher != nil {
			message.ExchangePayload(response)
			fC.genericPublisher.SetChannel(ch, nq)
			err := fC.genericPublisher.Publish(response)
			if err != nil {
				fC.publishErrorLog(err, ch, message)
				continue
			}
		}

		if err != nil {
			fC.publishErrorLog(err, ch, message)
			d.Ack(true)
			continue
		}

		d.Ack(true)

	}
	<-forever
}

func (fC *FilterConsumer) publishErrorLog(err error, ch *amqp.Channel, iE IntegrationEvent) {
	logPublisher := publisher.GenericPublisher{}
	logPublisher.SetChannel(ch, "LogQueue")
	iE.ExchangePayload([]byte(err.Error()))
	logPublisher.Publish([]byte(err.Error()))
}
