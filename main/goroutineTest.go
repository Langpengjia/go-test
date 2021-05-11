package main

import (
	"fmt"
	"sync"
)

func print(i int) {
	defer group.Done()
	fmt.Println(i)
}

func printRandomInt() {
	//rand.Seed()
}

var group sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		group.Add(i)
		go print(i)
		var mu sync.Mutex
		mu.Lock()

		mu.Unlock()

		//go func(i int) {
		//	fmt.Println(i)
		//}(i)
	}
	group.Wait()
	fmt.Println("执行结束")
}
