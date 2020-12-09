// go-test project main.go
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"runtime"
	"strings"
)

func main() {

	//数组是值类型

	//数组一般声明
	//a1 := [3]bool{true, true, true}
	//fmt.Println(a1)
	////根据初始值自动推断数组的长度
	//a100 := [...]int{1, 4, 15, 24, 51, 55}
	//fmt.Println(a100)
	//a2 := [5]int{1,2}
	//fmt.Println(a2)
	////根据索引指定数组的值
	//an := [5]int{1,4:2}
	//fmt.Println(an)

	////数组的遍历
	////1
	//citys := [...]string{"北京","上海","广州"}
	//for i := 0; i < len(citys); i++ {
	//	fmt.Println(citys[i])
	//}
	////2
	//for i,v:=range citys{
	//	fmt.Println(i,v)
	//}

	//多维数组
	//var aii [3][3]int
	//fmt.Println(aii)
	//ss := [3][2]int{
	//	[2]int{0,1},
	//	[2]int{0,1},
	//	[2]int{0,1},
	//}
	//fmt.Println(ss)

	//b1 := [3]int{1,2,3}
	//b2 := b1
	//b2[2] = 100
	//fmt.Println(b1,b2)

	////练习
	//arr := &[15]int{0,1,2,3,4,5,6,7,8,9,1,0,5,65,74}
	//sumd := sumArr(arr)
	//fmt.Printf("数组和为 %v \n",sumd)
	//findArrSum(5,arr)

	//数组局限性
	//数组的类型包括数组元素的类型和数组的长度
	//数组作为参数时收到数组长度的约束

	//切片（和数组的区别就是不设置长度）
	//切片的本质，框柱了一块连续的内存,真正的数据是保存在底层数组里的
	//一个nil的切片是没有底层数组的
	//定义一个切片
	//var s1 []int
	//var s2 []string
	//fmt.Printf("s1的类型为：%T ，s2的类型为: %T",s1,s2)

	//s3:=[]int{1,2,3,4,5,6,7,8,9,10}
	//fmt.Println(cap(s3))
	//s4 := s3[1:3]
	//s5:= s3[6:]
	//s6:= s4[:2]
	//s6[0] = 10086
	//fmt.Println(s3,len(s3),cap(s3))
	//fmt.Println(s4,len(s4),cap(s4))
	//fmt.Println(s5,len(s5),cap(s5))
	//fmt.Println(s6,len(s6),cap(s6))

	//用make构造切片
	sn := make([]int, 5, 15)
	sn[1] = 21
	fmt.Println(sn, len(sn), cap(sn))

	/**
	变量交换，跨文件访问变量
	*/
	//var a string = "a"
	//var b string = "b"
	//var c string = "c"
	//var d struct {
	//	s string
	//	x int
	//}

	//a,b = b,a
	//a,b,c = c,a,b
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	/**
	搭建一个简单的服务器
	*/
	//http.Handle("/", http.FileServer(http.Dir(".")))
	//http.ListenAndServe(":8111", nil)
	/**
	一个简单的demo
	*/
	//var a, b, c int = 1, 2, 3
	//fmt.Println(a)
	//go log(b)
	//fmt.Println(c)
	//
	//fmt.Println("Hello World!")

	/**
	demo2 输出一个正弦函数
	*/
	//makeSin()
	/**
	符号引用,无视转义
	*/
	//s := `
	//	第一行
	//	第二行
	//	第三行
	//	第四行
	//	/n/t/f/s`
	//fmt.Println(s)
	/**
	转换不同的数据类型
	*/
	//fmt.Println(" int8 ranges:",math.MinInt8,math.MaxInt8)
	//fmt.Println(" int16 ranges:",math.MinInt16,math.MaxInt16)
	//fmt.Println(" int32 ranges:",math.MinInt32,math.MaxInt32)
	//fmt.Println(" int64 ranges:",math.MinInt64,math.MaxInt64)
	//var a int32 = 1047483674
	//fmt.Printf("int32: 0x%x %d\n",a,a)
	//b := int16(a)
	//fmt.Printf("int16: 0x%x %d\n",b,b)
	//var c float32 = math.Pi
	//fmt.Println(int(c))
	/**
	指针
	*/
	//var cat = 1
	//var str = "banana"
	//fmt.Printf("%p  %p  ",&cat,&str)
	/**
	从指针获取指针指向的值
	*/
	//var str = "time is money"
	//point := &str
	//fmt.Println(point)
	//fmt.Println(*point)
	/**
	使用指针获取命令行输入信息（flag包）
	*/
	//var mode = flag.String("mode","","process mode")
	////解析命令行
	//flag.Parse()
	//fmt.Println(*mode)
	/**
	用new()函数创建指针
	*/
	//str := new(string)
	//*str = "ninjia"
	//fmt.Println(*str)

	/**
	变量和栈
	*/
	//fmt.Println(calc(10,2))
	/**
	变量逃逸分析
	*/
	//var a int
	//void()
	//fmt.Println(a,dummy(0))
	/**
	取地址发生逃逸
	*/
	//fmt.Println(dummy1())
	/**
	字符串应用
	*/
	////计算字符串长度  汉字长度为3
	//tip1 := "asdbh"
	//tip2 := "前不见古人，后不见来者，念天地之悠悠，独怆然而涕下 "
	//tip3 := "读"
	//fmt.Println(len(tip1))
	//fmt.Println(len(tip2))
	//fmt.Println(len(tip3))

	////ASCII 遍历字符串
	//theme := "狙击 start!"
	//for i:=0;i<len(theme);i++{
	//	fmt.Printf("ascii : %c  %d \n",theme[i],theme[i])
	//}
	////Unicode 遍历字符串
	//for _, s:=range theme {
	//	fmt.Printf("Unicode:%c  %d\n",s,s)
	//}

	////字符串截取
	//tracer := "死神来了，死神bye bye!"
	//comma := strings.Index(tracer, "，")
	//fmt.Println(tracer[comma:])
	//pos := strings.Index(tracer[comma:],"死神")
	//fmt.Println(comma,pos,tracer[comma+pos:])

	////修改字符串,字符串不可变，生成新的字符串
	//angel := "Heros never die"
	//angleBytes := []byte(angel)
	//for i := 5; i <= 10; i++ {
	//	angleBytes[i] = ' '
	//}
	//fmt.Println(string(angleBytes))

	////字符串拼接
	//hammer := "吃我一锤"
	//sickle := "死吧"
	////声明字节缓冲
	//var stringBuilder bytes.Buffer
	////把字符串写入缓冲
	//stringBuilder.WriteString(hammer)
	//stringBuilder.WriteString(sickle)
	//fmt.Println(hammer+sickle)
	//fmt.Println(stringBuilder.String())

	//格式化，跳过

	////Base64解码
	//message := "Away from keyboard. https://golang.org/"
	////编码，输出
	//encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))
	//fmt.Println(encodeMessage)
	////解码，输出
	//data,err := base64.StdEncoding.DecodeString(encodeMessage)
	////出错处理
	//if err!=nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(string(data))
	//}

	//从INI配置文件中查询需要的值

	////常量
	//const  a = "asdad"
	//const b = 111
	//const (
	//	c = &a
	//	d = &b
	//)

	//枚举，用其他方式替代
	//fmt.Printf("%s  %d",CPU,CPU)
	//fmt.Printf("%s  %d",GPU,GPU)

	////类型别名 Type Alias
	//var a newInt
	//fmt.Printf("a type: %T\n", a)
	//
	//var a2 intAlias
	//fmt.Printf("a2 type: %T\n", a2)

	////在结构体成员嵌入时使用别名
	//var a Vehicle
	//a.FakeBrand.show()
	//ta := reflect.TypeOf(a)
	//for i := 0; i < ta.NumField(); i++ {
	//	f := ta.Field(i)
	//	fmt.Println("FieldName :  %v, FieldType: %v\n",f.Name,f.Type.Name())
	//}

	////数组
	//var team = [3]int{1,2,3}
	//for i := 0; i < len(team); i++ {
	//	fmt.Println(team)
	//}

	////切片复制元素
	//const elementCount = 1000
	//srcData := make([]int, elementCount)
	//
	////将切片赋值
	//for i := 0; i < elementCount; i++ {
	//	srcData[i] = i
	//}

	////引用切片数据,复制切片数据
	//refData := srcData
	//
	//copyData := make([]int,elementCount)
	//
	//copy(copyData,srcData)
	//
	//srcData[0] = 999
	//
	//fmt.Println(refData[0])
	//
	//fmt.Println(copyData[0],copyData[elementCount-1])
	//
	//copy(copyData,srcData[4:6])
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("%d",copyData[i])
	//}

	////99乘法表
	//chengfa99biao()

	////遍历字符串
	//str:= "无名以观其妙，有名以观其徼"
	//for key,value := range str{
	//	fmt.Printf("key: %d  value : 0x%x\n  ",key,value)
	//}

	////遍历Map 获得map中的键和值
	//m := map[string]int{
	//	"hello": 100,
	//	"world": 200,
	//}
	////改版map的值
	//m["every"] = 300
	//m["hello"] = 0
	//for k,v:=range m{
	//	fmt.Println(k,v)
	//}

	////遍历通道
	//c := make(chan int)
	//go func() {
	//	c <- 1
	//	c <- 2
	//	c <- 3
	//	close(c)
	//}()
	//
	//for i:= range c{
	//	fmt.Println(i)
	//}

	//switch
	////一分支多值
	//a := "daddy"
	//switch a {
	//case "mum","daddy":
	//	fmt.Println("family!")
	//}
	////分支表达式
	//r := 11
	//switch {
	//case r > 10 && r < 20:
	//	fmt.Println(r)
	//}
	////case使用fallthrough
	//s := "hello"
	//switch  {
	//case s == "hello":
	//	fmt.Println("hello")
	//	fallthrough
	//case s != "world":
	//	fmt.Println("world")
	//}

	////跳转到指定代码标签（goto）
	//var breakAgain bool
	//for x := 0; x < 10; x++ {
	//	for y := 0; y < 10; y++ {
	//		if y == 2 {
	//			breakAgain = true
	//			break
	//		}
	//	}
	//	if breakAgain {
	//		break
	//	}
	//}
	//fmt.Println("done")

	//	//使用goto处理错误
	//	for x := 0; x < 10; x++ {
	//		for y := 0; y < 10; y++ {
	//			if y == 2 {
	//				goto breakHear
	//			}
	//		}
	//	}
	//	return
	//breakHear:
	//	fmt.Println("donbe")

	////将函数作为变量
	//var f func()
	//f= chengfa99biao
	//f()

	////匿名函数
	//visit([]int{1,2,3,4}, func(v int) {
	//	fmt.Println(v)
	//})

	//in := Data1{
	//	complex: []int{1, 2, 3},
	//	instance: InnerData{
	//		5,
	//	},
	//	prt: &InnerData{1},
	//}
	//

	////函数中参数传递效果测试
	//fmt.Printf("in value：%+v\n", in)
	//
	//fmt.Printf("in ptr: %p\n", &in)
	//
	//out := passByValue(in)
	//fmt.Printf("out value:  %+v\n", out)
	//
	//fmt.Printf("out ptr:  %p\n", &out)

	////函数定义
	//fmt.Println(resolveTime(25221511))
	//
	//_,hour,minute := resolveTime(10036)
	//
	//fmt.Println(hour,minute)

	////字符串处理链函数
	//list := []string{
	//	"go root",
	//	"go path",
	//	"deadman must be destroyed",
	//	"go through",
	//	"heil Hitler!",
	//}
	//
	//chain := []func(string) string{
	//	removePrefix,
	//	strings.TrimSpace,
	//	strings.ToUpper,
	//	strings.ToTitle,
	//}
	////处理字符串
	//StringProcess(list, chain)
	////输出处理完的字符串
	//for _, str := range list {
	//	fmt.Println(str)
	//}

	//匿名函数（匿名函数可以在声明后调用）
	//d:= func(data int) int{
	//	fmt.Println("hello",data)
	//	return data * 10
	//}(100)
	//fmt.Println(d)
	//f := func(a string){
	//	fmt.Println("this is the letter",a)
	//}
	//f("s")

	////使用匿名函数实现操作封装
	//generator := playerGen("high noon!")
	//name, hp := generator()
	//fmt.Println(name, hp)

	//可变参数--参数不固定的函数形式
	//func m(s ... interface()) return { ... }

	////遍历可变参数列表
	//f := joinStrings("a","b","c","d","e","f")
	//fmt.Println(f)

	////延迟执行语句-- defer
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)

	//使用延迟执行语句在函数退出时释放资源

	////自定义错误
	//var err = errors.New("这是我的自定义错误")
	//fmt.Println(err)

	////宕机
	//panic("crash")

	////从宕机中恢复
	//fmt.Println("运行前")
	//
	//ProtectRun(func() {
	//	fmt.Println("手动宕机前")
	//	//使用panic传递上下文
	//	panic(&panicContext{
	//		"手动触发panic",
	//	})
	//	fmt.Println("手动宕机后")
	//})
	////故意造成空指针异常
	//ProtectRun(func() {
	//	fmt.Println("赋值宕机前")
	//	var a *int
	//	*a = 1
	//	fmt.Println("赋值宕机后")
	//})
	//fmt.Println("运行后")

	////结构体
	//cat := newCatByName("喵喵喵")
	//fmt.Println(cat)

	//fmt.Println(fof, asd, asdwg, adtesg,sss,asdw)
	//fmt.Println(a1,a2,a3,a4,a5)

	//a := "赵客缦胡缨"
	//as := []rune(a)
	//as[0] = '唐'
	//fmt.Println(string(as))
	//
	//c1 := "唐"
	//c2 := '唐'
	//c3 := rune('糖')
	//fmt.Printf("%T \n %T \n %T \n",c1,c2,c3)

	var a [3]bool
	var b [4]bool
	fmt.Println(a, b)

}

