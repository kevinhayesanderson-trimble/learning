package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sys/cpu" // Install via: go get golang.org/x/sys/cpu
)

// BAD: Causes False Sharing. Both counters occupy the same 64-byte cache line.
type BadCounters struct {
	CounterA uint64 // 8 bytes
	CounterB uint64 // 8 bytes
}

// GOOD: Enforces Cache Isolation. 
type GoodCounters struct {
	CounterA uint64
	_        cpu.CacheLinePad // Injects 56 bytes of blank padding, pushing CounterB to a new line [1]
	CounterB uint64
	_        cpu.CacheLinePad // Prevents anything else following this struct from sharing the line [1]
}

func main1() {
	var wg sync.WaitGroup
	iterations := 500_000_000

	// Test Bad Counters (High Cache Bouncing)
	bad := BadCounters{}
	startBad := time.Now()
	
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range iterations { bad.CounterA++ }
	}()
	go func() {
		defer wg.Done()
		for range iterations { bad.CounterB++ }
	}()
	wg.Wait()
	fmt.Printf("Bad Counters (False Sharing): %v\n", time.Since(startBad))

	// Test Good Counters (Isolated Cache Lines)
	good := GoodCounters{}
	startGood := time.Now()
	
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range iterations { good.CounterA++ }
	}()
	go func() {
		defer wg.Done()
		for range iterations { good.CounterB++ }
	}()
	wg.Wait()
	fmt.Printf("Good Counters (Padded):        %v\n", time.Since(startGood))
}
