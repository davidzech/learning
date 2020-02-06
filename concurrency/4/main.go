package main

import (
	"fmt"
	"sync/atomic"
)

type Spinlock struct {
	acquired int32
}

func (s *Spinlock) Lock() {
	// try to set s.acquired to 1 if s.acquired == 0
	// CompareAndSwap safely checks if s.acquired == 0
	// if it couldn't be set, (the function returned false) try again infinitely
	for atomic.CompareAndSwapInt32(&s.acquired, 0, 1) == false {
	}
}

func (s *Spinlock) Unlock() {
	// safely set s.acquired to false
	atomic.StoreInt32(&s.acquired, 0)
}

func main() {
	var counter = 0
	var lock Spinlock

	done := make(chan bool)

	increaseFn := func() {
		for i := 0; i < 1000; i++ {
			// wrap critical sections with acquiring and releasing a lock

			lock.Lock()
			counter++ // our critical section!
			lock.Unlock()
		}
		done <- true // Ignore this for now, but its for waiting for a task to finish
	}

	go increaseFn()
	go increaseFn()
	go increaseFn()

	<-done
	<-done
	<-done

	lock.Lock()
	fmt.Println("spinlock: ", counter)
	lock.Unlock()
}
