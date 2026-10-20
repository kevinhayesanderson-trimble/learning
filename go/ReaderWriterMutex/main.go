package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	rwMutex := NewReadWriteMutex()
	var readers sync.WaitGroup
	readers.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			defer readers.Done()
			rwMutex.ReadLock()
			time.Sleep(time.Second)
			fmt.Println("Read done")
			rwMutex.ReadUnlock()
		}()
	}

	readers.Wait()
	rwMutex.WriteLock()
	fmt.Println("Write finished")
	rwMutex.WriteUnlock()
}
