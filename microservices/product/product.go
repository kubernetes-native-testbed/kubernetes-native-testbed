package main

import "time"

type product struct {
	UUID      string `gorm:"primary_key"`
	Name      string
	Price     uint64
	ImageURLs []productImage `gorm:"foreignkey:ProductUUID;association_foreignkey:UUID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type productImage struct {
	ProductUUID string `gorm:"primary_key"`
	URL         string `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
