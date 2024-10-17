package main

import "fmt"

/*
Constants

Constants are declared like variables, but with the
const keyword.

Constants can be character, string, boolean,
or numeric values.

Constants canNOT be declared using the := syntax.
*/

const Pi = 3.14

func main() {
	const World = "hello"
	fmt.Println(Pi, World)
}
