package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn          *amqp.Connection
	queue         string
	consumerTag   string
	retryInterval time.Duration
	maxRetries    int
}

type MessageHandler func(ctx context.Context, msg amqp.Delivery) error

func NewConsumer(conn *amqp.Connection, queue, consumerTag string) *Consumer {
	return &Consumer{
		conn:          conn,
		queue:         queue,
		consumerTag:   consumerTag,
		retryInterval: 5 * time.Second,
		maxRetries:    3,
	}
}

func (c *Consumer) Start(ctx context.Context, handler MessageHandler) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := c.consume(ctx, handler); err != nil {
				log.Printf("Consumer error: %v, retrying in %v...", err, c.retryInterval)
				time.Sleep(c.retryInterval)
				continue
			}
		}
	}
}

func (c *Consumer) consume(ctx context.Context, handler MessageHandler) error {
	channel, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}
	defer channel.Close()

	msgs, err := channel.Consume(
		c.queue,
		c.consumerTag,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg, ok := <-msgs:
			if !ok {
				return fmt.Errorf("consumer channel closed")
			}

			retryCount := 0
			for retryCount < c.maxRetries {
				if err := handler(ctx, msg); err != nil {
					retryCount++
					log.Printf("Handler error (attempt %d/%d): %v", retryCount, c.maxRetries, err)
					time.Sleep(c.retryInterval)
					continue
				}
				msg.Ack(false)
				break
			}

			if retryCount >= c.maxRetries {
				log.Printf("Max retries reached, sending to DLQ")
				msg.Nack(false, false)
			}
		}
	}
}
