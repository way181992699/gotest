package main

import "fmt"

var x int8 = 10

const pi = 3.14

/*
在Go语言程序执行时导入包语句会自动触发包内部的init函数
*/
func init() {
	fmt.Println(x)
}

func main() {
	fmt.Println("开始运行main函数")
	//10
	//开始运行main函数
}
