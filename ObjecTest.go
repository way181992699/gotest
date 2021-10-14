package main

import "fmt"

type People struct {
	age  int
	name string
}

func change(p People) {
	p.age = 18
}

func change2(p *People) {
	p.age = 18
}

func add(x, j int) {
	x++
}

func add1(x *int, j int) {
	*x = j + 10
}

/**
1.change函数是传递的对象，函数调用的时候，会拿到对象的拷贝。
2.change2函数是传递的指针，函数调用的时候，会拿到一个指向改对象的指针。
3.go没有引用传递
*/
func main() {
	//1.引用类型作为参数传入
	people := People{31, "july"}
	change(people)
	fmt.Println(people)
	change2(&people)
	fmt.Println(people)
	//2.基本类型作为参数传入
	var x = 10
	var j = 20
	add(x, j)
	fmt.Println(x)
	add1(&x, j)
	fmt.Println(x)
	//结论 两者都一样
}
