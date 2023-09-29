package amqp

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitAMQPClient struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	log_ex   string // exchange name for logs
	trac_ex  string // exchange name for tracker
	mail_key string // routing key for sending mail
}

var once sync.Once

var amqpClient *RabbitAMQPClient

func Get() *RabbitAMQPClient {
	once.Do(func() {

		conn, err := connectRabbit()
		if err != nil {
			log.Panic("unable to connect to rabbitmq")
		}

		ch, err := conn.Channel()
		if err != nil {
			log.Panic("unable to create channel on rabbitmq")
		}

		lexName := os.Getenv("LOG_EXCHANGE_NAME")
		if lexName == "" {
			lexName = "logs_topic"
		}

		err = ch.ExchangeDeclare(
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

		texName := os.Getenv("TRACKER_EXCHANGE_NAME")
		if texName == "" {
			texName = "tracker_topic"
		}

		err = ch.ExchangeDeclare(
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

		mailKey := os.Getenv("MAIL_ROUTING_KEY")
		if mailKey == "" {
			mailKey = "mail_queue"
		}

		amqpClient = &RabbitAMQPClient{
			conn:     conn,
			ch:       ch,
			log_ex:   lexName,
			trac_ex:  texName,
			mail_key: mailKey,
		}
	})
	return amqpClient
}

func connectRabbit() (*amqp.Connection, error) {
	var counts int64
	var backOff = 5 * time.Second
	var connection *amqp.Connection

	//don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not ready yet...")
			counts++
		} else {
			fmt.Println("Connected to rabbitMQ.")
			connection = c
			break
		}

		if counts > 10 {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println("backing off ....")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}

func (a *RabbitAMQPClient) Stop() {
	if a.ch != nil {
		a.ch.Close()
	}
	if a.conn != nil {
		a.conn.Close()
	}
}
