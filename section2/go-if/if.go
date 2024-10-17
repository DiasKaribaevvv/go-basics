package main

import (
	"fmt"
)

/*
If

Go's if statements are like its for loops;
the expression need not be surrounded
by parentheses ( ) but the braces { }are required.
*/

func main() {
	fmt.Println("Enter your age to check: ")

	var age int
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("You are to young !!!")
	} else {
		fmt.Println("Okay your age confirmed")
	}
}
