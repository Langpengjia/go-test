package struct_

import "fmt"

//person 结构体匿名字段 //不推荐

//结构体嵌套
type person struct {
	name    string
	age     int
	gender  bool
	address address
}

func main() {
	test5()
}

type address struct {
	province string
	city     string
}

//company 匿名嵌套结构体，也不推荐
type company struct {
	name    string
	address //匿名嵌套结构体
}

func test5() {
	p1 := person{
		name:    "111",
		age:     11,
		gender:  true,
		address: address{"aa", "bb"},
	}
	c1 := company{"111", address{"sss", "111"}}

	fmt.Println(p1.name, p1.address.city)
	//语法糖,只适用于匿名嵌套结构体
	fmt.Println(c1.city)
}
