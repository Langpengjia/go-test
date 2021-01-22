package file_

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	length = 128
)

//基础班读文件
func ReadFile_Test(path string) {

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var b [length]byte
	for {
		read, err := f.Read(b[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b[:read]))
		if read < length {
			return
		}
	}
	////工具版
	//b, err := ioutil.ReadFile(path)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}else {
	//	fmt.Println(string(b))
	//}
}

//系统现有工具文件读取
func ReadFile_Test2(path string) {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(b))
	}
}

//读文件至缓冲区
func ReadFile_Test3(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//读一行
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Printf("读文件异常：", err)
		}
		fmt.Printf(readString)
	}

}
