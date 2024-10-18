package main

import (
	"fmt"
	"strconv"
)

type myInterface interface {
	allInfo() string
}

type Person struct {
	name    string
	surname string
}

func (p *Person) allInfo() string {
	return p.name + p.surname
}

type myNum int

func (m myNum) allInfo() string {
	return "Number is: " + strconv.Itoa(int(m))
}

type F float64

func (f F) allInfo() string {
	return "Number is: " + strconv.FormatFloat(float64(f), 'f', -1, 64)
}

func main() {
	var myInterface myInterface
	myInterface = &Person{
		name:    "Dias",
		surname: "Karibaev",
	}
	fmt.Println(myInterface.allInfo())

	myInterface = myNum(5)
	fmt.Println(myInterface.allInfo())

	myInterface = F(1.5)
	fmt.Println(myInterface.allInfo())

}
