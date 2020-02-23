package main

import "time"

type DeliveryStatus struct {
	UUID          string `gorm:"primary_key"`
	OrderUUID     string
	Status        int // enum
	InquiryNumber string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}
