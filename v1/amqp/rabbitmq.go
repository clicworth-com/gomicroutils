package amqp

import (
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitAMQPClient struct {
	Conn           *amqp.Connection
	Ch             *amqp.Channel
	LogExName      string // exchange name for logs
	TracExName     string // exchange name for tracker
	MailRoutingKey string // routing key for sending mail
	MailQName      string
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

		amqpClient = &RabbitAMQPClient{
			Conn: conn,
			Ch:   ch,
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
	if a.Ch != nil {
		a.Ch.Close()
	}
	if a.Conn != nil {
		a.Conn.Close()
	}
}
