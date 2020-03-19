package order

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	FindByUUID(string) (*Order, error)
	Store(*Order) (string, error)
	Update(*Order) error
	DeleteByUUID(string) error

	InitDB() error
}

type orderRepositoryTiDB struct {
	db *gorm.DB
}

func (or *orderRepositoryTiDB) FindByUUID(uuid string) (*Order, error) {
	o := &Order{UUID: uuid}
	if err := or.db.Preload("OrderedProducts").Find(o).Error; err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}
	return o, nil
}

func (or *orderRepositoryTiDB) Store(o *Order) (string, error) {
	if !or.db.NewRecord(o) {
		return "", fmt.Errorf("store error: this key already exists")
	}

	o.UUID = uuid.New().String()
	if err := or.db.Create(o).Error; err != nil {
		return "", fmt.Errorf("store error: %w (order: %v)", err, o)
	}

	return o.UUID, nil
}

func (or *orderRepositoryTiDB) Update(o *Order) error {
	if err := or.db.Save(o).Error; err != nil {
		return fmt.Errorf("update error: %w (order: %v)", err, o)
	}
	return nil
}

func (or *orderRepositoryTiDB) DeleteByUUID(uuid string) error {
	if err := or.db.Delete(&Order{UUID: uuid}).Error; err != nil {
		return fmt.Errorf("deleteByID error: %w (uuid: %s)", err, uuid)
	}
	return nil
}

func (or *orderRepositoryTiDB) InitDB() error {
	if err := or.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Order{}, &OrderedProduct{}).Error; err != nil {
		return err
	}
	return nil
}

type OrderRepositoryTiDBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func (c *OrderRepositoryTiDBConfig) Connect() (OrderRepository, func() error, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("%w (dsn=%s)", err, dsn)
	}

	return &orderRepositoryTiDB{db: db}, db.Close, nil
}
