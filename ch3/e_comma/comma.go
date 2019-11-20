package main

import "fmt"

func main() {
	//should output "12,345,678"
	fmt.Println(comma("12345678"))
	//should output "123,456,789,012,345"
	fmt.Println(comma("123456789012345"))
}

//function should take in a string of numbers and insert a comma every
//spaces starting from the right
func comma(s string) string {
	n := len(s)
	//no comma needed if string is 3 numbers or less
	if n <= 3 {
		return s
	}
	//insert a comma three spaces from the right
	//send the remaining substring as an argument back into the function
	return comma(s[:n-3]) + "," + s[n-3:]
}
