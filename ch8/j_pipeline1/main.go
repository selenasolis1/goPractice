package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Sqaurer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer
	for {
		fmt.Println(<-squares)
	}
}
