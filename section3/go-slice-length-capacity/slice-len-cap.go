package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// here nil slice
	var slice []int
	fmt.Println(len(slice), cap(slice))
	if len(slice) == 0 && cap(slice) == 0 {
		fmt.Println("niiiiiiil!")
	}

	//here I create slice with using make function
	a := make([]int, 5)
	b := make([]int, 0, 5)
	printSlice(a)
	printSlice(b)

	// slices of slices
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	fmt.Println(board)
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
