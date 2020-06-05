package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Hello GO ControlFlow\n")

	//defer , panic , recover

	fmt.Println("Normal Exe")
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end\n")

	fmt.Println("Defer ")
	fmt.Println("start")
	defer fmt.Println("middle") // exe this after the main func exe but before the main returns (follow LIFO last deferred first called)
	fmt.Println("end\n")

	//Example of defer
	res, err := http.Get("http://www.google.com/robots.txt") //opening a connection
	defer res.Body.Close()                                   // closing the resource after the program
	if err != nil {
		// fmt.Println(err)
		log.Fatal("error in getting response ", err) // eq to printf and do os.exit(1)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// fmt.Println(err)
		log.Fatal("error in reading response ", err)
	}
	fmt.Println(string(data))

	aa := "start"
	defer fmt.Println(aa) // ans will be start as when this line exe it takes the argument at that point of time even it is changed later
	aa = "end"

}
