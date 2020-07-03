package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("postgres", "host=database-1.c1fzlyaabojw.eu-central-1.rds.amazonaws.com port=5432 user=postgres dbname=musicstore password=zw345b7u sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Track{})

	DB = db
}
