package main

import (
	"fmt"
)

func checkType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("string type %v \n", v)
	case int:
		fmt.Printf("int type %v \n", v)
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	checkType("ST")
	checkType(20)
	checkType(5.5)
}
