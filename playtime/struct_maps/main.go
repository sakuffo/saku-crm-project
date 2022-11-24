package main

import "fmt"

type TEST struct {
	x int
	y []string
}

var m = map[string]*TEST{"a": {2, []string{"a", "b"}}}
var m2 = map[string]TEST{"a": {2, []string{"a", "b"}}}

func main() {
	a := new(TEST)
	a.x = 0
	a.y = []string{"c", "d"}

	a2 := new(TEST)
	a2.x = 0
	a2.y = []string{"x", "z"}

	m["toto"] = a
	m2["bobo"] = *a2

	fmt.Println(m)
	fmt.Println(m2)
}
