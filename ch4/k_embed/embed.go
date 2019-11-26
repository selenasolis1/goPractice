package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point  //Do not have to put Center Point
	Radius int
}
type Wheel struct {
	Circle // do not have to put Circle Circle
	Spokes int
}

func main() {
	w := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, //Note trailing comma necessary
	}
	fmt.Println(w)         //no keys
	fmt.Printf("%#v\n", w) //keys with values
	fmt.Println(w.X)       //8
	fmt.Println(w.Y)       //8
	fmt.Println(w.Point)   //{8,8}
}
