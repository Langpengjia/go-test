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

type Data struct {
}

func main() {

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

	//切片复制元素
	const elementCount = 1000
	srcData := make([]int, elementCount)

	//将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

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

func void() {

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
