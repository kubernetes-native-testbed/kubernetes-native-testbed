package order

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

type orderSenderNATS struct {
	conn    *nats.Conn
	subject string
	retry   int
}

func (os *orderSenderNATS) send(o *Order) error {
	var err error
	for i := 0; i < os.retry; i++ {
		if err = os.conn.Publish(os.subject, []byte(o.UUID)); err == nil {
			log.Printf("send [%s] %s", os.subject, o)
			break
		}
	}
	return err
}
