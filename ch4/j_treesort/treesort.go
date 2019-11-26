package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var a = []int{7, 2, 5, 4, 3, 6, 1}
	Sort(a)
	fmt.Println(a)
}

//sort sorts values in place
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		fmt.Println(t)
		//Equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		fmt.Println("value: ", value)
		fmt.Println("t.value: ", t.value)
		t.left = add(t.left, value)
		fmt.Println("t.left", t.left)
	} else {
		fmt.Println("value: ", value)
		fmt.Println("t.value: ", t.value)
		t.right = add(t.right, value)
		fmt.Println("t.right", t.right)
	}
	fmt.Println(t)
	return t
}
