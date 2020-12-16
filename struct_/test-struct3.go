package struct_

import "fmt"

//Go语言中如果标识符首字母大写，表示对外部可见！！！！
//type person struct  ---- 外部不可见  main.go无法调用
type Person3 struct {
	name string
	age  int
}

//方法 --和函数有区别
//func (p Person3) wang()  --- 无法被别的包调用
func (p *Person3) Wang() {
	fmt.Println(p.name, p.age)
}

//--- 无法被别的包调用
func (p Person3) wang() {
	fmt.Println("人不如狗！！！")
}
func test3() {
	p1 := newPersion("aaa", 11)
	p2 := newPersion2("aaa", 22)
	fmt.Println(p1)
	fmt.Println(*p2)

}

//构造函数（面向对象）
//当结构体比较复杂的时候，构造函数尽量返回结构体的指针，
//以节省系统资源的开销（返回值有做拷贝阶段）
//具体情况具体分析
func newPersion(name string, age int) Person3 {
	return Person3{
		name, age,
	}
}

func newPersion2(name string, age int) *Person3 {
	return &Person3{
		name, age,
	}
}

func NewPersion3(name string, age int) *Person3 {
	return &Person3{
		name, age,
	}
}
