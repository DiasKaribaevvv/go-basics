package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "Hello"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "World"
	}()

	select {
	case ms1 := <-chan1:
		fmt.Println(ms1)
	case ms2 := <-chan2:
		fmt.Println(ms2)
	}

}
