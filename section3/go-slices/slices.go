package main

import "fmt"

/*
Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

a[low : high]

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

a[1:4]

*/

func main() {
	primes := [6]int{2, 3, 5, 9, 6, 56}
	var slice []int = primes[1:4]
	var slice2 []int = []int{1, 2, 4, 5, 6}
	fmt.Println(slice, slice2)
}
