package rabbit

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/nyatify/nyatify/pkg/model"
	"github.com/streadway/amqp"
)

// Client is a rabbitmq instance entry point
type Client struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func New(amqpURI string) (*Client, error) {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	return &Client{
		Conn:    conn,
		Channel: channel,
	}, nil
}

func NewWithQueue(amqpURI string, queueName string) (*Client, error) {
	client, err := New(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a client: %v", err)
	}

	queue, err := client.Channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %v", err)
	}

	return &Client{
		Conn:    client.Conn,
		Channel: client.Channel,
		Queue:   queue,
	}, nil
}

func (s *Client) Close() error {
	if err := s.Channel.Close(); err != nil {
		return fmt.Errorf("failed to close channel: %v", err)
	}

	if err := s.Conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %v", err)
	}

	return nil
}

func (s *Client) Schedule(n model.Notification) error {
	body, err := n.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %v", err)
	}

	err = s.Channel.Publish(
		"",       // exchange
		n.Client, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish notification: %v", err)
	}

	log.Debug().Msg("Scheduled a notification:" + n.Client)

	return nil
}
