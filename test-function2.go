package main

import (
	"fmt"
	"strings"
)

//函数练习
func main() {

	//字符串处理链函数
	list := []string{
		"go root",
		"go path",
		"deadman must be destroyed",
		"go through",
		"heil Hitler!",
	}

	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
		strings.ToTitle,
	}
	//处理字符串
	StringProcess(list, chain)
	//输出处理完的字符串
	for _, str := range list {
		fmt.Println(str)
	}

	//跳转到指定代码标签（goto）
	var breakAgain bool
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				breakAgain = true
				break
			}
		}
		if breakAgain {
			break
		}
	}
	fmt.Println("done")

	//处理科学计数法
	var (
		old = "1.000002e+09"
		f   float64
	)
	n, err := fmt.Sscanf(old, "%e", &f)
	if err != nil {
		fmt.Println(err)
	} else if 1 != n {
		fmt.Println("n is not one !")
	}
	fmt.Println(uint64(f))

}
