package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma2("1234567")) // "[1, 2, 3]"
}

//like fmt.Sprint(values) but adds commas
func comma2(s string) string {
	//initialize buffer
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		//rem holds how many digits until the first comma
		rem := len(s) % 3
		//current character
		char := s[i : i+1]
		// prints writes character to buffer
		fmt.Fprintf(&buf, "%v", char)
		//writes a comma to the buffer every 3 characters
		//last comma should be placed 3 from the end
		if i%3 == rem-1 && i < len(s)-2 {
			buf.WriteString(",")
		}
	}
	//returns the buffer as a string
	return buf.String()
}
