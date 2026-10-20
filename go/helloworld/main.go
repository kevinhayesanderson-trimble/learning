package main

import (
	"fmt"
	"sync"
	"unsafe"
	"github.com/cutajarj/ConcurrentProgrammingWithGo/chapter5/listing5.16"

)

func bool2Int(b bool) int {
	return int(*(*byte)(unsafe.Pointer(&b)))
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(bool2Int(5 >= 4))
	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	for range 50000 {
		go doWork1(cond)
		fmt.Println("Waiting for child goroutine ")
		cond.Wait() // releases the lock, waits for signal and suspends the main go-routine
		// cond.Wait() in main() wakes up, re-acquires cond.L.Lock(), and continues to the next iteration
		fmt.Println("Child goroutine finished")
	}
	cond.L.Unlock()

	//Semaphore doesn't have missed signal problem
	semaphore := listing5_16.NewSemaphore(0)       
    for i := 0; i < 50000; i++ {                   
        go doWork2(semaphore)                       
        fmt.Println("Waiting for child goroutine ")
        semaphore.Acquire()                        
        fmt.Println("Child goroutine finished")
    }
}

func doWork(cond *sync.Cond) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	cond.Signal()
}

func doWork1(cond *sync.Cond) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	cond.L.Lock() // cannot acquire the lock until the wait() in main releases it
	cond.Signal()
	cond.L.Unlock()
}

func doWork2(semaphore *listing5_16.Semaphore) {
    fmt.Println("Work started")
    fmt.Println("Work finished")
    semaphore.Release()                            
}
