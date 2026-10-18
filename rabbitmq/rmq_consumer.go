package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"CodeGo/logger"
	"CodeGo/config"
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"context"
)

func main() {
	logger := logger.GetLogger()
	// listen for interrupt signal(Ctrl+C) and terminate signal(SIGTERM) to gracefully shutdown the consumer
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	logger.Debug("Starting RMQ consumer...")
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

	// declare the same que as the producer
	_, err = chn.QueueDeclare(
		c.RabbitMQ.Queue,
		true,	// durable
		false,	// auto-delete
		false,	// exclusive
		false,	// noWait
		nil,
	)
	if err != nil {
		logger.Fatalf("Failed to declare queue: %v", err)
	}

	// limit the number of unacknowledged messages to 1
	// this is important to avoid overwhelming the consumer with too many messages at once
	// it ensures that the consumer processes one message at a time and acknowledges it before receiving the next one
	err = chn.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		logger.Fatalf("Failed to set QoS: %v", err)
	}

	// Use a named consumer so we can cancel it explicitly.
	consumerTag := "my-consumer"

	// start consuming messages from the queue
	msgs, err := chn.Consume(
		c.RabbitMQ.Queue,
		consumerTag,    // consumer tag
		false, // auto-ack  - set to false to manually acknowledge messages after processing
		false, // exclusive - set to false to allow multiple consumers on the same queue
		false, // no-local  - set to false to receive messages published by the same connection
		false, // no-wait	- set to false to wait for the server's response and ensure the consumer is registered before proceeding
		nil,   // args
	)
	if err != nil {
		logger.Fatalf("Failed to register a consumer: %v", err)
	}
	
	// create a channel to signal when the program should exit
	// for example, by using a signal handler to catch SIGINT or SIGTERM and then closing the channel and connection before exiting
	done := make(chan struct{})

	// start a goroutine to process incoming messages
	go func() {
		defer close(done)
		for d := range msgs {
			logger.Infof("Received a message: %s", d.Body)
			// simulate processing time
			time.Sleep(1 * time.Second)
			// acknowledge the message after processing
			if err := d.Ack(false); err != nil {
				logger.Errorf("Failed to acknowledge message: %v", err)
			}
		}
	}()

	logger.Debug("Consumer is running. Press Ctrl+C to Stop ...")
	// Wait until the OS signals a shutdown request
	<- ctx.Done()
	logger.Info("Shutdown received. Stopping consumer ...")

	// stop receiving new deliveries
	err = chn.Cancel(consumerTag, false)
	if err != nil {
		logger.Errorf("Failed to cancel the consumer: %v", err)
	}

	// wait for the goroutine to finish processing any remaining messages
	<- done
	logger.Info("Consumer stopped. Exiting ...")
}