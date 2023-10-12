package amqp

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var createlogex sync.Once

func (a *RabbitAMQPClient) Write(p []byte) (int, error) {

	createlogex.Do(func() {
		lexName := os.Getenv("LOG_EXCHANGE_NAME")
		if lexName == "" {
			lexName = "logs_topic"
		}

		err := a.Ch.ExchangeDeclare(
			lexName, // name
			"topic", // type
			true,    // durable
			false,   // auto-deleted
			false,   // internal
			false,   // no-wait
			nil,     // arguments
		)
		if err != nil {
			log.Panic("unable to declare exchange for logs topic")
		}

		a.LogExName = lexName
	})

	if a.LogExName == "" {
		return 0, errors.New("unable to create logger exchange")
	}

	err := a.Ch.PublishWithContext(context.Background(),
		a.LogExName, // exchange
		"",          // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        p,
		})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
