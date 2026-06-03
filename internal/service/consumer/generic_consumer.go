package consumer

import (
	"go_rabbitmqhandler/internal/interfaces"

	"github.com/streadway/amqp"
)

type GenericConsumer[T any] struct {
	name            string
	abstractFactory interfaces.AbstractFactoryHandler
	delivery        <-chan amqp.Delivery
}

func (Cc *GenericConsumer[T]) ConfigureConsumer(ch *amqp.Channel) error {
	q, err := ch.QueueDeclare(
		Cc.name, // nome
		true,    // durável
		false,   // auto-delete
		false,   // exclusiva
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		return err
	}
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
	if err != nil {
		return err
	}
	Cc.delivery = msgs
}

func (c *GenericConsumer[T]) Consume() {
	forever := make(chan bool)

	for d := range c.delivery {

		factory, err := c.abstractFactory.CreateStrategy(&d.Body)
		if err != nil {
			hm.failOnError(err, "Erro ao obter factory")
		}

		strategy, err := factory.CreateStrategy(&d.Body)
		if err != nil {
			hm.failOnError(err, "Erro ao criar estratégia")
		}
		response, err := strategy.Start()

		if publisher != nil {
			err := publisher.Publish(response)
			if err != nil {
				hm.failOnError(err, "Erro ao publicar mensagem")
			}
		}
		// ⚙️ Processamento da mensagem
		//		err := hm.processMessage( factory, d.Body )

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
