package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

func main() {
	fmt.Println("Hello GO Methods\n")
	g := greeter{
		greeting: "hi",
		name:     "gokul",
	}
	// fmt.Println(g)

	greetFunction(g)
	g.greetMethod()

}

func (g greeter) greetMethod() {
	fmt.Println(g.greeting, g.name)

}

func greetFunction(g greeter) {
	fmt.Println(g.greeting, g.name)
}
