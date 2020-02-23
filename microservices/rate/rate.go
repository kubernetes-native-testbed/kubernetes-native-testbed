package main

import "time"

type Rate struct {
	UUID        string `gorm:"primary_key"`
	Rating      int32
	UserUUID    string
	CommentUUID string
	ProductUUID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
