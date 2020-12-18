package main

import "fmt"

func main() {

	test2()

}

func test1() {
	var m map[string]interface{}
	m = make(map[string]interface{})

	m["a"] = "天下风云出我辈"
	m["b"] = 15
	m["c"] = []int{1, 2, 3, 4, 5}
	m["d"] = [...]string{"a", "b", "c", "d", "2"}
	m["e"] = true
	fmt.Printf(" 类型： %T， 值：%v \n ", m, m)
	for k, v := range m {
		fmt.Printf(k + "  ")
		fmt.Printf(" 类型： %T， 值：%v \n ", v, v)
	}
}

func test2() {

	a := "天下风云出我辈"
	b := 15
	c := []int{1, 2, 3, 4, 5}
	d := [...]string{"a", "b", "c", "d", "2"}
	e := true
	showx(a)
	showx(b)
	showx(c)
	showx(d)
	showx(e)

}
func showx(x interface{}) {
	fmt.Println(x)
}
