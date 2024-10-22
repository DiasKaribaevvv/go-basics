package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	// enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			panic(err)
		}

		for k, _ := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		fmt.Println(v)

	}

}
