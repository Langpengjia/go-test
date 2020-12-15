package main

import (
	"fmt"
	"sort"
)

//切片
//切片（和数组的区别就是不设置长度）
//切片的本质，框柱了一块连续的内存,真正的数据是保存在底层数组里的
//切片是一个引用类型，不保存值，值是底层数组决定的
//一个nil的切片是没有底层数组的
func main() {

	//定义一个切片
	var i1 []int
	var i2 []string
	fmt.Printf("s1的类型为：%T ，s2的类型为: %T", i1, i2)
	i3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(cap(i3))
	i4 := i3[1:3]
	i5 := i3[6:]
	i6 := i4[:2]
	//i7[0] = 10086
	fmt.Println(i3, len(i3), cap(i3))
	fmt.Println(i4, len(i4), cap(i4))
	fmt.Println(i5, len(i5), cap(i5))
	fmt.Println(i6, len(i6), cap(i6))

	//用make构造切片
	sn := make([]int, 5, 15)
	sn[1] = 21
	fmt.Println(sn, len(sn), cap(sn))

	//切片扩容 append,append操作会改变底层数组
	s1 := []string{"北京", "shanghai", "shenzhen", "guangzhou", "guangzhou"}
	fmt.Println(s1, len(s1), cap(s1))
	//扩容时原有数组容量会翻倍！！！
	s1 = append(s1, "guangzhou")
	fmt.Println(s1, len(s1), cap(s1))

	//-------注意！！！！！！！！
	//必须要用变量接收append参数的返回值
	//ss3 := make([]int,5,10)
	i1 = append(i1, 1, 52, 47, 6578, 248, 1475, 13687, 1347, 521, 852, 45, 81)
	sort.Ints(i1)
	fmt.Println(i1, len(i1), cap(i1))
	//此处修改了底层数组，数值发生了平移
	s2 := append(s1[0:8], s1[10:]...)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	//... 相当于拆开前面变量中的元素
	//copy
	a1 := []int{1, 2, 3}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//切片复制元素
	const elementCount = 1000
	srcData := make([]int, elementCount)

	//将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	//引用切片数据,复制切片数据
	refData := srcData

	copyData := make([]int, elementCount)

	copy(copyData, srcData)

	srcData[0] = 999

	fmt.Println(refData[0])

	fmt.Println(copyData[0], copyData[elementCount-1])

	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d", copyData[i])
	}

}
