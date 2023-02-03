package database

import (
	"github.com/Kdsingh333/week1-GL1-CipherSchools/models"
	"gorm.io/gorm"
)
// creating connection and interacting from goland app to db server via db variable

func GetBooks(db *gorm.DB) ([]models.Book,error){
	books := []models.Book{}
	query := db.Select("books.*")
	err:= query.Find(&books).Error
	if err!= nil{
		return nil,err
	}
	return books,nil
}

func GetBooksById(db *gorm.DB, id string)(*models.Book,error){
	books:= models.Book{}
	err := db.Select("books.*").Group("books.id").Where("books.id=?",id).First(&books).Error
	if err!= nil{
		return nil,err
	}
	return &books,nil
}
func DeleteBooksById(db *gorm.DB, id string)(error){
	var book models.Book
	err:= db.Where("id=?",id).Delete(&book).Error
	if err!= nil{
		return err
	}
	return nil
}
func UpdateBooksById(db *gorm.DB, book *models.Book)(error){
	err := db.Save(book).Error 
	if err!= nil{
		return err
	}
	return nil
}
func SaveBook(db *gorm.DB, book *models.Book)(error){
	
	err := db.Save(book).Error
	if err!= nil{
		return err
	}
	return nil
}
