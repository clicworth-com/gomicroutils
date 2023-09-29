package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type TrackEvent struct {
	Sender    string    `json:"sender"`
	EventName string    `json:"event_name"`
	EventData string    `json:"event_data"`
	CreatedAt time.Time `json:"created_at"`
}

var createtracex sync.Once

func (a *RabbitAMQPClient) TrackEvent(e []byte, sender, name string) error {

	createtracex.Do(func() {
		texName := os.Getenv("TRACKER_EXCHANGE_NAME")
		if texName == "" {
			texName = "tracker_topic"
		}

		err := a.Ch.ExchangeDeclare(
			texName, // name
			"topic", // type
			true,    // durable
			false,   // auto-deleted
			false,   // internal
			false,   // no-wait
			nil,     // arguments
		)
		if err != nil {
			log.Panic("unable to declare exchange for tracker topic")
		}
		a.TracExName = texName
	})

	if a.TracExName == "" {
		return errors.New("unable to create tracker exchange")
	}

	evt := TrackEvent{
		Sender:    sender,
		EventName: name,
		EventData: string(e),
		CreatedAt: time.Now(),
	}

	evtB, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	return a.track(evtB, a.TracExName)
}

func (a *RabbitAMQPClient) track(p []byte, ex string) error {
	err := a.Ch.PublishWithContext(context.Background(),
		ex,    // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        p,
		})
	if err != nil {
		return err
	}
	return nil
}
