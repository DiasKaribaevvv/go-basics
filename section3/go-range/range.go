package main

import "fmt"

var names []string = []string{"dias", "karibaev"}

func main() {
	for i, v := range names {
		fmt.Println(i, v)
	}
}
