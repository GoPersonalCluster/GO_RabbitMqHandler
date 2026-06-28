package consumer

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/parser"
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/publisher"

	"github.com/streadway/amqp"
)

type FilterConsumer struct {
	config           FilterConfig
	delivery         <-chan amqp.Delivery
	genericPublisher publisher.PublisherInterface
	logPublisher     publisher.PublisherInterface
}

func (sC *FilterConsumer) SetConfiguration(config *FilterConfig) {
	sC.config = *config
}
func (cC *FilterConsumer) configureConsumer(ch *amqp.Channel) error {
	q, err := ch.QueueDeclare(
		cC.config.QueueName,  // nome
		cC.config.Durable,    // durável
		cC.config.AutoDelete, // auto-delete
		cC.config.Exclusive,  // exclusiva
		cC.config.NoWait,     // no-wait
		cC.config.Args,       // args
	)
	println("declared queue ", cC.config.QueueName)
	if err != nil {
		return err
	}
	// 👂 Consumir mensagens
	msgs, err := ch.Consume(
		q.Name,
		cC.config.QueueName,  // nome
		cC.config.Durable,    // durável
		cC.config.AutoDelete, // auto-delete
		cC.config.Exclusive,  // exclusiva
		cC.config.NoWait,     // no-wait
		cC.config.Args,       // args
	)

	if err != nil {
		return err
	}
	cC.delivery = msgs
	cC.setGenericPublisher(ch)
	cC.setLogPublisher()

	return nil
}
func (cP *FilterConsumer) setLogPublisher() {
	cP.config.QueueName = "LogQueue"

}
func (cP *FilterConsumer) setGenericPublisher(ch *amqp.Channel) {
	publisher := publisher.GenericPublisher{}
	cP.genericPublisher = &publisher
}

func (c *FilterConsumer) Consume(ch *amqp.Channel) {
	c.configureConsumer(ch)

	println("end consumer configuration")
	forever := make(chan bool)

	for d := range c.delivery {

		parser := parser.JsonParser[IntegrationEvent]{}
		i := parser.NewParser()
		model, err := i.Decode(d.Body)
		if err != nil {
			c.publishErrorLog(err, ch, model)
			continue
		}
		iE, err := c.config.AbstractFactory.GetQueue(&model)
		if err != nil {
			c.publishErrorLog(err, ch, iE)
			continue
		}

		qn, err := iE.GetNextQueue()

		if err != nil {
			c.publishErrorLog(err, ch, iE)
			continue
		}
		c.genericPublisher.SetChannel(ch, qn)
		err = c.genericPublisher.Publish(iE.Payload)

		if err != nil {
			c.publishErrorLog(err, ch, model)
			d.Ack(true)
			continue
		}

		d.Ack(true)

	}
	<-forever

}

func (gC *FilterConsumer) publishErrorLog(err error, ch *amqp.Channel, iE IntegrationEvent) {
	logPublisher := publisher.GenericPublisher{}
	logPublisher.SetChannel(ch, "LogQueue")
	iE.ExchangePayload([]byte(err.Error()))
	logPublisher.Publish([]byte(err.Error()))
}
