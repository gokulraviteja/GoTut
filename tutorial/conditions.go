package main

import "fmt"

func main() {
	fmt.Println("Hello GO conditions")
	myMap := map[string]int{
		"gokul": 1,
		"teja":  2,
	}
	val, err := myMap["gokul"]
	fmt.Println(val, err)
	if err {
		fmt.Println(val)
	}

	if a, b := myMap["teja"]; b {
		fmt.Println(a)
	} else if b == false {
		fmt.Println("Impossible")
	} else {
		fmt.Println("IDK")
	}

	switch 2 {
	case 1, 3, 5:
		fmt.Println("one ,3,5 in switch")

	case 2, 4, 6:
		fmt.Println("two ,4,6in switch")

	default:
		fmt.Println("neither one nor two in switch")
	}

	var t interface{} = 1 // interface is a detatype which can take any type of data in go
	switch t.(type) {
	case int:
		fmt.Println("Integer")
	case float32:
		fmt.Println("Float")
	default:
		fmt.Println("Other type")
	}

}
