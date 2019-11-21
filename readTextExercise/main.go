package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("file reading error", err)
		return
	}
	fmt.Println("Contents of file:")
	fmt.Println(string(data))
}
