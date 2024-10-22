package main

import "fmt"

func recursion(num int) int {
	if num <= 0 {
		return num
	}

	return num + recursion(num-1)
}

func main() {
	fmt.Println(recursion(3))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(4))
}
