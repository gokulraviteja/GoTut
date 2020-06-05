package main

import "fmt"

// unlike structs interfaces dont describe data , they describe behavior
// type exampleInterface interface {
// 	fun([]byte) (int, error)
// }

type Writer interface {
	Write([]byte) (int, error)
}

type consoleWriter struct{}

func (cw consoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementor interface {
	Increment() int
}

type intgr int

func (count *intgr) Increment() {
	*count++
}

func main() {

	cw := consoleWriter{}
	cw.Write([]byte("Hello interface again!"))

	i := intgr(10000)

	i.Increment()
	fmt.Println("Incremented : ", i)

	fmt.Println(i)

	fmt.Printf("Hello GO Interface %v ,%T\n", cw, cw)

	//Tutorial Incomplete
}
