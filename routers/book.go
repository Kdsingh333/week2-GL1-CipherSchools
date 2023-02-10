package routers

import (
	"github.com/Kdsingh333/week1-GL1-CipherSchools/database"
	"github.com/Kdsingh333/week1-GL1-CipherSchools/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine{
	router:=gin.Default() //get the default engine for further customization
	api := handler.Handler{
       DB: database.GetDB(), // set the handler db
     }
	 router.GET("/books",api.GetBooks) // set the function for http://localhost:8080/books: get request
     //while calling handler function, gin will pass *gin.context in the handler function
	router.POST("/books",api.SaveBook)
	 return router
}