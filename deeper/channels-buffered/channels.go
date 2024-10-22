package main

import (
	"fmt"
	"math/rand"
)

func main() {
	myChannel := make(chan string)
	numCount := 5
	go scoreCounter(myChannel, numCount)
	// for message := range myChannel {

	// 	fmt.Println(message)
	// }

	for {
		message, open := <-myChannel
		if !open {
			break
		}
		fmt.Println(message)
	}
	fmt.Println(<-myChannel)
}

func scoreCounter(channel chan string, num int) {

	for i := 0; i < num; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprintf("Your score is: %v", score)
	}
	close(channel)
}
