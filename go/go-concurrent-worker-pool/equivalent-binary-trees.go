package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)


func Walk(t *tree.Tree, ch chan int){
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go func(){
		Walk(t1, c1)
		close(c1)
	}()
	go func(){
		Walk(t2, c2)
		close(c2)
	}()
	for {
		v1, ok1 := <- c1
		v2, ok2 := <- c2
		if !ok1 && !ok2 {
			break
		}
		fmt.Printf("Comparing: Tree1 = %v (open: %t) | Tree2 = %v (open: %t)\n", v1, ok1, v2, ok2)

		if ok1 != ok2 || v1 != v2 {
			fmt.Println("Mismatch found! Exiting early.")
			return false 
		}
	}
	 return true
}

func test() {
	t1 := tree.New(2)
	t2 := tree.New(2)
	print(Same(t1, t2))
}