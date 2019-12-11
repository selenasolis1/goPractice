package main

import (
	"fmt"

	geometry "../geometry"
)

func main() {
	q := geometry.Point{1, 2}
	p := geometry.Point{4, 6}
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{1, 4},
		{1, 1},
	}

	fmt.Println(geometry.Distance(p, q)) // "5" Function Call
	fmt.Println(p.Distance(q))           //"5" Method Call
	fmt.Println(perim.Distance())
}
