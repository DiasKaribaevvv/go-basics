package main

import (
	"fmt"
)

/*
Variables with initializers

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/

var a, b, c, d bool

var i, j int = 1, 2

func main() {
	var num int

	var python, golang, java = true, false, 5.5

	fmt.Println(a, b, c, d, num, i, j, python, golang, java)

}
