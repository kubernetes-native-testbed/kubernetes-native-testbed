package main

import "time"

type Point struct {
	UUID        string `gorm:"primary_key"`
	UserUUID    string
	Balance     int32
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
