package main

import (
	"fmt"
	"log"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type orderQueue interface {
	subscribe() (<-chan string, error)
	unsubscribe() error
}

type orderQueueKafka struct {
	consumer *kafka.Consumer
	topic    string
	retry    int
}

func (oq *orderQueueKafka) subscribe() (<-chan string, error) {
	orderCh := make(chan string, 64)
	err := oq.consumer.Subscribe(oq.topic, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			msg, err := oq.consumer.ReadMessage(-1)
			if err == nil {
				if data := string(msg.Value); data != "" {
					orderCh <- data
				}
			} else {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}()

	return orderCh, nil
}

func (oq *orderQueueKafka) unsubscribe() error {
	return oq.consumer.Unsubscribe()
}

type orderQueueKafkaConfig struct {
	host     string
	port     int
	username string
	password string
	topic    string
	group    string
	retry    int
}

func (c *orderQueueKafkaConfig) connect() (orderQueue, func() error, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%d", c.host, c.port),
		"sasl.username":     c.username,
		"sasl.password":     c.password,
		"group.id":          c.group,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, nil, err
	}

	return &orderQueueKafka{consumer: consumer, topic: c.topic, retry: c.retry}, consumer.Close, nil
}
