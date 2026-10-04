package main

import (
	"fmt"
	"runtime"
	//"time"
)

func sayHello() {
	//time.Sleep(1 * time.Microsecond)
	fmt.Println("Hello")
}

func sayBye() {
	//time.Sleep(1 * time.Microsecond)
	fmt.Println("Bye")
}

/*
# 1. Start all 10 processes concurrently using thread jobs
$jobs = 1..10 | ForEach-Object { Start-ThreadJob { .\main.exe } }

# 2. Wait for them to finish while streaming their output to the console in real-time
$jobs | Wait-Job | Receive-Job

the go scheduler will pick goroutines randomly, so the order of execution is not guaranteed.
so we'll see "Hello" and "Bye" printed in an unpredictable order,
and may not be printed at all if the main function exits before the goroutines get a chance to run.
*/
func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	go sayHello()
	go sayBye()
	runtime.Gosched()
	fmt.Println("Finished")

	main1()
	// detect race - go run -race .
	main2()
}
