package main

import "fmt"

/*
Functions

A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

(For more about why types look the way they do,
 see the article on Go's declaration syntax.)
*/

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(multiply(5, 2))
}
