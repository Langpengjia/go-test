// go-test project main.go
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

//此文件主要用于测试

func main() {

}

//func main() {
//	//var d struct_.Person3
//	//d.Wang()
//	//
//	//p := struct_.NewPersion3("张嘉倪",18)
//	//p.Wang()
//
//	struct_.Test2()
//
//	////make 也是分配类型的 只作用于 slice map chan
//	//arr := make([]int, 3, 10)
//	//fmt.Println(arr)
//
//	////  map 内部以hash实现
//	//var m map[string]int
//	//fmt.Println(m == nil)
//	//m = make(map[string]int, 5)
//	//m["a"] = 1
//	//m["b"] = 2
//	//fmt.Println(m)
//	//v,ok := m["c"]
//	//if ok{
//	//	fmt.Println(v)
//	//}
//	//delete(m,"a")
//	//
//	//
//	//m["sou"] = 214
//	//m["asda"] = 241
//	//fmt.Println(m)
//	//fmt.Println(len(m))
//	//for k,v :=range m{
//	//	fmt.Println(k,v)
//	//}
//
//	////判断字符中汉子的长度
//	//s:="asdadcvs赵姐阿萨法wang史蒂夫"
//	//fmt.Println(calsChinese(s))
//
//	////判断回文数
//	//s1 := "上海自来水来自海上"
//	//s2 := "asdafswfa"
//	//fmt.Println(huiwen(s1))
//	//fmt.Println(huiwen(s2))
//
//	//f := f2(1)(1, 2, 3)
//	//fmt.Println(f)
//
//	//x,y:=f3(0)
//	//fmt.Println(x,y)
//	////函数定义
//	//c:= func(x int) {
//	//	tmp := func() {
//	//		fmt.Println(x)
//	//	}
//	//	tmp()
//	//}
//
//	/**
//	变量交换，跨文件访问变量
//	*/
//	//var a string = "a"
//	//var b string = "b"
//	//var c string = "c"
//
//	//a,b = b,a
//	//a,b,c = c,a,b
//	//fmt.Println(a)
//	//fmt.Println(b)
//	//fmt.Println(c)
//	//fmt.Println(d)
//	/**
//	搭建一个简单的服务器
//	*/
//	//http.Handle("/", http.FileServer(http.Dir(".")))
//	//http.ListenAndServe(":8111", nil)
//
//	/**
//	demo2 输出一个正弦函数
//	*/
//	//makeSin()
//
//	/**
//	/**
//	变量和栈
//	*/
//	//fmt.Println(calc(10,2))
//	/**
//	变量逃逸分析
//	*/
//	//var a int
//	//void()
//	//fmt.Println(a,dummy(0))
//	/**
//	取地址发生逃逸
//	*/
//	//fmt.Println(dummy1())
//	/**
//	字符串应用
//	*/
//	////计算字符串长度  汉字长度为3
//	//tip1 := "asdbh"
//	//tip2 := "前不见古人，后不见来者，念天地之悠悠，独怆然而涕下 "
//	//tip3 := "读"
//	//fmt.Println(len(tip1))
//	//fmt.Println(len(tip2))
//	//fmt.Println(len(tip3))
//
//	////ASCII 遍历字符串
//	//theme := "狙击 start!"
//	//for i:=0;i<len(theme);i++{
//	//	fmt.Printf("ascii : %c  %d \n",theme[i],theme[i])
//	//}
//	////Unicode 遍历字符串
//	//for _, s:=range theme {
//	//	fmt.Printf("Unicode:%c  %d\n",s,s)
//	//}
//
//	////字符串截取
//	//tracer := "死神来了，死神bye bye!"
//	//comma := strings.Index(tracer, "，")
//	//fmt.Println(tracer[comma:])
//	//pos := strings.Index(tracer[comma:],"死神")
//	//fmt.Println(comma,pos,tracer[comma+pos:])
//
//	////修改字符串,字符串不可变，生成新的字符串
//	//angel := "Heros never die"
//	//angleBytes := []byte(angel)
//	//for i := 5; i <= 10; i++ {
//	//	angleBytes[i] = ' '
//	//}
//	//fmt.Println(string(angleBytes))
//
//	////字符串拼接
//	//hammer := "吃我一锤"
//	//sickle := "死吧"
//	////声明字节缓冲
//	//var stringBuilder bytes.Buffer
//	////把字符串写入缓冲
//	//stringBuilder.WriteString(hammer)
//	//stringBuilder.WriteString(sickle)
//	//fmt.Println(hammer+sickle)
//	//fmt.Println(stringBuilder.String())
//
//	//格式化，跳过
//
//	////Base64解码
//	//message := "Away from keyboard. https://golang.org/"
//	////编码，输出
//	//encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))
//	//fmt.Println(encodeMessage)
//	////解码，输出
//	//data,err := base64.StdEncoding.DecodeString(encodeMessage)
//	////出错处理
//	//if err!=nil{
//	//	fmt.Println(err)
//	//}else {
//	//	fmt.Println(string(data))
//	//}
//
//	//从INI配置文件中查询需要的值
//
//	////常量
//	//const  a = "asdad"
//	//const b = 111
//	//const (
//	//	c = &a
//	//	d = &b
//	//)
//
//	//枚举，用其他方式替代
//	//fmt.Printf("%s  %d",CPU,CPU)
//	//fmt.Printf("%s  %d",GPU,GPU)
//
//	////类型别名 Type Alias
//	//var a newInt
//	//fmt.Printf("a type: %T\n", a)
//	//
//	//var a2 intAlias
//	//fmt.Printf("a2 type: %T\n", a2)
//
//	////在结构体成员嵌入时使用别名
//	//var a Vehicle
//	//a.FakeBrand.show()
//	//ta := reflect.TypeOf(a)
//	//for i := 0; i < ta.NumField(); i++ {
//	//	f := ta.Field(i)
//	//	fmt.Println("FieldName :  %v, FieldType: %v\n",f.Name,f.Type.Name())
//	//}
//
//	////遍历字符串
//	//str:= "无名以观其妙，有名以观其徼"
//	//for key,value := range str{
//	//	fmt.Printf("key: %d  value : 0x%x\n  ",key,value)
//	//}
//
//	////遍历Map 获得map中的键和值
//	//m := map[string]int{
//	//	"hello": 100,
//	//	"world": 200,
//	//}
//	////改版map的值
//	//m["every"] = 300
//	//m["hello"] = 0
//	//for k,v:=range m{
//	//	fmt.Println(k,v)
//	//}
//
//	////遍历通道
//	//c := make(chan int)
//	//go func() {
//	//	c <- 1
//	//	c <- 2
//	//	c <- 3
//	//	close(c)
//	//}()
//	//
//	//for i:= range c{
//	//	fmt.Println(i)
//	//}
//
//	//switch
//	////一分支多值
//	//a := "daddy"
//	//switch a {
//	//case "mum","daddy":
//	//	fmt.Println("family!")
//	//}
//	////分支表达式
//	//r := 11
//	//switch {
//	//case r > 10 && r < 20:
//	//	fmt.Println(r)
//	//}
//	////case使用fallthrough
//	//s := "hello"
//	//switch  {
//	//case s == "hello":
//	//	fmt.Println("hello")
//	//	fallthrough
//	//case s != "world":
//	//	fmt.Println("world")
//	//}
//
//	////结构体
//	//cat := newCatByName("喵喵喵")
//	//fmt.Println(cat)
//
//	//fmt.Println(fof, asd, asdwg, adtesg,sss,asdw)
//	//fmt.Println(a1,a2,a3,a4,a5)
//
//	//a := "赵客缦胡缨"
//	//as := []rune(a)
//	//as[0] = '唐'
//	//fmt.Println(string(as))
//	//
//	//c1 := "唐"
//	//c2 := '唐'
//	//c3 := rune('糖')
//	//fmt.Printf("%T \n %T \n %T \n",c1,c2,c3)
//
//}

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
