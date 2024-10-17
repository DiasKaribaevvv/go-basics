package main

import "fmt"

type vertex struct {
	X, Y int
}

func main() {
	myVertex := vertex{1, 2}
	myVertex.X = 4
	fmt.Println(myVertex.X, myVertex.Y)

}
