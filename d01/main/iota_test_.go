package main

import (
	"fmt"
)

const (
	b = iota
	b1
	b2
)

func main() {
	fmt.Println(b)
	fmt.Println(b1)
	fmt.Println(b2)

	var aa = 0x66666
	var bb = 077
	fmt.Println(aa)
	fmt.Printf("type : %d ", aa)
	fmt.Printf("type : %T ", bb)

}
