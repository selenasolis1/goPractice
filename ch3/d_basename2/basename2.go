package main

import (
	"fmt"
	"strings"
)

func main() {
	//should output b.c
	fmt.Println(basename("a/b.c.go"))
	//should output 'c'
	fmt.Println(basename("a/b/c.go"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") //-1 if '/' not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
