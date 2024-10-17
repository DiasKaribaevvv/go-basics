package main

import "fmt"

func main() {
	var myArray []int

	for i := 0; i < 10; i++ {
		myArray = append(myArray, i)
	}

	fmt.Println(myArray)
}
