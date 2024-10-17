package main

import "fmt"

/*

Maps

A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.
*/

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex = map[string]Vertex{
	"Dias": Vertex{
		0.5,
		0.5,
	},
}

var myMap map[int]string = map[int]string{
	1: "Dias",
	2: "marat",
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Dias"])
	fmt.Println(myMap[1])

	delete(myMap, 1)

	fmt.Println(myMap)

	value, ok := myMap[1]
	fmt.Println(value, ok)
}
