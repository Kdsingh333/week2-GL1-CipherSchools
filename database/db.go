package database

import (
	"log"

	"github.com/Kdsingh333/week1-GL1-CipherSchools/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB
func GetDB() *gorm.DB{
	return DB
}
func Setup() {
	dsn := "host=localhost user=postgres password=postgres dbname=book port=5432 sslmode=disable "
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil{
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB= db
}