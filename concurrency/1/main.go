package main

import "fmt"

func entrypoint() {
	fmt.Println("this happens at some unknown time")
}

func main() {
	fmt.Println("This happens first")

	// dispatch a new go routine (thread)
	go entrypoint() // go means run call entrypoint() on new goroutine

	fmt.Println("This can happen second or third")

	fmt.Scanln()
	fmt.Println("done")
}
