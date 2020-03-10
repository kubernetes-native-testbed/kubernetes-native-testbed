package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type userRepository interface {
	findByUUID(string) (*User, error)
	store(*User) (string, error)
	update(*User) error
	deleteByUUID(string) error
	initDB() error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func (rr *userRepositoryImpl) findByUUID(uuid string) (*User, error) {
	p := &User{UUID: uuid}
	if err := rr.db.Preload("Addresses").Find(p).Error; err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}
	return p, nil
}

func (rr *userRepositoryImpl) store(u *User) (string, error) {
	if !rr.db.NewRecord(u) {
		return "", fmt.Errorf("store error: this key already exists")
	}

	u.UUID = uuid.New().String()
	if err := rr.db.Create(u).Error; err != nil {
		return "", fmt.Errorf("store error: %w (user: %v)", err, u)
	}

	return u.UUID, nil
}

func (rr *userRepositoryImpl) update(u *User) error {
	if err := rr.db.Save(u).Error; err != nil {
		return fmt.Errorf("update error: %w (user: %v)", err, u)
	}
	return nil
}

func (rr *userRepositoryImpl) deleteByUUID(uuid string) error {
	if err := rr.db.Delete(&User{UUID: uuid}).Error; err != nil {
		return fmt.Errorf("deleteByID error: %w (uuid: %s)", err, uuid)
	}
	return nil
}

func (rr *userRepositoryImpl) initDB() error {
	if err := rr.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Address{}).Error; err != nil {
		return err
	}
	return nil
}
