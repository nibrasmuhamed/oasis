package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Persons struct {
	sync.Mutex
	Person []Person `json:"person"`
}
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func hello() {
	var person Persons
	// Read initial data from file
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Unmarshal JSON into struct
	err = json.Unmarshal(data, &person.Person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Watch file for changes
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add("data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Start goroutine to handle file events
	go person.updateData(watcher)

	// Print initial data
	fmt.Println("Initial data:", person)

	// Keep the main goroutine running
	for {
		time.Sleep(time.Second)
	}
}

func (p *Persons) updateData(watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				// File was modified, read new data and update struct
				p.Lock()

				data, err := ioutil.ReadFile("data.json")
				if err != nil {
					fmt.Println("Error:", err)
					p.Unlock()
					continue
				}

				err = json.Unmarshal(data, &p.Person)
				if err != nil {
					fmt.Println("Error:", err)
					p.Unlock()
					continue
				}

				p.Unlock()
				fmt.Println("Updated data:", p.Person)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error:", err)
		}
	}
}
