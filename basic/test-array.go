package basic

import "fmt"

//数组
//数组是值类型
//数组局限性
//数组的类型包括数组元素的类型和数组的长度
//数组作为参数时收到数组长度的约束
func main() {

	//数组一般声明
	a1 := [3]bool{true, true, true}
	fmt.Println(a1)
	//根据初始值自动推断数组的长度
	a100 := [...]int{1, 4, 15, 24, 51, 55}
	fmt.Println(a100)
	a2 := [5]int{1, 2}
	fmt.Println(a2)
	//根据索引指定数组的值
	an := [5]int{1, 4: 2}
	fmt.Println(an)

	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[2] = 100
	fmt.Println(b1, b2)

	test1()
	test2()

	//数组求和
	arr := &[15]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 0, 5, 65, 74}
	sumd := sumArr(arr)
	fmt.Printf("数组和为 %v \n", sumd)
	findArrSum(5, arr)

}

//数组的遍历
func test1() {
	//1
	citys := [...]string{"北京", "上海", "广州"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2
	for i, v := range citys {
		fmt.Println(i, v)
	}

}

//多维数组
func test2() {
	var aii [3][3]int
	fmt.Println(aii)
	ss := [3][2]int{
		[2]int{0, 1},
		[2]int{0, 1},
		[2]int{0, 1},
	}
	fmt.Println(ss)
}

//练习
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
