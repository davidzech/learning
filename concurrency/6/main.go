package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func divideAndConquer() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c) // sum the first half
	go sum(s[len(s)/2:], c) // sum the second half
	x := <-c                // block and wait until we have a value to receive from c
	y := <-c                // block and wait again until we have yet another value to receive from c

	fmt.Println(x, y, x+y)
}

func queue() {

	queue := make(chan int, 3) // make a channel with size 3

	queue <- 1 // <- operator on the right of the stack "pushes" 1 to the queue
	queue <- 2
	queue <- 3

	val := <-queue // <- operator on the left of the stack "pops" a value off from the front
	fmt.Println(val)

	val = <-queue
	fmt.Println(val)

	val = <-queue
	fmt.Println(val)

}

func bufferedWait() {
	queue := make(chan int, 3) // make a channel with size 3

	queue <- 1
	queue <- 2
	queue <- 3

	fmt.Println(<-queue)
	fmt.Println(<-queue)
	fmt.Println(<-queue)

	go func() {
		time.Sleep(1 * time.Second)
		queue <- 4
	}()

	fmt.Println(<-queue)
}

func deadLock() {
	ch := make(chan int, 2)
	done := make(chan bool)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("this is never reached")
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("this is never printed. goroutine above deadlocked")
	default:
		fmt.Println("deadlock, exiting...")
	}
}

func main() {
	divideAndConquer()
	queue()
	bufferedWait()
	deadLock()
}
