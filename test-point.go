package main

import (
	"flag"
	"fmt"
)

//指针
//go中指针操作很简单，只有取地址和取地址对应所在值，无法移动指针变量
func main() {

	//指针
	//go语言中不存在指针操作
	c := "按时打发"
	p := &c
	fmt.Printf("p的类型 %T \n", p)
	fmt.Println(p)
	fmt.Println(*p)

	//new 申请内存空间,一般用于给基本数据类型申请内存
	var a = new(int) //a是指针
	*a = 1000
	fmt.Println(a)
	fmt.Println(*a)

	/**
	指针
	*/
	var cat = 1
	var s1 = "banana"
	fmt.Printf("%p  %p  ", &cat, &s1)
	/**
	从指针获取指针指向的值
	*/
	var str = "time is money"
	point := &str
	fmt.Println(point)
	fmt.Println(*point)
	/**
	使用指针获取命令行输入信息（flag包）
	*/
	var mode = flag.String("mode", "", "process mode")
	//解析命令行
	flag.Parse()
	fmt.Println(*mode)
	/**
	用new()函数创建指针
	//*/
	////str := new(string)
	////*str = "ninjia"
	////fmt.Println(*str)

}
