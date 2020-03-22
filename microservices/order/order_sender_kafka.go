package order

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type orderSenderKafka struct {
	producer *kafka.Producer
	topic    string
	retry    int
}

func (os *orderSenderKafka) Send(o *Order, op Operation) error {
	bytes, err := json.Marshal(o)
	if err != nil {
		return err
	}

	for i := 0; i < os.retry; i++ {
		err = os.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &os.topic, Partition: kafka.PartitionAny},
			Value:          []byte(fmt.Sprintf("%s::%s", op, bytes)),
		}, nil)
		if err == nil {
			log.Printf("send [%s] %s", os.topic, o)
			return nil
		}
	}
	return err
}

type OrderSenderKafkaConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Topic    string
	Retry    int
}

func (c *OrderSenderKafkaConfig) Connect() (OrderSender, func(), error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%d", c.Host, c.Port),
		"sasl.username":     c.Username,
		"sasl.password":     c.Password,
	})
	if err != nil {
		return nil, nil, err
	}

	return &orderSenderKafka{producer: producer, topic: c.Topic, retry: c.Retry}, producer.Close, nil
}
