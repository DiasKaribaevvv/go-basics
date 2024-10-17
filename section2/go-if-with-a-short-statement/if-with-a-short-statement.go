package main

import (
	"fmt"
)

/*

If with a short statement

Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

(Try using v in the last return statement.)
*/

func ageChecker(age, lim int) string {
	if v := age + 1; v < lim {
		return "You are too young !"
	}
	return "Thats okay"
}

func main() {
	fmt.Println(
		ageChecker(18, 20),
		ageChecker(15, 20),
	)

}
