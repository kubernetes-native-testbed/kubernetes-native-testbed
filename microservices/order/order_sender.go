package order

type orderSender interface {
	send(*Order) error
}
