package main

import (
	"encoding/json"
	"fmt"
)

//结构体与JSON 转换

type jsontest struct {
	Name string `json:"name"`
	I    int    `json:"i"`
}

func main() {
	j := jsontest{
		Name: "asdasdasd",
		I:    123,
	}
	b, err := json.Marshal(j)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(err)
	}
	var j2 jsontest
	err2 := json.Unmarshal(b, &j2)
	if err2 == nil {
		fmt.Println(j2.Name)
	}

}
