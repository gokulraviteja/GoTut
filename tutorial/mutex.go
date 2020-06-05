package main

import (
	"fmt"
	"sync"
)

func bestPractices() {
	//Dont create goroutines in library , let user of library decide when to use or not
	// knowing when to end a goroutine , when creating avoid mem leaks ...
	// checking for raceconditions at compile times as in routines.go  (useing -race flag)

}

var wg = sync.WaitGroup{} // sync multiple go routines
var counter = 0
var m = sync.RWMutex{}

// mutex -> only one can have the data lock
// rwmutex -> anyone can read the data , only one can write(cannot write when reading)

func main() {
	// by default go uses no of cores= no of threads
	// runtime.GOMAXPROCS(1) - makes the process single threaded
	// runtime.GOMAXPROCS(-1) - get no of threads prog is using

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayhello()
		m.Lock()
		go increment()

		// This way works fine but useless , as we not not using any concurrency, wose than single threaded as dealing with mutex
	}
	wg.Wait()
}

func sayhello() {
	// m.RLock()
	fmt.Println("Hello", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	// m.Lock()
	counter += 1
	m.Unlock()
	wg.Done()

}