//数组求和
func sumArr(arr *[15]int) int {
	var i int
	for _, v := range *arr {
		i += v
	}
	return i
}

//找出数组中和为指定值的两个数组元素的下标
func findArrSum(target int, arr *[15]int) {
	for i := 0; i < len(*arr); i++ {
		for j := i + 1; j < len(*arr); j++ {
			if target == arr[i]+arr[j] {
				fmt.Print(arr[i], arr[j], " :")
				fmt.Println(i, j)
			}
		}
	}
}

const (
	fof, coc = iota, iota
	assss    = iota
	sss      = 85174
	asdw     = iota
)

const (
	a1 = iota
	a2
	a3
	a4 = 10086
	a5
)

//结构体demo--二维矢量模拟玩家移动
//实现二维矢量解构
type Vec2 struct {
	X, Y float32
}

//矢量加法
func (v Vec2) add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

//矢量减法
func (v Vec2) minus(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

//矢量相乘
func (v Vec2) scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

//计算两个矢量的距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy + dy)))
}

//

//结构体
type Cat struct {
	Color string
	Name  string
}

func newCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func newCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
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

////使用延迟执行语句在函数退出时释放资源.
//
//var(
//	valueByKey = make(map[string]int)
//	valueByKeyGuard sync.Mutex
//)
//func readValue(key string)int{
//	valueByKeyGuard.Lock()
//	v := valueByKey(key)
//	valueByKeyGuard.Unlock()
//	return v
//}

