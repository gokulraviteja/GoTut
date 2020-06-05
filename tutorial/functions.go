package main

import "fmt"

func main() {
	fmt.Println("Hello GO Functions\n")
	sayMessage("Its a Func")
	name := "gokul"
	greet := "hi"

	sayMessages(greet, name)
	fmt.Println(greet, name)

	sayMessagesp(&greet, &name)
	fmt.Println(greet, name)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("Returned value : ", s)

	fmt.Println(abyb(1.0, 2.0))

	fmt.Println(abyb(1.0, 0.0))

	func() {
		fmt.Println("I am an Anonymous Function!")
	}()

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println("anonymous ", i)
		}(i)
	}

	fun := func() {
		fmt.Println(" func assigned to variable")
	}
	fun()

	// dclaring fun in diff way

	var fun1 func(int, int) int

	fun1 = func(a, b int) int {
		return a + b
	}
	fmt.Println(fun1(100000, 10))

}

func sayMessage(msg string) {
	fmt.Println(msg)
}

// func sayMessages(greet string, name string) { //can simply written
func sayMessages(greet, name string) {
	// fmt.Println(greet, name)
}

func sayMessagesp(greet, name *string) {
	*name = "ravi"
	// fmt.Println(*greet, *name)
}

func sum(sliceName ...int) int { //... indicates get  all last  parameters and keep in slice
	fmt.Println(sliceName)
	s := 0
	for _, v := range sliceName {
		// fmt.Println(idx, v)
		s += v
	}
	fmt.Println("sum", s)
	return s
}

func abyb(a, b float32) (float32, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by 0")
	}
	return a / b, nil
}
