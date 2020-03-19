package order

import (
	"fmt"
	"log"

	nats "github.com/nats-io/nats.go"
)

type orderSenderNATS struct {
	conn    *nats.Conn
	subject string
	retry   int
}

func (os *orderSenderNATS) Send(o *Order) error {
	var err error
	for i := 0; i < os.retry; i++ {
		if err = os.conn.Publish(os.subject, []byte(o.UUID)); err == nil {
			log.Printf("send [%s] %s", os.subject, o)
			break
		}
	}
	return err
}

type OrderSenderNATSConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Subject  string
	Retry    int
}

func (c *OrderSenderNATSConfig) Connect() (OrderSender, func(), error) {
	conn, err := nats.Connect(
		fmt.Sprintf("%s:%d", c.Host, c.Port),
		nats.UserInfo(c.Username, c.Password),
	)
	if err != nil {
		return nil, nil, err
	}
	return &orderSenderNATS{conn: conn, subject: c.Subject, retry: c.Retry}, conn.Close, nil
}
