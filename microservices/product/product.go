package main

import "github.com/jinzhu/gorm"

type product struct {
	gorm.Model
	UUID      string `gorm:"primary_key"`
	Name      string
	Price     int
	ImageURLs []string
}
