package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func serial() {
	counter := 0

	increaseFn := func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}

	increaseFn()
	increaseFn()
	increaseFn()

	fmt.Println("serial: ", counter)
}

func parallel() {
	var counter = 0

	increaseFn := func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}

	go increaseFn()
	go increaseFn()
	go increaseFn()

	fmt.Scanln()

	fmt.Println("parallel: ", counter)
}

func fixed() {
	var counter = 0
	var lock sync.Mutex

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
	fmt.Println("fixed: ", counter)
	lock.Unlock()

}

func atomicop() {
	var counter int32 = 0

	done := make(chan bool)

	increaseFn := func() {
		for i := 0; i < 1000; i++ {
			atomic.AddInt32(&counter, 1)
		}
		done <- true // Ignore this for now, but its for waiting for a task to finish
	}

	go increaseFn()
	go increaseFn()
	go increaseFn()

	<-done
	<-done
	<-done

	fmt.Println("atomic: ", counter)
}

func main() {
	serial()
	parallel()
	fixed()
	atomicop()
}
