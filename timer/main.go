package main

import (
	"fmt"
	"sync"
	"time"
)

const two_seconds = 2
const five_seconds = 5
const ten_seconds = 10

// Wait group variable for timers
var wg sync.WaitGroup

func main() {

	// Adds one event to wait group
	wg.Add(1)
	//Creates new go routine
	go NewTimer(two_seconds)

	// Adds one event to wait group
	wg.Add(1)
	//Creates new go routine
	go NewTimer(five_seconds)

	// Adds one event to wait group
	wg.Add(1)
	//Creates new go routine
	go NewTimer(ten_seconds)

	// Wait to finish main go routine until all events are finished
	wg.Wait()
}

// NewTimer creates a new timer
func NewTimer(seconds int) {
	timer := time.NewTimer(time.Second * time.Duration(seconds))
	for {
		select {
		case <-timer.C:
			fmt.Printf("It's been %v seconds: %v\n", seconds, time.Now())
			// Complete event in wait group
			wg.Done()
		}
	}
}
