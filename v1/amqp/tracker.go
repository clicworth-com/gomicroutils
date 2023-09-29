package amqp

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type TrackEvent struct {
	Sender    string    `json:"sender"`
	EventName string    `json:"event_name"`
	EventData string    `json:"event_data"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *RabbitAMQPClient) TrackEvent(e []byte, sender, name string) error {
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
	return a.track(evtB, a.trac_ex)
}

func (a *RabbitAMQPClient) track(p []byte, ex string) error {
	err := a.ch.PublishWithContext(context.Background(),
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
