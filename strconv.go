package main

import (
	"fmt"
	"strconv"
)

//Go语言中strconv包实现了基本数据类型和其字符串表示的相互转换。
func main() {
	//a的典故
	//【扩展阅读】这是C语言遗留下的典故。C语言中没有string类型而是用字符数组(array)表示字符串，所以Itoa对很多C系的程序员很好理解。
	//1.Atoi(): Array to int 这样记
	str0 := "7"
	fmt.Println(strconv.Atoi(str0)) //7 <nil>
	//2.Itoa(): Int to array
	i := 7
	strconv.Itoa(i)              // "7"
	fmt.Println(strconv.Itoa(i)) //7 <nil>
	//3.ParseBool()把字符串转换为bool值
	bool0 := "0"
	bool1 := "1"
	boolf := "f"
	boolt := "t"
	boolF := "F"
	boolT := "T"
	//他接受: 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	fmt.Println(strconv.ParseBool(bool0)) //false <nil>
	fmt.Println(strconv.ParseBool(bool1)) //true <nil>
	fmt.Println(strconv.ParseBool(boolf)) //false <nil>
	fmt.Println(strconv.ParseBool(boolt)) //true <nil>
	fmt.Println(strconv.ParseBool(boolF)) //false <nil>
	fmt.Println(strconv.ParseBool(boolT)) //true <nil>
	//4.ParseInt() 解析字符串转换为Int值，接受正负号
	//5.ParseUnit() 类似ParseInt()，但不接受正负号，用于无符号整型。
	//6.ParseFloat() 解析字符串转为浮点数
	//7.FormatBool() 转为字符串类型的"true" or "false"
	//8.FormatInt() 转为字符串类型的数值
	//9.FormatUnit() 转为字符串类型的无符号整数版本(绝对值)
	//10.FormatUnit() 转为字符串类型的无符号整数版本(绝对值)
	//上述为常用的函数，若需查看其他用法可以访问：https://pkg.go.dev/strconv
}
