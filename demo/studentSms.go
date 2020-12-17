package demo

import "fmt"

//Student 学生类型
type Student struct {
	//姓名
	Name string
	//学号
	Index int
	//性别
	Gender string
}

//Class 系统结构体
type Class struct {
	Clazzname string
	Stus      []*Student
}

//PrepareClass 初始化系统
func PrepareClass(s string) *Class {
	return &Class{
		Clazzname: s,
	}
}

func (c *Class) ListStus() {
	if len(c.Stus) == 0 {
		fmt.Println("尚未录入学生！")
	} else {
		for _, stu := range c.Stus {
			fmt.Printf("学生姓名：%s，学号：%d，性别：%s。\n",
				stu.Name, stu.Index, stu.Gender)
		}
	}
}

//AddStudent 添加一个学生
func (c *Class) AddStudent(name string, index int, gender string) {
	c.Stus = append(c.Stus, &Student{name, index, gender})
}

//FindStu 查找单一学生
func (c *Class) FindStu(index int) *Student {
	var stu *Student
	for _, v := range c.Stus {
		if index == v.Index {
			stu = v
			break
		}
	}
	return stu
}

//DeleteStudent 删除一个学生
func (c *Class) DeleteStudent(index int) {
	i := 0
	b := false
	for i < len(c.Stus) {
		if c.Stus[i].Index == index {
			b = true
			break
		}
		i++
	}
	if b {
		c.Stus = append(c.Stus[:i], c.Stus[i+1:]...)
		fmt.Println("学生删除成功！")
	} else {
		fmt.Println("未找到对应学号的学生！")
	}
}
