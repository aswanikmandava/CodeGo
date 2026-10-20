package main

import (
	"CodeGo/config"
	"CodeGo/logger"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"time"
)

func main() {
	logger := logger.GetLogger()
	logger.Debug("Starting RMQ producer...")
	if wd, err := os.Getwd(); err == nil {
		logger.Debug("Current working directory: ", wd)
	}
	// load application configuration from app.yml
	c, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("Failed to load config: %v", err)
	} else {
		logger.Debugf("Loaded config: %+v", c)
	}

	// make a connection to RabbitMQ server
	rmq_url := fmt.Sprintf("amqp://%s:%s@%s:%d", c.RabbitMQ.User, c.RabbitMQ.Pass, c.RabbitMQ.Host, c.RabbitMQ.Port)
	conn, err := amqp.Dial(rmq_url)
	if err != nil {
		logger.Fatalf("Failed to establish connection to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// get a channel using the connection
	chn, err := conn.Channel()
	if err != nil {
		logger.Fatalf("Unable to get a channel: %v", err)
	}
	defer chn.Close()

	// declare an exchange on the server
	err = chn.ExchangeDeclare(
		c.RabbitMQ.Exchange,
		c.RabbitMQ.ExchangeType,
		true,  // durable    - survives broker restarts
		false, // autodelete - if true, exchanges will be deleted when there are no remaining bindings and not restored on server restart
		false, // internal   - if true, it doesn't allow publishing to this exchange as it is meant for internal usage
		false, // noWait     - if true, declare without waiting for a confirmation from the server
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to declare the exchange: %v", err)
	}

	// declare a durable queue
	_, err = chn.QueueDeclare(
		c.RabbitMQ.Queue,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // noWait
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to declare the queue: %v", err)
	}

	// bind the queue to the exchange with a routing key
	err = chn.QueueBind(
		c.RabbitMQ.Queue,
		c.RabbitMQ.RoutingKey,
		c.RabbitMQ.Exchange,
		false, // noWait
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to bind the queue to the exchange: %v", err)
	}

	// Register listeners before publishing.
	returns := chn.NotifyReturn(make(chan amqp.Return, 1))
	confirms := chn.NotifyPublish(make(chan amqp.Confirmation, 1))

	// Enable publisher-confirm mode.
	if err := chn.Confirm(false); err != nil {
		logger.Errorf("failed to enable publisher confirms: %v", err)
	}

	// Publish the message. mandatory=true makes RabbitMQ
	// return it if no queue matches the routing key.
	err = chn.Publish(
		c.RabbitMQ.Exchange,
		c.RabbitMQ.RoutingKey,
		true,  // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp.Persistent,
			Body:         []byte("My first message to RabbitMQ"),
		},
	)
	if err != nil {
		logger.Errorf("failed to publish message: %v", err)
	}

	// Wait for RabbitMQ's publisher confirmation.
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	select {
	case confirmation, ok := <-confirms:
		if !ok {
			logger.Errorf("publisher confirmation channel closed")
		}

		if !confirmation.Ack {
			logger.Errorf("RabbitMQ negatively acknowledged the message")
		}

		// RabbitMQ sends basic.return before basic.ack for
		// an unroutable mandatory message. The buffered return
		// channel lets us inspect that return after the ACK.
		select {
		case r, ok := <-returns:
			if !ok {
				logger.Errorf("return notification channel closed")
			}

			logger.Errorf(
				"UNROUTABLE MESSAGE: exchange=%s routingKey=%s replyCode=%d replyText=%s body=%s",
				r.Exchange,
				r.RoutingKey,
				r.ReplyCode,
				r.ReplyText,
				string(r.Body),
			)
			logger.Error("message was returned as unroutable")

		default:
			logger.Info("RabbitMQ acknowledged the message; it was routed to at least one queue")
		}

	case <-timer.C:
		logger.Error("timed out waiting for RabbitMQ publisher confirmation")
	}

}
