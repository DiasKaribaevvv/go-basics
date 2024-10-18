package main

import (
	"fmt"
)

/*Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.
*/

type Person struct {
	name    string
	surname string
}

func (p Person) fullName() string {
	return p.name + " " + p.surname
}

func main() {
	fmt.Println(Person{
		name:    "Dias",
		surname: "Karibaev",
	}.fullName())

}
