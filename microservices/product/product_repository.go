package product

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	FindByUUID(string) (*Product, error)
	Store(*Product) (string, error)
	Update(*Product) error
	DeleteByUUID(string) error

	InitDB() error
}

type productRepositoryMySQL struct {
	db *gorm.DB
}

func (pr *productRepositoryMySQL) FindByUUID(uuid string) (*Product, error) {
	p := &Product{UUID: uuid}
	if err := pr.db.Preload("ImageURLs").Find(p).Error; err != nil {
		return nil, fmt.Errorf("findByID error: %w (uuid: %s)", err, uuid)
	}
	return p, nil
}

func (pr *productRepositoryMySQL) Store(p *Product) (string, error) {
	if !pr.db.NewRecord(p) {
		return "", fmt.Errorf("store error: this key already exists")
	}

	p.UUID = uuid.New().String()
	if err := pr.db.Create(p).Error; err != nil {
		return "", fmt.Errorf("store error: %w (product: %v)", err, p)
	}

	return p.UUID, nil
}

func (pr *productRepositoryMySQL) Update(p *Product) error {
	if err := pr.db.Save(p).Error; err != nil {
		return fmt.Errorf("update error: %w (product: %v)", err, p)
	}
	return nil
}

func (pr *productRepositoryMySQL) DeleteByUUID(uuid string) error {
	if err := pr.db.Delete(&Product{UUID: uuid}).Error; err != nil {
		return fmt.Errorf("deleteByID error: %w (uuid: %s)", err, uuid)
	}
	return nil
}

func (pr *productRepositoryMySQL) InitDB() error {
	if err := pr.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Product{}, &ProductImage{}).Error; err != nil {
		return err
	}
	return nil
}

type ProductRepositoryMySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func (c *ProductRepositoryMySQLConfig) Connect() (ProductRepository, func() error, error) {
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

	return &productRepositoryMySQL{db: db}, db.Close, nil
}
