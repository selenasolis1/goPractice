package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f1, err := os.Open("ab.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer f1.Close()
	f2, err := os.Open("ab.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer f2.Close()
	buf1, _ := ioutil.ReadAll(f1)
	buf2, _ := ioutil.ReadAll(f2)

	c1 := sha256.Sum256(buf1)
	c2 := sha256.Sum256(buf2)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
