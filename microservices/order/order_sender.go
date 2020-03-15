package main

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

type orderSender interface {
	send(*Order) error
}

type orderSenderImpl struct {
	conn *nats.Conn
}

func (os *orderSenderImpl) send(o *Order) error {
	log.Printf("send %s", o)
	return nil
}
