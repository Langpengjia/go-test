package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type person1 struct {
	person
	Sku string `json:"sku"`
}

func (p person) personAge() {
	fmt.Println(p.Age)
}

func main() {
	per := `{"name":"aaa","age":11}`
	var p person1
	json.Unmarshal([]byte(per), &p)
	of := reflect.TypeOf(p)
	fmt.Println(of)
	p.personAge()
	ss := `asd`
	fmt.Println(reflect.TypeOf(ss), per)
}
