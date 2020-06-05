package main

import (
	"fmt"
	"strconv"
)

// declaring at package level - cannot declare as i:=1
// if name is given in caps it can be used by other packages else no , can be used in the same package

var ij int = 23

// declaring block of variables

var (
	name  string = "gokul"
	age   int    = 21
	idiot bool   = true
)

const (
	a1 = iota
	a2
	a3
)

const (
	b1 = iota
	b2
	b3
)

const (
	cons1 string = "constant1"
	cons2 string = "constant2"
)

func main() {

	// declaration
	var i int
	fmt.Println(i)

	// initialization
	i = 2
	i = 3
	fmt.Println(i)

	// both
	var j int = 22
	fmt.Println(j)

	//simple
	k := 42 // k should be a new variable
	fmt.Println(k)

	fmt.Printf("%v , %T\n", k, k)
	fmt.Println("Hello Go Variables")
	fmt.Println(name, " ", idiot)

	//Shadowing
	// ij is declared in package level
	// redeclaring is fine inside function

	var ij int = 100

	var ifl float32

	ifl = float32(ij)
	fmt.Println("Shaddowing : ", ij)
	fmt.Println("Typecasted : ", ifl)

	// by default strings castinf from numbers is not possible

	var st string = string(ij)
	fmt.Println("Type casted ", ij, " with string : ", st)
	fmt.Println("Type casted ", ij, " with strconv : ", strconv.Itoa(ij))

	s := "this is a string"
	b := []byte(s)
	fmt.Println(b)

	//constants

	const myCons int = 2
	// const myCons int = math.Sin(2)  --> error constants cannot be initialized at runtime

	fmt.Println("IOTA specifically in a constant block is something kindof enumerator , ", a1, a2, a3, b1, b2, b3, "scoped to that block only")

	fmt.Println("Constants : ", cons1, cons2)

}
