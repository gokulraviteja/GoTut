package main

import (
	"fmt"
)

func panicker() {
	fmt.Println("About to Panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error : ", err)
			// log.Println("Error : ", err)
			// log.Fatal("Error : ", err)
			// panic(err)   // use this when u want to stop main func
		}
	}()
	panic("Went wrong")
	fmt.Println("Will not be exe as after panic in the same func")
}

func main() {

	fmt.Println("Hello GO Panic\n")

	//Example case of panic
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// Practical Example
	// err = http.ListenAndServe(":8080", nil)
	// err  will be returned and execution continues , so will will panic the situatuion as it shld not be continued and to know the error
	// using
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println("start")
	// defer fmt.Println("middle")
	// panic("something wrong happened") // defer lines will be executed before panic
	// fmt.Println("end")

	// fmt.Println("start")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error : ", err)
	// 	}
	// }() // () -> invoking the func
	// panic("something wrong happened") // defer lines will be executed before panic
	// fmt.Println("end")

	fmt.Println("start")
	panicker()
	fmt.Println("end")

}
