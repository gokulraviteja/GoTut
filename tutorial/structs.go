package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number    int
	actorName string
	patients  []string
}

type Obj struct {
	name    string `required max="100"` // two conditions one is required and other is max length as 100
	countrt string
}

type SubObj struct {
	Obj
	speed float32
}

func main() {

	fmt.Println("Hello GO structs")

	doctor := Doctor{
		number:    3,
		actorName: "gokul",
		patients: []string{
			"ravi",
			"teja",
		},
	}

	fmt.Println(doctor, doctor.actorName)

	//anonymous struct

	aDoctor := struct{ name string }{name: "gokul"}
	fmt.Println(aDoctor)

	subobj := SubObj{}
	subobj.name = "emu"
	subobj.countrt = "australia"
	subobj.speed = 12.1
	fmt.Println(subobj)

	t := reflect.TypeOf(SubObj{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)

}
