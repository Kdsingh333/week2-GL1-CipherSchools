package main

import (
	"fmt"
	"time"
)

func main() {
	helloCh := make(chan string, 1)
	goodByeCh := make(chan string, 1)
	quitCh := make(chan bool)

	
	go receiveMessage(helloCh, goodByeCh, quitCh)
	go sendMessage(helloCh, "Hello World")
	time.Sleep(time.Second)
	go sendMessage(helloCh,"Good Bye world")
	result := <-quitCh
	fmt.Println("message from quitCh", result)
}

func sendMessage(ch chan<- string, message string) {
	ch <- message
}

func receiveMessage(helloCh, goodByeCh <-chan string, quitch chan<- bool) {
    for{
		select{
		case message:=<-helloCh :
			fmt.Println("MEssage from helloCh:",message)
		case message:=<-goodByeCh :
			fmt.Println("MEssage from goodByCh:",message)
		case <-time.After(time.Second*2) :
			fmt.Println("Nothing receiving in 2 second ")
			quitch<-true
			break
	

		}
	}
}
