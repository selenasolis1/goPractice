package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "might be down. I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep, its up"
}
