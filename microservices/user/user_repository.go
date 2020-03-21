package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	FindByUUID(string) (*User, error)
	Store(*User) (string, error)
	Update(*User) error
	DeleteByUUID(string) error
	InitDB() error
}

type userRepositoryMySQL struct {
	db *gorm.DB
}

func (rr *userRepositoryMySQL) FindByUUID(uuid string) (*User, error) {
	p := &User{UUID: uuid}
	if err := rr.db.Preload("Addresses").Find(p).Error; err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}
	return p, nil
}

func (rr *userRepositoryMySQL) Store(u *User) (string, error) {
	if !rr.db.NewRecord(u) {
		return "", fmt.Errorf("store error: this key already exists")
	}

	u.UUID = uuid.New().String()
	for i := 0; i < len(u.Addresses); i++ {
		u.Addresses[i].UUID = uuid.New().String()
		u.Addresses[i].UserUUID = u.UUID
	}
	if err := rr.db.Create(u).Error; err != nil {
		return "", fmt.Errorf("store error: %w (user: %v)", err, u)
	}

	return u.UUID, nil
}

func (rr *userRepositoryMySQL) Update(u *User) error {
	if err := rr.db.Save(u).Error; err != nil {
		return fmt.Errorf("update error: %w (user: %v)", err, u)
	}
	return nil
}

func (rr *userRepositoryMySQL) DeleteByUUID(uuid string) error {
	if err := rr.db.Delete(&User{UUID: uuid}).Error; err != nil {
		return fmt.Errorf("deleteByID error: %w (uuid: %s)", err, uuid)
	}
	return nil
}

func (rr *userRepositoryMySQL) InitDB() error {
	if err := rr.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Address{}).Error; err != nil {
		return err
	}
	return nil
}

type UserRepositoryMySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func (c *UserRepositoryMySQLConfig) Connect() (UserRepository, func() error, error) {
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

	return &userRepositoryMySQL{db: db}, db.Close, nil
}
