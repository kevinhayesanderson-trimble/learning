package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, dataChan chan string,  arg string) {
	if arg != "" {
		select{
		case <-ctx.Done():
			return
		case dataChan <- arg:
			
		}
	}
}

func contextMain() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	dataChan := make(chan string)
	worker(ctx, dataChan, "test string")
	for strValue := range dataChan {
		fmt.Printf("returnChan: %v\n", strValue)
	}
}
