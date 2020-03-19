package main

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

type orderQueue interface {
	subscribe() (<-chan string, error)
	unsubscribe() error
}

type orderQueueNATS struct {
	conn    *nats.Conn
	subject string
	retry   int

	subscription *nats.Subscription
}

func (oq *orderQueueNATS) subscribe() (<-chan string, error) {
	msgCh := make(chan *nats.Msg, 64)
	orderCh := make(chan string, 64)
	sub, err := oq.conn.ChanQueueSubscribe(oq.subject, "delivery-status-group", msgCh)
	if err != nil {
		return nil, err
	}
	oq.subscription = sub

	go func() {
		for {
			msg := <-msgCh
			orderCh <- msg.Reply
		}
	}()
	return orderCh, nil
}

func (oq *orderQueueNATS) unsubscribe() error {
	return oq.subscription.Unsubscribe()
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
