package main

import (
	"sync"
)

// read-preferring lock
type ReaderWriterMutex struct {
	readersCount int
	readerLock   sync.Mutex
	globalLock   sync.Mutex
}

func (rwm *ReaderWriterMutex) ReadLock() {
	rwm.readerLock.Lock()
	rwm.readersCount++
	if rwm.readersCount == 1 {
		rwm.globalLock.Lock()
	}
	rwm.readerLock.Unlock()

}
func (rwm *ReaderWriterMutex) ReadUnlock() {
	rwm.readerLock.Lock()
	rwm.readersCount--
	if rwm.readersCount == 0 {
		rwm.globalLock.Unlock()
	}
	rwm.readerLock.Unlock()
}
func (rwm *ReaderWriterMutex) WriteLock() {
	rwm.globalLock.Lock()
}
func (rwm *ReaderWriterMutex) WriteUnlock() {
	rwm.globalLock.Unlock()
}

func main() {
	rwm := &ReaderWriterMutex{}
	rwm.ReadLock()
	rwm.ReadUnlock()
	rwm.WriteLock()
	rwm.WriteUnlock()
}
