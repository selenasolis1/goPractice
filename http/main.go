package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}
	//has same result as code below
	//func Copy(dst Writer, src Reader) (written int64, err error)
	//
	io.Copy(lw, resp.Body)

	// make an empty byte slice with 99,999 spaces for elements
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// prints out HTML
	// fmt.Println(string(bs))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
