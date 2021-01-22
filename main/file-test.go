package main

import "go-test/file_"

//main 文件操作
func main() {
	//基础版读文件
	//file_.ReadFile_Test("D:/studentSms.go")
	//系统现有工具文件读取,整个文件读取
	//file_.ReadFile_Test2("D:/studentSms.go")
	//将文件读进缓冲区
	//file_.ReadFile_Test3("D:/studentSms.go")
	file_.BufIoread()
}
