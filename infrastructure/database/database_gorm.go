package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm interface {
	Create(data interface{}, table string) (interface{}, error)
	AutoMigrate(entities ...interface{}) error
}

type gormHandler struct {
	db *gorm.DB
}

func NewGormHandler(c *config) (*gormHandler, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		// "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC/GMT",
		c.host,
		c.user,
		c.password,
		c.database,
		c.port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &gormHandler{}, err
	}

	return &gormHandler{db}, nil
}

func (gorm *gormHandler) AutoMigrate(entities ...interface{}) error {
	if err := gorm.db.AutoMigrate(entities...); err != nil {
		return err
	} else {
		return nil
	}
}

var (
	errFailedSaveData = errors.New("failed to save data to database")
)

func (gorm *gormHandler) Create(data interface{}, table string) (interface{}, error) {
	result := gorm.db.Table(table).Create(data)

	if result.RowsAffected < 1 {
		return nil, errFailedSaveData
	}

	return data, nil
}
