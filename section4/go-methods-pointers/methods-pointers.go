package main

import (
	"fmt"
)

type Person struct {
	name    string
	surname string
}

// we use here pointer to do not pass copy of our Person
func (p *Person) changeName(newName string) string {
	p.name = newName
	return "New name now: " + newName
}

func main() {
	myPerson := Person{
		name:    "Dias",
		surname: "Karibaev",
	}
	fmt.Println(myPerson)
	//changing the name
	myPerson.changeName("Marat")

	fmt.Println(myPerson)
}
