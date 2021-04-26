package main

import "fmt"

//切片扩容队底层数组的影响
func main() {

	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)

	s1 := arr[1:3]
	s2 := arr[3:6]

	fmt.Println(s1)
	fmt.Println(s2)

	//s1 = append(s1, 0,0,0,0)

	s1[1] = 0
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(arr)

}
