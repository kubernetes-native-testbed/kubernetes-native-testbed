package order

type orderSenderKafka struct{}

func (os *orderSenderKafka) send(o *Order) error {
	return nil
}
