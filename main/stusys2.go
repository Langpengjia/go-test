package main

import (
	"fmt"
	"go-test/demo"
	"os"
	"time"
)

var (
	class = demo.PrepareClass("学生管理系统")
)

//另一个小练习
func main() {
	fmt.Println("欢迎使用" + class.Clazzname + "！")
	for {
		fmt.Println("菜单：")
		fmt.Println(`1. 查询学生信息：
2. 添加学生信息：
3. 删除学生：
4. 学生列表：
5. 退出系统。`)
		var choose int
		fmt.Print("请选择要进行的操作：")
		_, err := fmt.Scanln(&choose)
		if err != nil {
			os.Exit(500)
		}
		switch choose {
		case 1:
			sysFind()
		case 2:
			sysAdd()
		case 3:
			sysDelete()
		case 4:
			sysList()
		case 5:
			fmt.Println("系统正在退出...")
			time.Sleep(time.Second * 2)
			os.Exit(200)
		}

		time.Sleep(time.Second * 2)
	}

}
func sysList() {
	class.ListStus()
}
func sysAdd() {
	var name string
	var index int
	var gender string

	fmt.Print("请输入学号：")
	fmt.Scanln(&index)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入性别：")
	fmt.Scanln(&gender)
	class.AddStudent(name, index, gender)
}
func sysFind() {
	var index int
	fmt.Print("请输入学号：")
	fmt.Scanln(&index)

	stu := class.FindStu(index)
	if stu != nil {
		fmt.Printf("学生姓名：%s，学号：%d，性别：%s。\n",
			stu.Name, stu.Index, stu.Gender)
	} else {
		fmt.Println("未找到对应学号的学生！")
	}

}
func sysDelete() {
	var index int
	fmt.Print("请输入学号：")
	fmt.Scanln(&index)
	class.DeleteStudent(index)

}
