package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	jsonFile, err := os.Open("something.json")
	if err!= nil{
		log.Fatal(err)
	}
	
	defer jsonFile.Close()
	jsonByteValues,err := ioutil.ReadAll(jsonFile)
	if err!= nil{
		log.Fatal(err)
	}
	var something something
	json.Unmarshal(jsonByteValues,&something)
	log.Println(something)
}