package main

import "fmt"

func main() {
	defer func() { fmt.Println("不管出不出错，我都会这么做！") }()
	funcA()
	funcB()
	funcC()
}

func funcA() {
	panic("人为制造了一个错误！！！")
	fmt.Println("a")
}
func funcB() {
	fmt.Println("b")
}
func funcC() {
	fmt.Println("c")
}
