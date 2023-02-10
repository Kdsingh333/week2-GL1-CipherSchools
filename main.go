package main

import (
	"fmt"
	"log"

	"github.com/Kdsingh333/week1-GL1-CipherSchools/database"
	"github.com/Kdsingh333/week1-GL1-CipherSchools/handler"
	"github.com/Kdsingh333/week1-GL1-CipherSchools/routers"
	"github.com/gin-gonic/gin"
)

func init(){
	database.Setup() //establish the database connection
}
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func AuthMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL)
		token := c.GetHeader("token")
		fmt.Println("got token:	" + token)
		isValid, err := handler.ValidateToken(token)
		if err != nil && !isValid {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}
func main() {
	engine:=gin.Default() //get the default engine for further customization
	api := handler.Handler{
       DB: database.GetDB(), // set the handler db
     }
	 routers.BookRouter(engine,api)
	 routers.AuthRouter(engine,api)
	err := engine.Run("127.0.0.1:8080") // start the engine
	if err != nil{
		log.Fatal(err)
	}
}
