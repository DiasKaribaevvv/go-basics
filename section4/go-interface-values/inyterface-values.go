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
	return "Number is :" + strconv.Itoa(int(m))
}

type F float64

func (f F) allInfo() {
	fmt.Println(f)
}

func main() {
	var myInterface myInterface
	myInterface = &Person{
		name:    "Dias",
		surname: "Karibaev",
	}
	myInterface.allInfo()

	myInterface = myNum(5)
	myInterface.allInfo()

	myInterface = F(1.5)

}
