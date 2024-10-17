package main

import "fmt"

/*
Short variable declarations

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/

func main() {
	var myNum int = 1
	myNewNum := 1

	boolean, myString := true, "Dias"

	fmt.Println(myNum, myNewNum, boolean, myString)
}
