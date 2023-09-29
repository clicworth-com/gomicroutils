package amqp

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MailRequest struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Subject      string `json:"subject"`
	Message      string `json:"message"`
	MessageType  string `json:"message_type"`
	TemplateName string `json:"template_name"`
	AckRequired  bool   `json:"ack_required"`
	Priority     string `json:"priority"`
}

type MailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var createmailq sync.Once

func (r *RabbitAMQPClient) SendMail(msg []byte) error {

	createmailq.Do(func() {

		rkname := os.Getenv("MAIL_ROUTING_KEY_NAME")
		if rkname == "" {
			rkname = "mail_queue"
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
		r.MailRoutingKey = rkname
		r.MailQName = q.Name
	})

	if r.MailQName == "" {
		return errors.New("unable to create mail queue")
	}

	msgs, err := r.Ch.Consume(
		r.MailQName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		log.Println("Failed to register a consumer ", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	corrId := randomString(32)

	log.Println("sending mail request with corrrId: " + corrId)

	err = r.Ch.PublishWithContext(ctx,
		"",               // exchange
		r.MailRoutingKey, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       r.MailQName,
			Body:          msg,
		})
	if err != nil {
		log.Println("Failed to publish a message")
		return err
	}

	for d := range msgs {
		if corrId == d.CorrelationId {
			log.Println("Mail response for corrrId: " + corrId + " " + string(d.Body))
			break
		}
	}

	return nil

}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
