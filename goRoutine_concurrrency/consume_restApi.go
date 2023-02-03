package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Data struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Year         int    `json:"year"`
		Color        string `json:"color"`
		PantoneValue string `json:"pantone_value"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func main() {
	url := "https://reqres.in/api/products/3"
	response, err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}

	data,err := ioutil.ReadAll(response.Body)
	// fmt.Println(string(data))
	var responseObject Response
	json.Unmarshal(data,&responseObject)
	log.Println(responseObject)
 }