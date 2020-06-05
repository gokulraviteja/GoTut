package main

import (
	"fmt"
	"time"
)

func main() {

	msg := "hey gokul"

	go func() {
		fmt.Println(msg)
	}()

	//Race condition to detect when compiling use -race flag.
	// Avoid that use below practice

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }(msg)

	msg = "hey ravi"
	// before goroutine prints ,it is changing the variable to ravi -> race condition
	// even passing arg, will be  fine and prints gokul , using sleep is not a good practice , alternative is wait group

	time.Sleep(100 * time.Millisecond)

}
