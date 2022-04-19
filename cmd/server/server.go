package server

import (
	"ecommerce-app/internal/ports"
	"ecommerce-app/pkg/config"
	"gorm.io/gorm"
	"log"
)

// DBConnection stores the instance of the Database
var DBConnection *gorm.DB

// Run function starts the database connection
func Run(database ports.Db) error {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
		return err
	}

	DBConnection = database.ConnectDB(config.Instance.DatabaseURL)
	err = database.MigrateAll(DBConnection)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
