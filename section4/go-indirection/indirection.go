package main

import (
	"fmt"
	"strconv"
)

/*
Methods and pointer indirection

Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK

For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
*/

type Person struct {
	name string
	age  int
}

func (p *Person) fullInfo() string {
	return p.name + strconv.Itoa(p.age)
}

// method which using pointer

func (p *Person) changeInfo(newName string, newAge int) {
	p.name = newName
	p.age = newAge
}

// function which takes as an argument pointer

func changeName(p *Person, newName string) {
	p.name = newName
}

func main() {
	person1 := Person{
		name: "Dias",
		age:  22,
	}
	person1.changeInfo("marat", 21)
	fmt.Println(person1)
	changeName(&person1, "Dias")
	fmt.Println(person1)

	// or another way

	person2 := &Person{
		name: "Nariman",
		age:  27,
	}
	changeName(person2, "Dias")
	person2.changeInfo("Dias", 21)
}
