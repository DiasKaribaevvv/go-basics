package main

import "fmt"

/*

Arrays

The type [n]T is an array of n values of type T.

The expression

var a [10]int

declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
*/

func main() {
	var a [2]string
	a[0] = "Dias"
	a[1] = "Karibaev"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)
}
