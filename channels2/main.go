package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		//waits for response before moving on to the next link
		go checkLink(link, c)
	}
	//blocking call so only 1 response will be printed
	//fmt.Println(<-c)
	//instead ...

	for l := range c {
		//do not want to put sleep function here because it
		//it will be inside the main routine and it will pause
		//the main routine. No other go routines can get started

		//function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
	// for l := range c {
	//		checkLink(l, c)
	//	}
	//same as...
	//for {
	//	  go checkLink(<-c, c)
	// }

}

func checkLink(link string, c chan string) {
	//get request
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//send link back into channel
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
