package main

import (
	"fmt"
)

type personalInfo interface {
	fullName() string
}

type Person struct {
	name    string
	surname string
}

func (p *Person) fullName() string {
	return p.name + " " + p.surname
}

func main() {
	var personalInfo personalInfo = &Person{}
	var person1 Person = Person{
		name:    "Dias",
		surname: "Karibaev",
	}
	personalInfo = &person1

	fmt.Println(personalInfo.fullName())
}
