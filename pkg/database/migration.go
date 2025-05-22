package database

import (
	"log"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("migration completed")
}
