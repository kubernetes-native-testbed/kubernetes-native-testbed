package main

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

type orderSender interface {
	send(*Order) error
}

type orderSenderImpl struct {
	conn    *nats.Conn
	subject string
	retry   int
}

func (os *orderSenderImpl) send(o *Order) error {
	var err error
	for i := 0; i < os.retry; i++ {
		if err = os.conn.Publish(os.subject, []byte(o.String())); err == nil {
			log.Printf("send [%s] %s", os.subject, o)
			break
		}
	}
	return err
}
