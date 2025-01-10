package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSessionCounter(t *testing.T) {
	counter := SessionCounter{}

	// Keep track of the previous unique value
	var prevUniqueValue int

	for j := 1; j <= 100000; j++ {
		uniqueValue := counter.GetUniqueValue("firstworld")

		// Check if the current unique value is greater than the previous one
		if j > 1 && uniqueValue <= prevUniqueValue {
			t.Errorf("Generated unique value is not strictly increasing. Previous: %d, Current: %d", prevUniqueValue, uniqueValue)
		}

		// Print the session ID and unique value for verification
		fmt.Printf("Session ID: %s, Unique Value: %d\n", "firstworld", uniqueValue)

		// Set the previous unique value for the next iteration
		prevUniqueValue = uniqueValue

		// Simulate some processing time
		time.Sleep(time.Millisecond * 100)
	}
}
