package main

import "fmt"

//声明一个结构体
type Person struct {
	name  string
	age   int
	sex   string
	hobby []string
}

func main() {
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
}
