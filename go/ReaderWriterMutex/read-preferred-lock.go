package main

import "sync"

/* read-preferring lock:
a writer goroutine can only acquire lock only if all the readers have released thier locks
write-starvation
*/
type RPReadWriteMutex struct {
	readersCount int
	readerLock   sync.Mutex
	globalLock   sync.Mutex
}

func NewRPReadWriteMutex() *RPReadWriteMutex {
	return &RPReadWriteMutex{}
}

func (rwm *RPReadWriteMutex) ReadLock() {
	rwm.readerLock.Lock()
	rwm.readersCount++
	if rwm.readersCount == 1 {
		rwm.globalLock.Lock()
	}
	rwm.readerLock.Unlock()

}
func (rwm *RPReadWriteMutex) ReadUnlock() {
	rwm.readerLock.Lock()
	rwm.readersCount--
	if rwm.readersCount == 0 {
		rwm.globalLock.Unlock()
	}
	rwm.readerLock.Unlock()
}
func (rwm *RPReadWriteMutex) WriteLock() {
	rwm.globalLock.Lock()
}
func (rwm *RPReadWriteMutex) WriteUnlock() {
	rwm.globalLock.Unlock()
}
