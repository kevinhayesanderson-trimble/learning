package main

import (
	"container/list"
	"sync"

	listing516 "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter5/listing5.16"
)

type NaiveChannel struct {
	buffer            *list.List
	mutex             sync.Mutex
	capacitySemaphore *listing516.Semaphore //blocks the sender when the buffer is full
	bufferSemaphore   *listing516.Semaphore //blocks the receiver when the receiver is empty
}

func NewNaiveChannel(capacity int) *NaiveChannel {
	return &NaiveChannel{
		buffer:            list.New(),
		capacitySemaphore: listing516.NewSemaphore(capacity),
		bufferSemaphore:   listing516.NewSemaphore(0),
		mutex:             sync.Mutex{},
	}
}

func (nc *NaiveChannel) Send(){
	if nc.capacitySemaphore.permits 
}
