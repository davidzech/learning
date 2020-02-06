package main

import (
	"fmt"
	"time"
)

func taskEntry() {
	time.Sleep(3 * time.Second)
	fmt.Println("Task finished!")
}

func main() {
	fmt.Println("starting long running task")
	task := StartThread(taskEntry) // start a new thread that begins executing taskEntry on another thread
	fmt.Println("waiting for task to finish...")
	task.Join() // wait for this other task to finish running
	fmt.Println("done")
}
