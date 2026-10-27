package main

import (
    "fmt"
    "sync"
    "time"
)

func receiver(messages <-chan int, wGroup *sync.WaitGroup, seconds float64) {
    for msg := range messages {                    
        time.Sleep(time.Duration(seconds) * time.Second)    
        // msg,  more:= <-messages               
        // if !more {
        //     fmt.Println(time.Now().Format("15:04:05"), "Channel closed, stopping receiver")
        //     break
        // }
        fmt.Println(time.Now().Format("15:04:05"), "Received:", msg)
    }
    fmt.Println(time.Now().Format("15:04:05"), "Channel closed, stopping receiver")
    wGroup.Done()                      
}

func sender(messages chan<- int, wGroup *sync.WaitGroup, seconds float64) {
    for i := 1; i <= 6; i++ {
        time.Sleep(time.Duration(seconds) * time.Second)
        messages <- i
        fmt.Println(time.Now().Format("15:04:05"), "Sending:", i)
    }
    // messages <- -1 
    close(messages) // Close the channel to signal the receiver that no more messages will be sent 
    wGroup.Done()
}

func main() {
    msgChannel := make(chan int, 3)              
    wGroup := sync.WaitGroup{}                   
    wGroup.Add(2)                                
    go receiver(msgChannel, &wGroup, 3.0)             
    go sender(msgChannel, &wGroup, 0.0)                                   
    wGroup.Wait()                                
}