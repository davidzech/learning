package main

import (
	"fmt"
	"time"
)

func main() {
	// go routines are great for running long running tasks

	fmt.Println("Starting long task")

	go func() { // you can also use "go" on an anonymous func locally
		for i := 0; i < 4; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}
		fmt.Print("task finished!")
	}() // remember to actually call the function we just defined!

	fmt.Scanln()
	fmt.Println("done")
}
