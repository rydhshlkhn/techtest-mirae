package infra

import (
	"log"
	"os"

	"github.com/rydhshlkhn/techtest-mirae/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func IniitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&domain.Stock{})

	return db
}
