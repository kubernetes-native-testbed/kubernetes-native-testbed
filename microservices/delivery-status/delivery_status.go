package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

const (
	Waiting = Status(iota + 1)
	Preparing
	Shipping
	Delivered
)

var StatusMap = map[Status]string{
	Waiting:   "Waiting",
	Preparing: "Preparing",
	Shipping:  "Shipping",
	Delivered: "Delivered",
}

type Status int

type DeliveryStatus struct {
	OrderUUID     string     `json:"order_uuid"`
	Status        Status     `json:"status"`
	InquiryNumber string     `json:"inquiry_number"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

func (ds *DeliveryStatus) String() string {
	b, _ := json.Marshal(ds)
	return string(b)
}

func generateInquiryNumber() string {
	// format: 000-0000000-0000000
	n1 := rand.Intn(1000)
	n2 := rand.Intn(10000000)
	n3 := rand.Intn(10000000)
	return fmt.Sprintf("%03d-%07d-%07d", n1, n2, n3)
}
