package main

import "fmt"

func main() {
	var students [3]string
	var array2d [3][3]int = [3][3]int{[3]int{1, 2, 3}, [3]int{1, 2, 3}, [3]int{1, 2, 3}}
	students[0] = "hello"

	fmt.Println(students, array2d)

	a := []int{1, 2, 3}            // its a slice
	arr := [...]int{1, 2, 3, 4, 5} // its an array

	fmt.Println(a, len(a), a[2:], arr)

	// b := a
	// b := &a

	sliceEx := make([]int, 3, 100) // 100 is the capacity here
	sliceEx = append(sliceEx, 1)

	fmt.Println(sliceEx)

	mapex := map[string]int{
		"abc1": 1,
		"abc2": 2,
		"abc3": 3,
	}
	fmt.Println(mapex)

	// map1 := map[[]int]int{}  // error as slice cannot be a key , but array can be a key
	// map1 := map[[2]int]int{}

	delete(mapex, "abc1")

	val, ok := mapex["abc1"]
	fmt.Println(val, " ", ok)

	fmt.Println("Hello GO Arrays")
}
