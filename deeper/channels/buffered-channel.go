package main

import "fmt"

func Buffered() {
	newChannel := make(chan string, 2)
	newChannel <- "Hello"
	newChannel <- "World!"
	fmt.Println(<-newChannel)
	fmt.Println(<-newChannel)
}
