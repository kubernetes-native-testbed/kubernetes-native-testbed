package order

type orderSenderKafka struct{}

func (os *orderSenderKafka) Send(o *Order) error {
	return nil
}

type OrderSenderKafkaConfig struct{}

func (c *OrderSenderKafkaConfig) Connect() (OrderSender, func(), error) {
	return nil, nil, nil
}
