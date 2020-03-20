package order

type OrderSender interface {
	Send(*Order, Operation) error
}

type Operation string

const (
	CreateOperation Operation = "create"
	UpdateOperation Operation = "update"
	DeleteOperation Operation = "delete"
)
