package main

import "sync"

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func NewSemaphore(n int) *Semaphore{
	return &Semaphore{permits: n, cond: sync.NewCond(&sync.Mutex{})}
}

func (sem Semaphore) Acquire()  {
	sem.cond.L.Lock()
	for sem.permits <= 0{
		sem.cond.Wait()
	}
	sem.permits--
	sem.cond.L.Unlock()
}

func (sem Semaphore) Release()  {
	sem.cond.L.Lock()
	sem.permits++
	sem.cond.Signal()
	sem.cond.L.Unlock()
}