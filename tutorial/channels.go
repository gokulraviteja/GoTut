package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Hello GO Channels")

	// ch := make(chan int)
	// wg.Add(2)
	// go func() {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// go func() {
	// 	ch <- 42
	// 	wg.Done()
	// }()

	// by default we are working with unbuffered channels so only one piece of data can be at a time present in the channel

	// wg.Add(1)
	// go func() {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()

	// }()
	// for ij := 0; ij < 5; ij++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		ch <- 41
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// Above commented code will give us error(deadlock),
	//as once 41 is pushed in channel it is retrieved and used(printed),
	//and more 41's are  being pushed into channel and no resource is using them which leads to a deadlock.

	// wg.Add(2)
	// go func() {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	ch <- 20
	// 	wg.Done()
	// }()
	// go func() {
	// 	ch <- 40
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// wg.Wait()

	ch := make(chan int, 50) // buffer 50 can store 50 integers inside channel
	wg.Add(2)
	go func(ch <-chan int) {
		//recieve only channel
		//receive data from channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		//send only
		//only send data to channel
		ch <- 10
		wg.Done()
	}(ch)

	// wg.Add(2)
	// go func(ch <-chan int) {
	// 	for i := range ch {
	// 		fmt.Println(i)
	// 	}
	// 	wg.Done()
	// }(ch)

	// go func(ch chan<- int) {
	// 	ch <- 10
	// 	ch <- 15
	// 	close(ch)
	// 	wg.Done()
	// }(ch)
	wg.Wait()
}
