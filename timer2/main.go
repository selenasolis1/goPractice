package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

const two_seconds = 2
const five_seconds = 5
const ten_seconds = 10

type Events struct {
	Events []Event
}

type Event struct {
	Component string `json:"component"`
	State     int    `json:"state"`
	Time      int    `json:"time"`
}

var events Events

// Wait group variable for timers
var wg sync.WaitGroup

func main() {
	jsonFile, err := os.Open("./data.json")
	if err != nil {
		fmt.Println("Error reading json file: ", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &events.Events)

	fmt.Println(events.Events)
	fmt.Println("Simulation Started: ", time.Now())
	for _, e := range events.Events {
		wg.Add(1)
		go NewTimer(e)
	}
	wg.Wait()
	fmt.Println("Simulation Ended: ", time.Now())
}

// NewTimer creates a new timer
func NewTimer(e Event) {
	timer := time.NewTimer(time.Second * time.Duration(e.Time))
	for {
		select {
		case <-timer.C:
			if e.Component[1:3] == "MV" {
				if e.State == 1 {
					fmt.Printf("%v ON %02d:%02d:%02d\n", e.Component, time.Now().Hour(), time.Now().Minute(), time.Now().Second())
				} else {
					fmt.Printf("%v OFF %02d:%02d:%02d\n", e.Component, time.Now().Hour(), time.Now().Minute(), time.Now().Second())
				}
			} else if e.Component[1:3] == "RP" {
				if e.State != 0 {
					fmt.Printf("%v at %v%% %02d:%02d:%02d\n", e.Component, e.State, time.Now().Hour(), time.Now().Minute(), time.Now().Second())
				} else {
					fmt.Printf("%v at %v%% %02d:%02d:%02d\n", e.Component, e.State, time.Now().Hour(), time.Now().Minute(), time.Now().Second())
				}
			} else if e.Component[1:3] == "RV" {
				if e.State == 1 {
					fmt.Printf("%v ON %02d:%02d:%02d\n", e.Component, time.Now().Hour(), time.Now().Minute(), time.Now().Second())
				} else {
					fmt.Printf("%v OFF %02d:%02d:%02d\n", e.Component, time.Now().Hour(), time.Now().Minute(), time.Now().Second())
				}
			}
			// Complete event in wait group
			wg.Done()
			return
		}
	}
}
