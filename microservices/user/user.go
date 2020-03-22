package user

import "time"

type User struct {
	UUID                   string     `json:"uuid" gorm:"primary_key"`
	Username               string     `json:"username"`
	FirstName              string     `json:"first_name"`
	LastName               string     `json:"last_name"`
	Age                    int32      `json:"age"`
	Addresses              []Address  `json:"addresses" gorm:"foreignkey:UserUUID;association_foreignkey:UUID"`
	PasswordHash           string     `json:"password_hash"`
	DefaultPaymentInfoUUID string     `json:"default_paymentinfo_uuid"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
	DeletedAt              *time.Time `json:"deleted_at" sql:"index"`
}

type Address struct {
	UUID        string     `json:"uuid" gorm:"primary_key"`
	UserUUID    string     `json:"user_uuid"`
	ZipCode     string     `json:"zipcode"`
	Country     string     `json:"country"`
	State       string     `json:"state"`
	City        string     `json:"city"`
	AddressLine string     `json:"address_line"`
	Disabled    bool       `json:"disabled"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" sql:"index"`
}
