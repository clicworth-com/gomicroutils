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

type SMSRequest struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Subject      string `json:"subject"`
	Message      string `json:"message"`
	MessageType  string `json:"message_type"`
	TemplateName string `json:"template_name"`
	AckRequired  bool   `json:"ack_required"`
	Priority     string `json:"priority"`
}

type SMSResponse struct {
	Status        bool   `json:"status"`
	Message       string `json:"message"`
	CorrelationId string `json:"correlation_id"`
}

var creatsmsq sync.Once

func (r *RabbitAMQPClient) SendSMS(msg []byte, cb func(SMSResponse)) (string, error) {

	creatsmsq.Do(func() {

		rqname := os.Getenv("SMS_QUEUE_NAME")
		if rqname == "" {
			rqname = "sms_queue"
		}

		q, err := r.Ch.QueueDeclare(
			"",    // name
			false, // durable
			true,  // delete when unused
			true,  // exclusive
			false, // noWait
			nil,   // arguments
		)
		if err != nil {
			log.Println("Unable to create mail queue ", err)
		}
		r.SMSReqQName = rqname
		r.SMSResQName = q.Name
	})

	if r.SMSReqQName == "" {
		return "", errors.New("unable to create mail queue")
	}

	msgs, err := r.Ch.Consume(
		r.SMSResQName, // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Println("Failed to register a consumer ", err)
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	corrId := randomString(32)

	err = r.Ch.PublishWithContext(ctx,
		"",            // exchange
		r.SMSReqQName, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       r.SMSResQName,
			Body:          msg,
		})
	if err != nil {
		log.Println("Failed to publish a message")
		return "", err
	}

	go func() {
		for d := range msgs {
			if corrId == d.CorrelationId {
				if cb != nil {
					var res SMSResponse
					if err := json.Unmarshal(d.Body, &res); err != nil {
						log.Println("unble to invoke sms response callback function")
					}
					cb(res)
				}
				break
			}
		}
	}()

	return corrId, nil

}
