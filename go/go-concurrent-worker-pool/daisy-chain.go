package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

func daisyChain() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for range n {
		right = make(chan int) 
		/*these channel variable are pointer to channel data in memeory, 
		so the on each new assignment , the pointer is pointing to new data
		the old channel data still exists in memory
		the old channel data in memory is not collected by GC since we
		are holding the reference to it in the chain
		*/
		go f(left, right)
		left = right // passing reference around the chain
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
 