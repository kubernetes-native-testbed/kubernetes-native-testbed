package main

import "time"

type Point struct {
	UUID      string `gorm:"primary_key"`
	UserUUID  string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
