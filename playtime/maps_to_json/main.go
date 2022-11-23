package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// R&D on how to pass a map[int]string as json

	// attempting to pass a map with int keys is failing in crm backend
	// trying to see if there is a way to get around this

	map1 := map[int]string{
		1: "apple",
		2: "banana",
	}

	jsonStr, err := json.Marshal(map1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
}
