package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type paymentInfoRepository interface {
	findByUUID(string) (*PaymentInfo, error)
	store(*PaymentInfo) (string, error)
	update(*PaymentInfo) error
	deleteByUUID(string) error

	initDB() error
}

type paymentInfoRepositoryImpl struct {
	db *gorm.DB
}

func (pir *paymentInfoRepositoryImpl) findByUUID(uuid string) (*PaymentInfo, error) {
	pi := &PaymentInfo{UUID: uuid}
	if err := pir.db.Find(pi).Error; err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}
	return pi, nil
}

func (pir *paymentInfoRepositoryImpl) store(pi *PaymentInfo) (string, error) {
	if !pir.db.NewRecord(pi) {
		return "", fmt.Errorf("store error: this key already exists")
	}

	pi.UUID = uuid.New().String()
	if err := pir.db.Create(pi).Error; err != nil {
		return "", fmt.Errorf("store error: %w (paymentInfo: %v)", err, pi)
	}

	return pi.UUID, nil
}

func (pir *paymentInfoRepositoryImpl) update(pi *PaymentInfo) error {
	if err := pir.db.Save(pi).Error; err != nil {
		return fmt.Errorf("update error: %w (paymentInfo: %v)", err, pi)
	}
	return nil
}

func (pir *paymentInfoRepositoryImpl) deleteByUUID(uuid string) error {
	if err := pir.db.Delete(&PaymentInfo{UUID: uuid}).Error; err != nil {
		return fmt.Errorf("deleteByID error: %w (uuid: %s)", err, uuid)
	}
	return nil
}

func (pir *paymentInfoRepositoryImpl) initDB() error {
	if err := pir.db.AutoMigrate(&PaymentInfo{}).Error; err != nil {
		return err
	}
	return nil
}
