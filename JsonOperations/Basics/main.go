package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
{
	"data": {
		"object":"card",
		"id":"4543600307892860",
		"first_name":"Hakan",
		"last_name":"Altındiş",
		"balance": 500000
	}
}
`

	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}

	fmt.Println(m)

	fmt.Println("--------")

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}
