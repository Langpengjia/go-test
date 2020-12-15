package function

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"time"
	unicode "unicode"
)

//函数
//函数的作用域
//函数中变量先在内部查找（局部变量）
//内部找不到就去外层查找，一直找到全局为止
func main() {
	//函数定义
	fmt.Println(resolveTime(25221511))
	_, hour, minute := resolveTime(10036)
	fmt.Println(hour, minute)

	//九九乘法表
	chengfa99biao()

	//匿名函数（匿名函数可以在声明后调用）
	d := func(data int) int {
		fmt.Println("hello", data)
		return data * 10
	}(100)
	fmt.Println(d)
	f := func(a string) {
		fmt.Println("this is the letter", a)
	}
	f("s")

	//使用匿名函数实现操作封装
	generator := playerGen("high noon!")
	name, hp := generator()
	fmt.Println(name, hp)

	//可变参数--参数不固定的函数形式
	//func m(s ... interface()) return { ... }
	//遍历可变参数列表
	s := joinStrings("a", "b", "c", "d", "e", "f")
	fmt.Println(s)

	//闭包
	c()

	//defer
	deferPanic()

}

const (
	SecondsPerMinute = 60
	SecondsPerHour   = SecondsPerMinute * 60
	SecondsPerDay    = SecondsPerHour * 60
)

func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}

//九九乘法表
func chengfa99biao() {

	for x := 1; x < 10; x++ {
		for y := 1; y <= x; y++ {
			fmt.Print(x * y)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

//遍历可变参数列表
func joinStrings(slist ...string) string {

	//定义一个字节缓冲，快速地连接字符串
	var b bytes.Buffer
	for _, s := range slist {
		b.WriteString(s)
	}
	return b.String()
}

//函数作为参数
func StringProcess(list []string, chain []func(string string) string) {
	for index, str := range list {
		//第一个需要处理的字符串
		result := str
		//遍历每一个处理链
		for _, proc := range chain {
			result = proc(result)
		}
		//将处理后的结果放回切片中
		list[index] = result
	}
}

//自定义处理函数 -- 高阶函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

//函数的返回值是函数
func f1(x, y, z int) int {
	return x + y + z
}
func f2(x int) func(int, int, int) int {
	fmt.Println(x)
	return f1
}
func f3(z int) (x, y int) {
	x = 5
	y = z + 1
	return
}

//闭包
func playerGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}

//闭包
func bibao(x int, y int) func() {
	return func() {
		fmt.Println(x, y)
	}
}
func c() {
	f1 := bibao(1, 2)
	demo1(f1)
}
func demo1(f func()) {
	f()
}

////自定义错误
//var err = errors.New("这是我的自定义错误")
//fmt.Println(err)

//使用goto处理错误
func testGoto() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				goto breakHear
			}
		}
	}
	return
breakHear:
	fmt.Println("donbe")
}

//宕机和从宕机中恢复
//panic   recover(尽量避免使用)
func testPanic() {
	//宕机
	panic("crash")

	//从宕机中恢复
	fmt.Println("运行前")

	ProtectRun(func() {
		fmt.Println("手动宕机前")
		//使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})
		fmt.Println("手动宕机后")
	})
	//故意造成空指针异常
	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}

//panic时需要传递的上下文信息
type panicContext struct {
	function string
}

//保护方式允许一个函数
func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error: ", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()
}

//延时加载 -- defer（多用于释放资源）
func deferPanic() {

	fmt.Println("开始执行操作----")
	time.Sleep(time.Second)
	fmt.Println("函数执行中------")
	time.Sleep(time.Second * 2)
	defer fmt.Println("延时操作1-----")
	defer fmt.Println("延时操作2-----")
	fmt.Println("函数执行结束-----")

}

//求字符串中汉子的长度
func calsChinese(s string) int {
	i := 0
	for _, ss := range s {
		if unicode.Is(unicode.Han, ss) {
			i++
		}
	}
	return i
}

//回文判断
func huiwen(s string) bool {
	var r []rune
	i := 0

	for _, k := range s {
		r = append(r, k)
	}
	j := len(r) - i - 1

	for j > i {
		if r[i] != r[j] {
			return false
		}
		i++
		j--
	}
	return true
}
