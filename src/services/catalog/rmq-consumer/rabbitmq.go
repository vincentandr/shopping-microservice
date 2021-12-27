package rmqConsumer

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/streadway/amqp"
	rbmq "github.com/vincentandr/shopping-microservice/src/internal/rabbitmq"
	"github.com/vincentandr/shopping-microservice/src/model"
	db "github.com/vincentandr/shopping-microservice/src/services/catalog/catalogdb"
)

// RabbitMQ ...
type RbmqListener struct {
	Msgs <-chan amqp.Delivery
}

// NewRabbitMQ instantiates consumer instance
func NewConsumer(r *rbmq.Rabbitmq) (*RbmqListener, error) {
	q, err := r.Channel.QueueDeclare(
                "catalogQueue",    // name
                false, // durable
                false, // delete when unused
                true,  // exclusive
                false, // no-wait
                nil,   // arguments
    )
	if err != nil {
		return nil, fmt.Errorf("amqChannel.QueueDeclare %w", err)
	}

	err = r.Channel.QueueBind(
                        q.Name,
                        "event.payment.success",
                        "tasks",
                        false,
                        nil)
	if err != nil {
		return nil, fmt.Errorf("amqChannel.ExchangeDeclare %w", err)
	}

	msgs, err := r.Channel.Consume(
                q.Name, // queue
                "",     // consumer
                false,   // auto ack
                false,  // exclusive
                false,  // no local
                false,  // no wait
                nil,    // args
    )
	if err != nil {
		return nil, fmt.Errorf("ch.Consume %w", err)
	}

	return &RbmqListener{Msgs: msgs}, nil
}

func (l *RbmqListener) EventHandler(a *db.Action) {
	go func(){
		for msg := range l.Msgs {
			var order model.UserOrder

			switch msg.RoutingKey {
			case "event.payment.success":
				gob.NewDecoder(bytes.NewReader(msg.Body)).Decode(&order)
				err := EventPaymentSuccessful(a, order)
				if err != nil{
					fmt.Println(err)
				}
			}
		}
	}()
}

func EventPaymentSuccessful(a *db.Action, order model.UserOrder) error {
	err := a.UpdateProducts(context.Background(), order.Items)
	if err != nil {
		return fmt.Errorf("failed to execute remove cart items event payment: %v", err)
	}

	return nil
}