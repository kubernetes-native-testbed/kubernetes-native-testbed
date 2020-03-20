package product

import (
	"encoding/json"
	"time"
)

type Product struct {
	UUID      string `gorm:"primary_key"`
	Name      string
	Price     uint64
	ImageURLs []ProductImage `gorm:"foreignkey:ProductUUID;association_foreignkey:UUID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type ProductImage struct {
	ProductUUID string `gorm:"primary_key"`
	URL         string `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

func (p *Product) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func (pi *ProductImage) String() string {
	b, _ := json.Marshal(pi)
	return string(b)
}
