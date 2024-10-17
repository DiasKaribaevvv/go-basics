package main

import "fmt"

/*
Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
*/

func deferLoop(lim int) {

	/*
	Stacking defers

	Deferred function calls are pushed onto a stack.
	When a function returns, its deferred calls are
	executed in last-in-first-out order.

	To learn more about defer statements
	read this blog post.
	*/

	fmt.Println("....")
	for i := 0; i < lim; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("finish")

}

func main() {
	defer fmt.Println("hello my friend")

	fmt.Println("hi")

	defer deferLoop(5)
}
