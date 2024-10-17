package main

import "fmt"

/*

Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

    0 for numeric types,
    false for the boolean type, and
    "" (the empty string) for strings.

*/

func main() {
	var myNum int
	var myName string
	var myBool bool
	var myFloat float64

	fmt.Printf("%v %q %v %v \n", myNum, myName, myBool, myFloat)
}
