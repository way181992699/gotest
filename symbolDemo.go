package main

import "fmt"

//特俗符号
func main() {
	var a int = 1
	var b *int = &a
	var d *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
	fmt.Println("d = ", d)
	fmt.Println("*d = ", *d)
	fmt.Println("*&d = ", *&d)
	s1 := `第一行
		第二行
第三行
`
	s2 := "第一行"
	s3 := "aaa"
	s4 := "第1行"

	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
	//1个汉字暂用3个长度。 字母和数字占用1个长度
	fmt.Println("s1 len is ", len(s1))
	fmt.Println("s2 len is ", len(s2), s2)
	fmt.Println("s3 len is ", len(s3), s3)
	fmt.Println("s4 len is ", len(s4), s4)

}
