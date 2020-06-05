package main

import "fmt"

func main() {
	fmt.Println("Hello GO Variables")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i <= 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	ij := 0
	for {
		ij += 1
		fmt.Println(ij)
		if ij == 5 {
			break
		}
	}

	s := []int{11, 22, 33}
	fmt.Println(s)

	for k, v := range s {
		fmt.Println(k, v)
	}

	st := "abcd"

	for k, v := range st {
		fmt.Println(k, v, string(v))
	}

}
