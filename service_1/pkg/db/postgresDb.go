package db

import (
	"log"

	"github.com/ajalck/service_1/pkg/config"
	"github.com/ajalck/service_1/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func ConnectPostgres(config *config.Config) *UserDB {
	db, err := gorm.Open(postgres.Open(config.SqlDSN), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return &UserDB{DB: db}
}
func SyncDB(db *gorm.DB) (err error) {
	if err = (db.AutoMigrate(&models.User{})); err != nil {
		return err
	}
	return nil
}
