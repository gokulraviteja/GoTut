package main

import (
	"encoding/json"
	"fmt"
)

type objec struct {
	Title string
	Desc  string
}
type objecjson struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func main() {
	fmt.Println("Hello GO Json")
	obb := objec{Title: "hey", Desc: "Whatsup"}
	fmt.Println("Plain : ", obb)

	obb1 := objecjson{Title: "hey", Desc: "Whatsup"}
	byt, err := json.Marshal(obb1)
	if err != nil {
		panic(err)
	}

	out := string(byt)
	fmt.Println("Json : ", out)

	var getback objec
	err1 := json.Unmarshal(byt, &getback)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("Json to struct : ", getback)

}
