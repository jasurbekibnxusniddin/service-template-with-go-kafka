package postgres

import (
	"log"

	"github.com/jasurbekibnxusniddin/service-template-with-go-kafka/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnToDB() (*gorm.DB, error) {

	dsn := "host=54.93.101.96 user=postgres password=!(((1001#!gO dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		log.Println("error on ConnToDB(). Error:", err)
		return nil, err
	}

	return db, nil
}

func AoutoMigrate(db *gorm.DB) error {

	if err := db.AutoMigrate(
		&models.TodoModel{},
	); err != nil {
		log.Println("error on AoutoMigrate(). Error:", err)
		return err
	}

	return nil
}
