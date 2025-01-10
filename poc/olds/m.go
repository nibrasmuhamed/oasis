package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Customer struct {
	CustomerID     string `json:"customerid"`
	SubscriptionID string `json:"subscriptionid"`
}

func soo() {
	// Read customer data from the "customers" file
	customerData, err := readLines("data/customers.txt")
	if err != nil {
		fmt.Println("Error reading customer data:", err)
		return
	}

	// Read subscription data from the "subscription" file
	subscriptionData, err := readLines("data/subscriptions.txt")
	if err != nil {
		fmt.Println("Error reading subscription data:", err)
		return
	}

	// Check if the number of customer IDs and subscription IDs match
	if len(customerData) != len(subscriptionData) {
		fmt.Println("Number of customer IDs does not match the number of subscription IDs.")
		return
	}

	var customers []Customer

	// Create a JSON object for each pair of customer and subscription IDs
	for i := 0; i < len(customerData); i++ {
		customer := Customer{
			CustomerID:     customerData[i],
			SubscriptionID: subscriptionData[i],
		}
		customers = append(customers, customer)
	}

	// Encode the customers into JSON format
	jsonData, err := json.Marshal(customers)
	if err != nil {
		fmt.Println("Error encoding data into JSON:", err)
		return
	}

	// Write the JSON data to the "out.json" file
	err = ioutil.WriteFile("out.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON data written to out.json")
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
