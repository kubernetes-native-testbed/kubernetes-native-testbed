package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type productRepository interface {
	findByUUID(string) (*product, error)
	store(*product) (string, error)
	update(*product) error
	deleteByUUID(string) error

	initDB() error
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func (pr *productRepositoryImpl) findByUUID(uuid string) (*product, error) {
	p := &product{}
	if err := pr.db.Where(&product{UUID: uuid}).First(p).Error; err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}
	return p, nil
}

func (pr *productRepositoryImpl) store(p *product) (string, error) {
	if !pr.db.NewRecord(p) {
		return "", fmt.Errorf("store error: this key already exists")

	}

	p.UUID = uuid.New().String()
	if err := pr.db.Create(p).Error; err != nil {
		return "", fmt.Errorf("store error: %w (product: %v)", err, p)
	}

	return p.UUID, nil
}

func (pr *productRepositoryImpl) update(p *product) error {
	if err := pr.db.Update(p).Error; err != nil {
		return fmt.Errorf("update error: %w (product: %v)", err, p)
	}
	return nil
}

func (pr *productRepositoryImpl) deleteByUUID(uuid string) error {
	if err := pr.db.Delete(&product{UUID: uuid}).Error; err != nil {
		return fmt.Errorf("deleteByID error: %w (uuid: %s)", err, uuid)
	}
	return nil
}

func (pr *productRepositoryImpl) initDB() error {
	if err := pr.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&product{}, &productImage{}).Error; err != nil {
		return err
	}
	return nil
}
