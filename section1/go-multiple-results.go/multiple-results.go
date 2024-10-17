package main

import (
	"fmt"
	"math"
)

//A function can return any number of results.

func anynum(x float64) (float64, float64, float64) {
	return math.Pow(x, 2), math.Pow(x, 3), math.Pow(x, 4)
}

func main() {
	fmt.Println("Enter any number I will show to the power of 2,3,4")
	var num float64
	fmt.Scan(&num)
	first, second, third := anynum(num)
	fmt.Printf("to the power of 2 -> %v to the power of 3 -> %v to the power of 4 -> %v ", first, second, third)

}
