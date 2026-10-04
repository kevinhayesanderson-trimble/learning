package main

import (
	"fmt"
	"time"
)

func main2() {
	var balance int = 0

	// Goroutine 1: Modifying memory directly
	go func() {
		for range 100 {
			balance += 10 // Data Race: Write operation
		}
	}()

	// Goroutine 2: Reading memory directly
	go func() {
		for range 100 {
			fmt.Println("Current balance:", balance) // Data Race: Read operation
		}
	}()

	time.Sleep(100 * time.Millisecond)
}
