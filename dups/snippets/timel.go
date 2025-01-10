package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Add(24 * time.Hour)
	fmt.Println(t)
	now := time.Now()
	fmt.Println(int64(t.Sub(now).Hours() / 24))
}
