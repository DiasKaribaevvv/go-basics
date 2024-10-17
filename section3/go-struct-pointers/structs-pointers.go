package main

import "fmt"

/*
Pointers to structs

Struct fields can be accessed through a
struct pointer.

To access the field X of a struct when we have the
struct pointer p we could write (*p).X.
However, that notation is cumbersome, so the
language permits us instead to write just p.X,
without the explicit dereference.
*/

type Person struct {
	Name    string
	Surname string
}

func main() {
	person1 := Person{"Dias", "Karibaev"}
	myPointer := &person1
	myPointer.Name = "Almaz"
	(*myPointer).Name = "Almaz" //same with before
	fmt.Println(person1)
}
