package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

// Config is the struct that represents your configuration.
type Config struct {
	ServerPort int    `json:"serverPort"`
	Database   string `json:"database"`
}

func not() {
	configFile := "config.json"

	// Initialize the initial configuration.
	config, err := loadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Start watching for file changes using fsnotify.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(configFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Watching for changes to", configFile)

	// Create a channel to receive fsnotify events.
	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Reloading configuration due to file change...")
					config, err = loadConfig(configFile)
					if err != nil {
						log.Println("Error reloading config:", err)
					} else {
						fmt.Println("Configuration reloaded:", config)
					}
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	<-done
}

// loadConfig loads the configuration from a JSON file.
func loadConfig(filename string) (Config, error) {
	var config Config

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
