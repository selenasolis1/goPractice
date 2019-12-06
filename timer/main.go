package main

import (
	"fmt"
	"time"
)

const two_seconds = 2
const five_seconds = 5
const ten_seconds = 10

func main() {

	// Run timer for 60 seconds. When timer expires,
	// call a function to print out that our timer is done.
	timer1 := NewTimer(two_seconds, func() {
		fmt.Println("It's been two seconds ", time.Now())
	})

	timer2 := NewTimer(five_seconds, func() {
		fmt.Println("It's been five seconds ", time.Now())
	})

	timer3 := NewTimer(ten_seconds, func() {
		fmt.Println("It's been ten seconds ", time.Now())
	})

	defer timer1.Stop()
	defer timer2.Stop()
	defer timer3.Stop()

	countdownBeforeExit := time.NewTimer(time.Second * 30)
	<-countdownBeforeExit.C

}

// NewTimer creates a timer that runs for a specified number of seconds.
// When the timer finishes, it calls the action function.
// Use the returned timer to stop the timer early, if needed.
func NewTimer(seconds int, action func()) *time.Timer {

	timer := time.NewTimer(time.Second * time.Duration(seconds))

	go func() {
		<-timer.C
		action()
	}()

	return timer
}
