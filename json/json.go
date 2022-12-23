package json

import (
	"encoding/json"
	"fmt"
)
func Json() {
	jsonString := `
	{ 
		"data": {
			"username" : "Jonik",
			"passowrd" : "1111111",
			"id" : 1
		}
	}
	`
	var m map[string]map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &m); err != nil {
		panic(err)
	}
	fmt.Println(m["data"]["username"])
	fmt.Println("*-----------------------------*")
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
