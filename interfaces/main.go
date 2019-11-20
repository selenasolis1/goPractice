package main

import "fmt"

//solution to declaring two functions of the same name
//to all types there is a new type "bot"
type bot interface {
	//if you are a type with a function called getGreeting and return a type string,
	//you are also of type bot and can use the getGreeting() function
	getGreeting() string
}

//both of type "bot" as well because of line 22 & 23
type englishBot struct{}
type spanishBot struct{}

func main() {
	//new value of type struct
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//do not need "eb" since we are not making use of it in function
func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

//no "sb"
func (spanishBot) getGreeting() string {
	return "Hola!"
}

//cannot declare functions of the same name
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
