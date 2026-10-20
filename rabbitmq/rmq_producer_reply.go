package main

// this script uses request-reply pattern to send a request to a service and wait for the response
// it uses a temporary exclusive queue for the reply and a correlation ID to match the response with the request
// it also handles gzip compressed response and uncompresses it before parsing the JSON
// it also handles timeout for the response and prints an error message if the response is not received within the timeout period
// it also handles JSON parsing errors and prints an error message if the response is not valid JSON

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/url"
	"bytes"
	"compress/gzip"
	"io"

	"github.com/streadway/amqp"
)

func main() {
	// init RabbitMQ settings
	var rmq_host string = "localhost"
	var rmq_port string = "5672"
	var rmq_user string = "appuser"
	var rmq_pwd string  = "xxxxxxxxxxxxxx"

	// Construct the RabbitMQ URL using the net/url struct
	rabbitmq_url := &url.URL{
		Scheme: "amqp",
		User:   url.UserPassword(rmq_user, rmq_pwd),
		Host:   rmq_host + ":" + rmq_port,
		Path:   "/", 
	}

	var rmq_routing_key string   = "demo.route"
	var rmq_exchange string      = "demo.exchange"
	var rmq_exchange_type string = "topic"

	// service response timeout
	var service_resp_timeout int = 5

	// make a connection
	conn, err := amqp.Dial(rabbitmq_url.String())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	// Declare the exchange (must match service config)
	err = ch.ExchangeDeclare(
		rmq_exchange,      // name
		rmq_exchange_type, // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	// Create exclusive reply queue
	replyQueue, err := ch.QueueDeclare(
		"amandavaRQ", // name
		false,        // durable
		true,         // delete when unused
		true,         // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare reply queue: %v", err)
	}

	// Bind reply queue to exchange so we can receive responses
	err = ch.QueueBind(
		replyQueue.Name, // queue name
		replyQueue.Name, // routing key (same as queue name for direct reply)
		rmq_exchange,    // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind reply queue: %v", err)
	}

	// Start consuming from reply queue
	msgs, err := ch.Consume(
		replyQueue.Name, // queue
		"amandavaC",     // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Fatalf("Failed to consume: %v", err)
	}

	// Build request payload
	// Options: {"ip": "x.x.x.x"}, {"mac": "aabbccddeeff"}, {"modemMac": "aabbccddeeff"}
	request := map[string]interface{}{
		// "ip": "10.213.131.18", // single ip lookup
		// "mac": "287AEE0982BF", // single mac lookup
		// "modemMac": "D46C6D39EF11", // single modemMac lookup 
		"mac": []string{"D46C6D39EF11", "3C37863AE288", "287AEE0982BF"}, // bulk mac lookup
	}
	body, _ := json.Marshal(request)

	// Generate correlation ID
	corrID := fmt.Sprintf("test-%d", time.Now().UnixNano())

	// Publish message
	err = ch.Publish(
		rmq_exchange,    // exchange
		rmq_routing_key, // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrID,
			ReplyTo:       replyQueue.Name,
			Timestamp:     time.Now(),
			Body:          body,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish: %v", err)
	}

	fmt.Printf("Sent request:%s\n  Routing Key: %s\n  Correlation ID: %s\n  Body: %s\n\n", request, rmq_routing_key, corrID, body)
	fmt.Println("Waiting for response...")

	// Wait for response with timeout
	select {
	case msg := <-msgs:
		fmt.Println("CorrId: ", msg.CorrelationId)
		if msg.CorrelationId == corrID {
			// get the response message
			bodyBytes := msg.Body

			// Check if the sender compressed the message
			if msg.ContentEncoding == "gzip" {
				fmt.Println("----------- Message is gzip compressed --------------")
				reader, err := gzip.NewReader(bytes.NewReader(msg.Body))
				if err == nil {
					defer reader.Close()
					if uncompressed, err := io.ReadAll(reader); err == nil {
						bodyBytes = uncompressed
					}
				}
			}
			// Now use bodyBytes instead of msg.Body
			fmt.Printf("Response received:\n  Correlation ID: %s\n  Body: %s\n", msg.CorrelationId, string(bodyBytes))

			var response map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &response); err != nil {
				// JSON parsing failed
				fmt.Printf("JSON Unmarshal Error: %v\n", err)
			} else {
				pretty, _ := json.MarshalIndent(response, "", "  ")
				fmt.Printf("\nFormatted:\n%s\n", pretty)
			}
		}
	case <-time.After(time.Duration(service_resp_timeout) * time.Second):
		fmt.Println("Timeout: No response received within ", service_resp_timeout, " seconds")
	}
}
