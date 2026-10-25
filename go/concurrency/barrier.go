package main

import "sync"

type Barrier struct {
	size      int
	waitCount int
	condVar   *sync.Cond
}

func NewBarrier(size int) *Barrier {
	return &Barrier{size, 0, sync.NewCond(&sync.Mutex{})}
}

func (b *Barrier) Wait() {
	b.condVar.L.Lock()
	b.waitCount++
	if b.waitCount == b.size {
		b.waitCount = 0
		b.condVar.Broadcast()
	} else{
		b.condVar.Wait()
	}
	b.condVar.L.Unlock()
}