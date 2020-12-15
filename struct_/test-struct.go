package struct_

import "fmt"

//声明一个结构体
type Person struct {
	name  string
	age   int
	sex   string
	hobby []string
}

func test() {
	//实例化一个类型----1
	var p Person
	p.name = "Peter"
	p.age = 20
	p.sex = "男"
	p.hobby = []string{"喝酒", "撸串", "打豆豆"}
	fmt.Println(p)
	fmt.Println(p.name)

	//实例化一个类型----2
	var peter = Person{"Peter", 18, "男", []string{"喝酒", "撸串", "打豆豆"}}
	fmt.Println(peter)

	//打印类型和值
	fmt.Printf("P 的type= %T ，值 = %v  ", p, p)

	//匿名结构体（多用于临时场景）
	//声明一个结构体变量s
	var s struct {
		name string
		age  int
	}
	s.name = "hh"
	s.age = 100
	//注意区分Person
	fmt.Printf("s 的type= %T ，值 = %v  ", s, s)
}
