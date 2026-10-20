package main

import "sync"

//write-preferring
// when writers are waiting, new readers are blocked
type WPReadWriteMutex struct {
	readersCounter        int
	writerWaitingCounter  int
	writerActiveIndicator bool
	condVar               *sync.Cond
}

func NewReadWriteMutex() *WPReadWriteMutex {
	return &WPReadWriteMutex{condVar: sync.NewCond(&sync.Mutex{})}
}

func (rw *WPReadWriteMutex) ReadLock() {
	rw.condVar.L.Lock()
	for rw.writerActiveIndicator || rw.writerWaitingCounter > 0 {
		rw.condVar.Wait()
	}
	rw.readersCounter++
	rw.condVar.L.Unlock()
}

func (rw *WPReadWriteMutex) WriteLock() {
	rw.condVar.L.Lock()
	rw.writerWaitingCounter++
	for rw.writerActiveIndicator || rw.readersCounter > 0 {
		rw.condVar.Wait()
	}
	rw.writerWaitingCounter--
	rw.writerActiveIndicator = true
	rw.condVar.L.Unlock()
}

func (rw *WPReadWriteMutex) ReadUnlock() {
	rw.condVar.L.Lock()
	rw.readersCounter--
	if rw.readersCounter == 0 {
		rw.condVar.Broadcast()
	}
	rw.condVar.L.Unlock()
}

func (rw *WPReadWriteMutex) WriteUnlock() {
	rw.condVar.L.Lock()
	rw.writerActiveIndicator = false
	rw.condVar.Broadcast()
	rw.condVar.L.Unlock()
}
