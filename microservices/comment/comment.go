package main

import "time"

type Comment struct {
	UUID          string `gorm:"primary_key"`
	Message       string
	UserUUID      string
	ParentComment []Comment `gorm:"many2many:friendships;association_jointable_foreignkey:parend_comment_uuid"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}
