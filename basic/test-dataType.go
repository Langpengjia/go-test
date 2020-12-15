package basic

import (
	"fmt"
	"math"
)

//基本数据类型

//声明常量
//常量是可以跨文件和包被使用的

func main() {
	//注意iota值的变化
	fmt.Println(const1)
	fmt.Println(const2)
	fmt.Println(const3)
	fmt.Println(const4)
	fmt.Println(const5)
	fmt.Println(const6)
	fmt.Println(const7)
	fmt.Println(const8)
	fmt.Println(const9)
	fmt.Println(const10)
	fmt.Println(const11)

	//基本数据类型
	/**
	一个简单的demo
	*/
	var a1, c1 int = 1, 2
	fmt.Println(a1)
	fmt.Println(c1)

	fmt.Println("Hello World!")
	/**
	符号引用,无视转义
	*/
	s := `
		第一行
		第二行
		第三行
		第四行
		/n/t/f/s`
	fmt.Println(s)
	/**
	转换不同的数据类型
	*/
	fmt.Println(" int8 ranges:", math.MinInt8, math.MaxInt8)
	fmt.Println(" int16 ranges:", math.MinInt16, math.MaxInt16)
	fmt.Println(" int32 ranges:", math.MinInt32, math.MaxInt32)
	fmt.Println(" int64 ranges:", math.MinInt64, math.MaxInt64)
	var a int32 = 1047483674
	fmt.Printf("int32: 0x%x %d\n", a, a)
	b := int16(a)
	fmt.Printf("int16: 0x%x %d\n", b, b)
	var c float32 = math.Pi
	fmt.Println(int(c))
}

const const1 = "这是一个常量"

//若常量未赋值，则跟随前一个常量的值
const (
	const2 = 1
	const3
	const4
)

//iota：在  \\同一个常量块中\\  初始化时为0
//常量块中每增加一行iota默认加一，也包括单独定义变量
const (
	const5 = iota
	const6
	const7
	const8 = 200
	const9
	const10 = iota
	const11
)
