package file_

import (
	"bufio"
	"fmt"
	"os"
)

func BufIoread() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容：")
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(readString + "    asdasdasd")

}
