package struct_

import "fmt"

//这是一个自定义类型
type myInt int

type MyTest int

//不允许为别的包的自定义类型定义方法，
//只允许为当前包的类型定义方法
func (m myInt) hello() {
	fmt.Println(m)
	fmt.Println("我其实是一个 int ！")
}
func test1() {

	i := myInt(100)
	i.hello()
}
