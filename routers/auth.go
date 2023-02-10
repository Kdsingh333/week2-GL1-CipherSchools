package routers

import (
	"github.com/Kdsingh333/week1-GL1-CipherSchools/handler"
	"github.com/gin-gonic/gin"
)


func AuthRouter(router *gin.Engine, api handler.Handler) {

	router.POST("/login", api.SignIn) // set the function for http://localhost:8080/books: get request
	//while calling handler function, gin will pass *gin.context in the handler function
	router.POST("/register", api.SignUp)

}