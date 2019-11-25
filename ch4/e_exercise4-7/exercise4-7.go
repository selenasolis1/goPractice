package main

import "fmt"

func main() {
	str := "hello"
	b := []byte(str)
	fmt.Println("Original String: ", str)
	fmt.Println("Original Bytes: ", b)
	fmt.Println("Reversed Bytes: ", reverse(b))
	fmt.Println("Reversed String: ", string(b))
}

func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
