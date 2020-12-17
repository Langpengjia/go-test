package struct_

import "fmt"

/**
结构体实现继承(模拟)
*/

type human struct {
}

func (h human) thinking() {
	fmt.Println("正常人都会思考！！")
}

type father struct {
	name string
	age  int
	human
}

func (f father) thinking() {
	fmt.Println("父亲也会会思考！！")
}

func test6() {
	var f father
	f.thinking()
}
