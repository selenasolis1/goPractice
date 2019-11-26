package main

import "fmt"

var a = [6]int{0, 1, 2, 3, 4, 5}

func main() {

	// fmt.Println("reverse: ", reverse(a[:]))
	fmt.Println("original: ", a)
	b := &a
	reverse(b)
	fmt.Println("reversed: ", b)
}

func reverse(s *[6]int) {
	for i := 0; i < len(s)/2; i++ {
		end := len(s) - i - 1
		s[i], s[end] = s[end], s[i]
	}
}
