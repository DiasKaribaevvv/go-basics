package main

import (
	"fmt"
)

/*
Methods continued

You can declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.

You can only declare a method with a receiver whose type is defined in the same package as the method.
You cannot declare a method with a receiver whose type is defined in another package (which includes
the built-in types such as int).
*/

type MyFloat float64

func (my MyFloat) square() float64 {
	return float64(my) * float64(my)
}

func main() {
	fmt.Println(MyFloat(-2.0).square())
}
