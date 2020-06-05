package main

import "fmt"

type myStuct struct {
	number int
}

func main() {
	fmt.Println("Hello GO Pointers\n")

	var a int = 1
	var b int = a
	b = 2
	fmt.Println("No pointers : ", a, b)

	var ap int = 1
	var bp *int = &ap
	*bp = 2
	fmt.Println("Pointers : ", ap, *bp)

	arr := []int{1, 2, 3}
	a1 := &arr[0]
	a2 := &arr[1]
	fmt.Println(arr, *a1, *a2)

	var ms myStuct
	ms = myStuct{number: 111}
	fmt.Println(ms)

	var ms1 *myStuct
	ms1 = new(myStuct)
	fmt.Println(*ms1)

	(*ms1).number = 222
	fmt.Println((*ms1).number)

	//above and below both works fine

	ms1.number = 333
	fmt.Println(ms1.number)

	//slices deal with pointers when copied same with maps (they deal with references but not directly with underlying data)
	ars := []int{1, 2, 3}
	brs := ars
	ars[1] = 3
	fmt.Println(ars, brs)

	//arrays doesnot
	ar := [3]int{1, 2, 3}
	br := ar
	ar[1] = 3
	fmt.Println(ar, br)

}
