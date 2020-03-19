package order

import (
	"encoding/json"
	"time"
)

type Order struct {
	UUID            string           `gorm:"primary_key" json:"uuid"`
	OrderedProducts []OrderedProduct `gorm:"foreignkey:OrderUUID;association_foreignkey:UUID" json:"ordered_products"`
	UserUUID        string           `json:"user_uuid"`
	PaymentInfoUUID string           `json:"payment_info_uuid"`
	AddressUUID     string           `json:"address_uuid"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DeletedAt       *time.Time       `sql:"index" json:"deleted_at"`
}

type OrderedProduct struct {
	OrderUUID   string     `gorm:"primary_key" json:"order_uuid"`
	ProductUUID string     `gorm:"primary_key" json:"product_uuid"`
	Count       int        `json:"count"`
	Price       int        `json:"price"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
}

func (o *Order) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (op *OrderedProduct) String() string {
	b, _ := json.Marshal(op)
	return string(b)
}
