package main

import "fmt"

const a = 3333
const (
	a1 = 100
	a2
	a3
	a4
)
const (
	s1 = "aaa"
	s2
	s3
)

func main() {
	//	//todo go语言是静态语言
	//	ss := [7]int{1, 2, 3, 4, 5, 6, 7}
	//	for i, s := range ss {
	//		if i == 2 {
	//			s = 252
	//			goto SF
	//		}
	//		s = s
	//
	//	}
	//	fmt.Println(ss)
	//SF:
	//	fmt.Println("早就饿瑟")

	//数组元素的类型和数量是数组类型的一部分
	//数组是值类型
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 5}

	//a11 := a1
	//a11[0] = 100
	//fmt.Println(a1, a11)
	fmt.Println(a1 == a2)

	//fmt.Printf("a1 type is %T \n", a1)
	//fmt.Printf("a2 type is %T \n", a2)

	//切片是引用类型
	//s1 := []int{1, 2, 3}
	//s2 := []int{1, 2, 3, 4, 5}

	//s11 := s1
	//s11[0] = 100
	//fmt.Println(s11,s1)

	//fmt.Printf("s1 type is %T \n", s1)
	//fmt.Printf("s2 type is %T \n", s2)

	////range 是值传递·
	//for _, v := range s1 {
	//	fmt.Println(v)
	//}

}
