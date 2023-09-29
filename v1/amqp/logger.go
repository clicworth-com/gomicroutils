package amqp

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (a *RabbitAMQPClient) Write(p []byte) (int, error) {
	err := a.ch.PublishWithContext(context.Background(),
		a.log_ex, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        p,
		})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
