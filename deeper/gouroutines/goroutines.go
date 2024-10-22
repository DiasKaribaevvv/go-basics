package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start).Seconds())
	}()
	var evils []string = []string{"Dias", "Marat", "Adil"}

	for _, value := range evils {
		go attackEvil(value)
	}

	time.Sleep(2 * time.Second)
}

func attackEvil(target string) {
	fmt.Println("Attacking evil named:", target)
	time.Sleep(1 * time.Second)
}
