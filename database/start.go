package database

import (
	"fmt"
	"mygram-api/config"
	"mygram-api/models/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Start() (Database, error) {
	dbInfo := config.GetDatabaseEnv()

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.Name)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		fmt.Println("error open connection to db", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(&entity.User{}, &entity.Photo{})
	if err != nil {
		fmt.Println("error on migration", err)
		return Database{}, err
	}

	return Database{
		db,
	}, nil
}
