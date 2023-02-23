package database

import (
	"github.com/bthornhill123/go-auth-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	logger := logger.New(nil, logger.Config{
		IgnoreRecordNotFoundError: true,
	})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{})
	DB = db
}
