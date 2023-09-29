package amqp

import (
	"context"
	"log"
	"math/rand"
	"os"
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

func (r *RabbitAMQPClient) SendMail(msg []byte) error {

	qname := os.Getenv("MAIL_QUEUE_NAME")
	if qname == "" {
		qname = "mail_queue"
	}

	q, err := r.ch.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		log.Println("Unable to create mail queue ", err)
		return err
	}

	msgs, err := r.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println("Failed to register a consumer ", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	corrId := randomString(32)

	log.Println("sending mail request with corrrId: " + corrId)

	err = r.ch.PublishWithContext(ctx,
		"",    // exchange
		qname, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
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
