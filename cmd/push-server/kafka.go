package main

import (
	"context"

	"github.com/gwaylib/errors"
	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	server string
}

func NewKafka(uri string) *Kafka {
	return &Kafka{
		server: uri,
	}
}

func (k *Kafka) dial(topic string) (*kafka.Conn, error) {
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", k.server, topic, partition)
	if err != nil {
		return nil, errors.As(err, topic)
	}
	return conn, nil
}
