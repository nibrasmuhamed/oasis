package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// ... (rest of the struct definitions and main function)

func com() {
	s := 0
	f := 0
	// Load before.json and after.json
	beforeData, err := ioutil.ReadFile("after.json")
	if err != nil {
		log.Fatal("Error reading before.json:", err)
	}

	afterData, err := ioutil.ReadFile("x.json")
	if err != nil {
		log.Fatal("Error reading after.json:", err)
	}

	var beforeBalances []Balance
	var afterBalances []Balance

	// Unmarshal JSON data
	if err := json.Unmarshal(beforeData, &beforeBalances); err != nil {
		log.Fatal("Error unmarshaling before.json:", err)
	}

	if err := json.Unmarshal(afterData, &afterBalances); err != nil {
		log.Fatal("Error unmarshaling after.json:", err)
	}

	// Calculate differences
	for i, b := range beforeBalances {
		a := b.BalanceItemSpec[0].BalanceItem[0].ConsumedBalance
		c := afterBalances[i].BalanceItemSpec[0].BalanceItem[0].ConsumedBalance
		diff := c - a
		x := 8.0

		if diff == x {
			fmt.Printf("Customer %s consumedBalance difference matches %v: got %v\n", b.CustomerID, x, diff)
			f++
		} else {
			fmt.Printf("Customer %s consumedBalance difference does not match %v: got %v\n", b.CustomerID, x, diff)
			s++
		}
	}
	fmt.Println("success: ", f, "\n", "failed: ", s)
}
