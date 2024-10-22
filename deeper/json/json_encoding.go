package main

import (
	"encoding/json"
	"fmt"
)

type PersonalInfo struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint32 `json:"age"`
}

func main() {
	var person1 PersonalInfo = PersonalInfo{
		Name:    "Dias",
		Surname: "Karibaev",
		Age:     22,
	}

	b, err := json.Marshal(person1) // Marshall encode
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	var dataDecoded PersonalInfo

	err1 := json.Unmarshal(b, &dataDecoded) // decode
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(dataDecoded)

	var i interface{}

	byt := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	err2 := json.Unmarshal(byt, &i)
	if err2 != nil {
		panic(err2)
	}

	newInteface := i.(map[string]interface{})
	fmt.Println(newInteface["Name"])
	fmt.Println("----------------")
	for k, v := range newInteface {

		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
