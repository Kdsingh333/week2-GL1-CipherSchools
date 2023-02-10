package routers

import (
	"github.com/Kdsingh333/week1-GL1-CipherSchools/handler"
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine, api handler.Handler) {

	 router.GET("/books",api.GetBooks) // set the function for http://localhost:8080/books: get request
     //while calling handler function, gin will pass *gin.context in the handler function
	router.POST("/books",api.SaveBook)
	
}