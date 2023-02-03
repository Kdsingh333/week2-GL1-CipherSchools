package main

import (
	"log"

	"github.com/Kdsingh333/week1-GL1-CipherSchools/database"
	"github.com/Kdsingh333/week1-GL1-CipherSchools/routers"
)

func init(){
	database.Setup() //establish the database connection
}

func main() {
	engine := routers.Router() //get the customized engine
	err := engine.Run("127.0.0.1:8080") // start the engine
	if err != nil{
		log.Fatal(err)
	}
}
