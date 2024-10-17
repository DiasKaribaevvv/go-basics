package main

import (
	"fmt"
	"runtime"
	"time"
)

/*

Switch

A switch statement is a shorter way to
write a sequence of if - else statements.
It runs the first case whose value is equal
to the condition expression.

Go's switch is like the one in C, C++, Java,
JavaScript, and PHP, except that Go only
runs the selected case, not all the cases
that follow. In effect, the break statement
that is needed at the end of each case in
those languages is provided automatically
in Go. Another important difference is
that Go's switch cases need not be constants,
and the values involved need not be integers.
*/

func whenSunday() {
	var toDay time.Weekday = time.Now().Weekday()
	switch time.Thursday {
	case toDay + 0:
		fmt.Println("Today is sunday")
	case toDay + 1:
		fmt.Println("Tomorrow")
	default:
		fmt.Println("Work work work")
	}
}

func switchWithNoCondition(age int) string {

	/*
		Switch with no condition

		Switch without a condition is the same as switch true.

		This construct can be a clean way to write long if-then-else chains.
	*/

	switch {
	case age < 18:
		return "You can not watch this video"
	case age >= 18:
		return "Enjoy)"

	default:
		return "------"
	}

}

func main() {

	fmt.Println("Go runs on ... ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("LINUX")
	default:
		fmt.Println("--------")
	}

	whenSunday()
	fmt.Println(switchWithNoCondition(18))

}