//遍历可变参数列表
func joinStrings(slist ...string) string {

	//定义一个字节缓冲，快速地连接字符串
	var b bytes.Buffer

	for _, s := range slist {
		b.WriteString(s)
	}
	return b.String()
}

func playerGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}

//声明常量
const (
	SecondsPerMinute = 60
	SecondsPerHour   = SecondsPerMinute * 60
	SecondsPerDay    = SecondsPerHour * 60
)

var skillParam = "s"

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

//自定义处理函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func resolveTime(seconds int) (day, hour, minute int) {

	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}

func passByValue(inFunc Data1) Data1 {

	fmt.Printf("inFunc value: %+v\n", inFunc)

	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}

type Data1 struct {
	complex  []int
	instance InnerData
	prt      *InnerData
}

type InnerData struct {
	a int
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
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

//在结构体成员嵌入时使用别名
type Brand struct {
}

func (a Brand) show() {
}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

//枚举，用其他方式替代
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

type newInt int
type intAlias = int

type Data struct {
}

func dummy1() *Data {
	var c Data
	return &c
}

func dummy(b int) int {
	var c int
	c = b
	return c
}

func calc(a, b int) int {
	var c int
	c = a * b
	var x int
	x = c * 10
	return x
}

func makeSin() {
	//绘制背景
	const size = 1000
	pic := image.NewGray(image.Rect(0, 0, size, size))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	//绘制正弦函数轨迹
	for x := 0; x < size; x++ {
		//让S的值处于0和2pi之间
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y), color.Gray{0})
	}

	//写入图片文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, pic)
	//关闭文件、
	file.Close()
}

func glog(b int) {
	fmt.Println(b)

}
