package user

import "time"

type User struct {
	UUID                   string `gorm:"primary_key"`
	Username               string
	FirstName              string
	LastName               string
	Age                    int32
	Addresses              []Address `gorm:"foreignkey:UserUUID;association_foreignkey:UUID"`
	PasswordHash           string
	DefaultPaymentInfoUUID string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              *time.Time `sql:"index"`
}

type Address struct {
	UUID        string `gorm:"primary_key"`
	UserUUID    string
	ZipCode     string
	Country     string
	State       string
	City        string
	AddressLine string
	Disabled    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
