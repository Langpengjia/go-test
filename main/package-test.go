package main

import (
	"fmt"
	"go-test/package_"
)

func init() {
	fmt.Println("main.go")
}

func main() {
	package_.Test2()
	package_.Test1()

}
