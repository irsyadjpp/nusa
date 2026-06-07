package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn *amqp.Connection
}

func NewPublisher(conn *amqp.Connection) *Publisher {
	return &Publisher{conn: conn}
}

func (p *Publisher) Publish(ctx context.Context, exchange, routingKey string, body []byte) error {
	channel, err := p.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}
	defer channel.Close()

	err = channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
