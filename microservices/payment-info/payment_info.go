package main

import "time"

type PaymentInfo struct {
	UUID        string `gorm:"primary_key"`
	UserUUID    string
	CardNumber  string
	CommentUUID string
	ProductUUID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
