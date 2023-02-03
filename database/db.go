package database

import (
	"log"
	"github.com/Suryaakula888/models"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"

)

var DB *gorm.DB

func GetDB() *gorm.DB{
	return DB
}

func setup() {
	host := "localhost"
	port := "5432"
	dbname := "book"
	username := "postgres"
	passward := "postgres"
	args:="host=" + host + "port=" + port + "user=" + username + "dbname=" + dbname + "sslmode=display.passward=" + passward 
	db,err:= gorm.Open("postgres" ,args)
	if err != nil{
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}

