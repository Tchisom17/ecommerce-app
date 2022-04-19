package database

import (
	"ecommerce-app/internal/core/domain"
	"ecommerce-app/internal/ports"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type datastore struct {
}

// NewDatabase creates a new instance for managing database
func NewDatabase() ports.Db {
	return &datastore{}
}

func (d *datastore) ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(
			err.Error(),
		)
		panic(any("failed to connect database"))
	}

	fmt.Println("Established database connection")

	return db
}

func (d *datastore) MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&domain.User{}, &domain.Item{}, &domain.Role{}, &domain.Category{}, &domain.Blacklist{}, &domain.Images{})
}
