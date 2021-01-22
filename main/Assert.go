package main

import "fmt"

func main() {

	var a interface{}
	a = 100

	switch t := a.(type) {
	case int8:
		fmt.Println("int8")
		fmt.Println(t)
	case int32:
		fmt.Println("human")
		fmt.Println(t)
	case int:
		fmt.Println("int")
		fmt.Println(t)
	}

}
