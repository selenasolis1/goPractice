package main

import "fmt"

var a = []int{0, 1, 2, 3, 4, 5}

func main() {

	// fmt.Println("reverse: ", reverse(a[:]))
	fmt.Println("original: ", a)
	fmt.Println("left 2: ", rotateLeft(a[:], 2))
	fmt.Println("Right 2: ", rotateRight(a[:], 2))
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func rotateLeft(s []int, r int) []int {
	b := append(s[:0:0], s...)
	reverse(b[:2])
	reverse(b[2:])
	reverse(b)
	return b
}

func rotateRight(s []int, r int) []int {
	c := append(s[:0:0], s...)
	reverse(c)
	reverse(c[:r])
	reverse(c[r:])
	return c
}
