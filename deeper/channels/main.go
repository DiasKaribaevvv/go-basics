package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %v µs\n", time.Since(now).Microseconds()) // prints in microseconds
	}()
	myChan := make(chan bool)
	go attackEvil("Dias", myChan)
	fmt.Println(<-myChan)
	Buffered()
	time.Sleep(10 * time.Second)

}

func attackEvil(target string, isCompleted chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println("Attacking -->", target)
	isCompleted <- true
}
