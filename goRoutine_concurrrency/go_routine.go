package main

import (
	"fmt"
	"sync"
	
)


// var wait sync.WaitGroup
func main() {
	// wait.Add(1)
	// go hello()
	// go hello()
	// wait.Wait()
	var wait sync.WaitGroup
	counter := 200
	wait.Add(counter)

	for i:=0; i<counter; i++{
		go hello(&wait,i)
	}
	defer wait.Wait()
	
}



func hello(wait *sync.WaitGroup,counter int) {
	fmt.Println(counter)
	wait.Done()
}
// func main() {
// 	go hello()
// 	time.Sleep(time.Second*5)
// }



// func hello() {
// 	fmt.Println(1)
// }



// go func(){
// 	fmt.Println(10)
// 	fmt.Println(10)
// 	fmt.Println(10)
// 	fmt.Println(10)
// 	fmt.Println(10)
// }()