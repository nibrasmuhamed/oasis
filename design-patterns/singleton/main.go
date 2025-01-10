package main

import (
	"fmt"
	"singleton/single"
)

func main() {
	// Get the singleton instance
	db1 := single.GetInstance()
	db2 := single.GetInstance()

	// Both db1 and db2 should point to the same instance
	if db1 == db2 {
		fmt.Println("Both db1 and db2 are the same instance")
	} else {
		fmt.Println("db1 and db2 are different instances")
	}
}
