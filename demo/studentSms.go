package demo

//Student 学生类型
type Student struct {
	name   string
	age    int
	gender string
}

//
type Class struct {
	clazzname string
	stus      []*Student
}

func NewStudent(name string, age int, gender string) *Student {
	return &Student{name, age, gender}
}

func (c *Class) addStudent(student *Student) {
	c.stus = append(c.stus, student)
}
func (c *Class) findStu(name string) *Student {
	var stu *Student
	for _, v := range c.stus {
		if name == v.name {
			stu = v
		}
	}
	return stu
}

func (c *Class) deleteStudent(student *Student) {
	i := 0
	b := false
	for i < len(c.stus) {
		if *(c.stus[i]) == *student {
			b = true
			break
		}
		i++
	}
	if b {
		c.stus = append(c.stus[:i], c.stus[i+1:]...)
	}
}
