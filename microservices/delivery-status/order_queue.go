package main

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

type orderQueue interface {
	subscribe() (<-chan string, error)
}

type orderQueueNATS struct {
	conn    *nats.Conn
	subject string
	retry   int
}

func (oq *orderQueueNATS) subscribe() (<-chan string, error) {
	msgCh := make(chan *nats.Msg, 64)
	orderCh := make(chan string, 64)
	if _, err := oq.conn.ChanSubscribe(oq.subject, msgCh); err != nil {
		return nil, err
	}
	go func() {
		for {
			msg := <-msgCh
			orderCh <- msg.Reply
		}
	}()
	return orderCh, nil
}

type orderQueueNATSConfig struct {
	host     string
	port     int
	username string
	password string
	subject  string
	retry    int
}

func (c *orderQueueNATSConfig) connect() (orderQueue, func(), error) {
	conn, err := nats.Connect(
		fmt.Sprintf("%s:%d", c.host, c.port),
		nats.UserInfo(c.username, c.password),
	)
	if err != nil {
		return nil, nil, err
	}
	return &orderQueueNATS{conn: conn, subject: c.subject, retry: c.retry}, conn.Close, nil
}
