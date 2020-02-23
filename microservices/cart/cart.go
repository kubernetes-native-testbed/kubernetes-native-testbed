package main

import "time"

type Cart struct {
	UUID         string        `gorm:"primary_key"`
	CartProducts []CartProduct `gorm:"foreignkey:CartUUID;association_foreignkey:UUID"`
	UserUUID     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

type CartProduct struct {
	CartUUID    string `gorm:"primary_key"`
	ProductUUID string `gorm:"primary_key"`
	Count       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
