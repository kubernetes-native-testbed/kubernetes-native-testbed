package main

import "time"

type PaymentInfo struct {
	UUID           string `gorm:"primary_key"`
	UserUUID       string
	Name           string
	CardNumber     string
	SecurityCode   string
	ExpirationDate string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}
