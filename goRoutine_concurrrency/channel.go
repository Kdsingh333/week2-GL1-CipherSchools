package main

import (
	"time"

	"fmt"

	"golang.org/x/text/message"
)

func main() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "2"
		fmt.Println(1)
	}(channel)

	message := <-channel
	time.Sleep(time.Second*5)
	fmt.Println(message)
}

func main1(){
	channel := make(chan string,1)
	go func(ch <-chan string){
		mess := <- ch
		fmt.Print(mess)
		fmt.Println(1)
	}(channel)
	message := "Hello from Main function"

	channel <- message

	fmt.Println(message)
}