package order

type OrderSender interface {
	Send(*Order) error
}
