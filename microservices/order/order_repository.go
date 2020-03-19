package order

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type orderRepository interface {
	findByUUID(string) (*Order, error)
	store(*Order) (string, error)
	update(*Order) error
	deleteByUUID(string) error

	initDB() error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func (or *orderRepositoryImpl) findByUUID(uuid string) (*Order, error) {
	o := &Order{UUID: uuid}
	if err := or.db.Preload("OrderedProducts").Find(o).Error; err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}
	return o, nil
}

func (or *orderRepositoryImpl) store(o *Order) (string, error) {
	if !or.db.NewRecord(o) {
		return "", fmt.Errorf("store error: this key already exists")
	}

	o.UUID = uuid.New().String()
	if err := or.db.Create(o).Error; err != nil {
		return "", fmt.Errorf("store error: %w (order: %v)", err, o)
	}

	return o.UUID, nil
}

func (or *orderRepositoryImpl) update(o *Order) error {
	if err := or.db.Save(o).Error; err != nil {
		return fmt.Errorf("update error: %w (order: %v)", err, o)
	}
	return nil
}

func (or *orderRepositoryImpl) deleteByUUID(uuid string) error {
	if err := or.db.Delete(&Order{UUID: uuid}).Error; err != nil {
		return fmt.Errorf("deleteByID error: %w (uuid: %s)", err, uuid)
	}
	return nil
}

func (or *orderRepositoryImpl) initDB() error {
	if err := or.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Order{}, &OrderedProduct{}).Error; err != nil {
		return err
	}
	return nil
}
