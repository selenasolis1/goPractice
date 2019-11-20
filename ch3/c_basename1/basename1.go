//basename removes directory components and a .suffix.
//e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"fmt"
)

func main() {
	//should output b.c
	fmt.Println(basename("a/b.c.go"))
	//should output 'c'
	fmt.Println(basename("a/b/c.go"))
}

func basename(s string) string {
	//Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			// s holds everything after the last '/'
			s = s[i+1:]
			break
		}
	}
	//Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			//s holds everything before the '.'
			s = s[:i]
			break
		}
	}
	//s should hold everything after the last '/' and before the last '.'
	return s
}
