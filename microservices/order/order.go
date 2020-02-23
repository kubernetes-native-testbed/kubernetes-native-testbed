package main

import "time"

type Order struct {
	UUID            string `gorm:"primary_key"`
	OrderedProducts []OrderedProduct
	UserUUID        string
	PaymentInfoUUID string
	AddressUUID     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

type OrderedProduct struct {
	OrderUUID   string
	ProductUUID string
	Count       int
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
