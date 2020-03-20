package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type pointRepository interface {
	findByUUID(string) (*Point, error)
	store(*Point) (string, error)
	update(*Point) error
	deleteByUUID(string) error

	getAmount(string) (int32, error)

	initDB() error
}

type pointRepositoryImpl struct {
	db *gorm.DB
}

func (pr *pointRepositoryImpl) findByUUID(uuid string) (*Point, error) {
	p := &Point{UUID: uuid}
	if err := pr.db.Find(p).Error; err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}
	return p, nil
}

func (pr *pointRepositoryImpl) store(p *Point) (string, error) {
	if !pr.db.NewRecord(p) {
		return "", fmt.Errorf("store error: this key already exists")
	}

	p.UUID = uuid.New().String()
	if err := pr.db.Create(p).Error; err != nil {
		return "", fmt.Errorf("store error: %w (point: %v)", err, p)
	}

	return p.UUID, nil
}

func (pr *pointRepositoryImpl) update(p *Point) error {
	if err := pr.db.Save(p).Error; err != nil {
		return fmt.Errorf("update error: %w (point: %v)", err, p)
	}
	return nil
}

func (pr *pointRepositoryImpl) deleteByUUID(uuid string) error {
	if err := pr.db.Delete(&Point{UUID: uuid}).Error; err != nil {
		return fmt.Errorf("deleteByID error: %w (uuid: %s)", err, uuid)
	}
	return nil
}

func (pr *pointRepositoryImpl) getAmount(useruuid string) (int32, error) {
	rows, err := pr.db.Model(&Point{}).Where("UserUUID = ?", useruuid).Select("Balance").Rows() // (*sql.Rows, error)
	defer rows.Close()
	if err != nil {
		return 0, fmt.Errorf("getAmount error: %w (useruuid: %s)", err, useruuid)
	}

	amount := int32(0)
	for rows.Next() {
		var point Point
		pr.db.ScanRows(rows, &point)
		amount += point.Balance
	}

	return amount, nil
}

func (pr *pointRepositoryImpl) initDB() error {
	if err := pr.db.AutoMigrate(&Point{}).Error; err != nil {
		return err
	}
	return nil
}
