package main

import (
	"fmt"
	"os"
)

var (
	students = make(map[int64]*student)
)

type student struct {
	index int64
	name  string
}

//一个练习
func main() {
	//显示菜单
	fmt.Println("welcome!")
	fmt.Println(
		`1.查看学生
				2.新增学生
				3.删除学生`)
	for {
		input := prepare()
		//用户选择功能
		switch input {
		case 1:
			show()
		case 2:
			add()
		case 3:
			deleteStu()
		case 4:
			fmt.Println("退出！")
			os.Exit(205)
		default:
			fmt.Println("选择错误！！")
		}
	}
}
func prepare() int {
	fmt.Print("请选择操作：")
	var input int
	fmt.Scanln(&input)
	fmt.Printf("您选择了 %d \n", input)
	return input
}

func show() {
	for _, v := range students {
		fmt.Printf("姓名：%v ,学号： %d \n", v.name, v.index)
	}
}
func add() {
	var name string
	var index int64
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入学号：")
	fmt.Scanln(&index)

	s := &student{index, name}
	b := false
	for k, _ := range students {
		if k == s.index {
			b = true
			break
		}

	}
	if b {
		fmt.Println("该学号已经被录入！")
	} else {
		students[s.index] = s
		fmt.Println("学生录入成功！")
	}
}
func deleteStu() {
	var index int64
	fmt.Print("请输入要删除的学号：")
	fmt.Scanln(&index)
	delete(students, index)
}
