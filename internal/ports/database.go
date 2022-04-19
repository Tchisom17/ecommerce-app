package ports

import "gorm.io/gorm"

// Database defines the interface for database management
type Db interface {
	ConnectDB(url string) *gorm.DB
	MigrateAll(db *gorm.DB) error
}
