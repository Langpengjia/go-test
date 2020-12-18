package interface_

import "fmt"

/**
接口是一种抽象类型  参照java

*/
func main() {
	var c cat
	var a Animal
	a = &c
	da(a)
}

type dog struct {
}

type cat struct {
}

type Animal interface {
	speak()
}

//da cat和dog若想调用此方法，必须实现Aninal中所有的方法！！
func da(x Animal) {
	x.speak()
}

func (c *cat) speak() {
	fmt.Println("猫叫")
}
func (c *dog) speak() {
	fmt.Println("狗叫")
}
