package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // sync multiple go routines
var counter = 0

func main() {

	msg := "hey gokul"

	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "hey ravi"
	wg.Wait()

	// ok , wg is letting the code run before terminating the main thread(routine). But unable to sync  if multiple go routines
	// are accessing the same data in parallel
	// we will use mutex for this

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayhello()
		go increment()
	}
	wg.Wait()
}

func sayhello() {
	fmt.Println("Hello", counter)
	wg.Done()
}

func increment() {
	counter += 1
	wg.Done()
}
