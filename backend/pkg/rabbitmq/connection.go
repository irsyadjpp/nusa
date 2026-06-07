package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConnectionManager struct {
	conn *amqp.Connection
	url  string
}

func NewConnectionManager(url string) *ConnectionManager {
	return &ConnectionManager{
		url: url,
	}
}

func (cm *ConnectionManager) Connect(ctx context.Context) error {
	var err error
	cm.conn, err = amqp.Dial(cm.url)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return nil
}

func (cm *ConnectionManager) Close() error {
	if cm.conn != nil {
		return cm.conn.Close()
	}
	return nil
}

func (cm *ConnectionManager) GetConnection() *amqp.Connection {
	return cm.conn
}

func (cm *ConnectionManager) HealthCheck(ctx context.Context) error {
	if cm.conn == nil || cm.conn.IsClosed() {
		return fmt.Errorf("connection is closed")
	}
	return nil
}
