package handler

import (
	"log"
	"net/http"

	"github.com/Kdsingh333/week1-GL1-CipherSchools/database"
	"github.com/Kdsingh333/week1-GL1-CipherSchools/models"
	"github.com/gin-gonic/gin"
	
)



func (h *Handler)GetBooks(in *gin.Context) {
      books,err := database.GetBooks(h.DB)

	  if err!=nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError,err)
	  }

	  in.JSON(200,books)
}
func (h *Handler)SaveBook(in *gin.Context) {
	book := models.Book{}
	 err :=in.BindJSON(&book)
	 if err!= nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError,err)
	 }
     err = database.SaveBook(h.DB,&book)
	 if err != nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError,err)
	 }
	 in.JSON(http.StatusOK,book)
}