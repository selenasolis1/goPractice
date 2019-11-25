package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//func Open(name string) (*File, error)
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
