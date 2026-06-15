package consumer

import (
	"go_rabbitmqhandler/internal/service/publisher"

	"github.com/streadway/amqp"
)

type GenericConsumer struct {
	config          ConsumerConfig
	delivery        <-chan amqp.Delivery
	filterPublisher publisher.PublisherInterface
	logPublisher    publisher.PublisherInterface
}

func (gC *GenericConsumer) ConfigureConsumer(ch *amqp.Channel) error {
	q, err := ch.QueueDeclare(
		gC.config.QueueName,  // nome
		gC.config.Durable,    // durável
		gC.config.AutoDelete, // auto-delete
		gC.config.Exclusive,  // exclusiva
		gC.config.NoWait,     // no-wait
		gC.config.Args,       // args
	)
	if err != nil {
		return err
	}
	// 👂 Consumir mensagens
	msgs, err := ch.Consume(
		q.Name,
		gC.config.QueueName,  // nome
		gC.config.Durable,    // durável
		gC.config.AutoDelete, // auto-delete
		gC.config.Exclusive,  // exclusiva
		gC.config.NoWait,     // no-wait
		gC.config.Args,       // args
	)

	if err != nil {
		return err
	}
	gC.delivery = msgs
	gC.setFilterPublisher(ch)
	return nil
}
func (cP *GenericConsumer) setLogPublisher(queueName string) {
	cP.config.QueueName = queueName

}

func (cP *GenericConsumer) setFilterPublisher(ch *amqp.Channel) {
	publisher := publisher.GenericPublisher{}
	publisher.SetChannel(ch, "FilterQueue")

	cP.filterPublisher = &publisher
}
func (cP *GenericConsumer) getStrategy(body []byte) (StrategyHandler, error) {
	strategy, err := cP.config.AbstractFactory.CreateStrategy(&body)

	if err != nil {
		//cP.failOnError(err, "Erro ao obter factory")
	}

	return strategy, nil
}


func (c *GenericConsumer) Consume(ch *amqp.Channel) {

	forever := make(chan bool)

	for d := range c.delivery {

		strategy, err := c.getStrategy(d.Body)
		if err != nil {
			//hm.failOnError(err, "Erro ao obter estratégia")
		}

		response, err := strategy.Start()

		if c.filterPublisher != nil {
			err := c.filterPublisher.Publish(response)
			if err != nil {
				//hm.failOnError(err, "Erro ao publicar mensagem")

			}
		}

		if err != nil {
			//log.Printf("❌ Erro ao processar: %s", err)
			//d.Nack(false, true) // requeue
			d.Ack(false)
			continue
		}

		// ✅ Confirma processamento
		d.Ack(false)

	}
	<-forever

}
