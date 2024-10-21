package main

import "fmt"

func main() {
	var myInterface interface{} = "Dias"

	s := myInterface.(string)

	s, ok := myInterface.(string)

	fmt.Println(s, ok)
}
