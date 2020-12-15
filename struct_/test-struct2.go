package struct_

import (
	"fmt"
)

//声明一个结构体
type Person1 struct {
	name   string
	age    int
	gender string
}

type student struct {
	name string
	age  int
}

type tes struct {
	a int8
	b int8
	c string
	d int8
}

func test2() {

	var p1 Person1
	p1.name = "李雷"
	p1.gender = "男"
	//p1.gender 未改变
	//函数中是值传递
	//f中p1更改的是副本
	f(p1)
	//fmt.Println(p1)

	//想改变值的话需要传递指针
	ff(&p1)
	//fmt.Println(p1)

	////两种初始化的区别
	//p3 := Person1{
	//	name:"ss",
	//	age: 11,
	//	gender: "男",
	//}
	//p4 := &Person1{
	//	name:"ss",
	//	age: 11,
	//	gender: "男",
	//}
	//fmt.Printf("p3 :  %T  %#v\n",p3,p3)
	//fmt.Printf("p4 :  %T  %#v\n",p4,p4)
	//fmt.Println(p4)

	//p5 := new(Person1)
	//p5.name = "asd"
	//fmt.Printf("%p\n",p5)
	//
	//var a int
	//a = 100
	//b:=&a
	//fmt.Printf("%t \n",a)
	//fmt.Printf("%v \n",b)
	//fmt.Println(b)

	////结构体占用一块连续的内存
	//cc := tes{100,100,"sada",100}
	//fmt.Printf("n.a : %p\n",&cc.a)
	//fmt.Printf("n.b : %p\n",&cc.b)
	//fmt.Printf("n.b : %p\n",&cc.c)
	//fmt.Printf("n.d : %p\n",&cc.d)

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		fmt.Println(&stu)
		m[stu.name] = &stu
		fmt.Println(m[stu.name])

	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}

func f(s Person1) {
	s.gender = "女"
}

func ff(s *Person1) {
	(*s).gender = "女"
	//go语言语法糖,自动根据指针找对应变量
	s.gender = "女"
}
