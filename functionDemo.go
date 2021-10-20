package main

import (
	"errors"
	"fmt"
	"strings"
)

type cb func(int) int

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

var a []string

func main() {

	//函数式编程
	i := func(x, y int) int {
		return x * y
	}
	fmt.Println(i(4, 5)) //20

	//回调函数 (感觉不像)
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})

	/*
		Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。
		这里主要列出go语言中的函数与java中的方法（函数）的一些不同
	*/
	/*
		1.语法不同,虽然不同，但其实很好记忆，就是他们之间相互写法是相反的，反着记就行。
		go:  func 函数名(参数)(返回值){
		函数体
		}
		java: 修饰符 返回值 函数名(参数){
		函数体
		}
		2.go有返回值的函数可以不用去接收返回值也可执行。
		3.go函数的参数中如果相邻变量的类型系统，则可以省略类型: func 函数名(i int , j int) ==> func 函数名(i, j, int)
		4.go函数中的返回值可以有多个: return x,m
		5.全局变量定义在函数外部的变量，有点js的意思，函数内部都可以访问到这个全局变量，若局部变量和全局变量重名， 则函数会先访问局部变量。

	*/
	//6.函数可以作为参数
	i2 := calculate(10, 20, 30, add11)
	fmt.Println(i2)
	//7.函数作为返回值
	fmt.Println(todoIt("+"))
	fmt.Println(todoIt("-"))
	fmt.Println(todoIt("*"))
	// 8.将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//9.自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	/*	练习题:
		你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
		分配规则如下：
		a. 名字中每包含1个'e'或'E'分1枚金币
		b. 名字中每包含1个'i'或'I'分2枚金币
		c. 名字中每包含1个'o'或'O'分3枚金币
		d: 名字中每包含1个'u'或'U'分4枚金币
		写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
	*/

	//自己写的，往下是目前看到的最优解
	for _, e := range users {
		_, ok := distribution[e]

		if !ok {
			distribution[e] = 0
		}

		a = strings.Split(e, "")
		for _, e1 := range a {

			space := strings.TrimSpace(e1)
			//1.传统if
			if strings.EqualFold(space, "e") {
				distribution[e]++
			}
			if strings.EqualFold(space, "i") {
				distribution[e] += 2
			}
			if strings.EqualFold(space, "o") {
				distribution[e] += 3
			}
			if strings.EqualFold(space, "u") {
				distribution[e] += 4
			}

		}
		if coins <= distribution[e] {
			fmt.Printf("用户%s分到%d金币，金币已分完", e, coins)
			coins = coins - coins
			break
		}
		coins -= distribution[e]
		fmt.Printf("用户%s分到%d金币，金币还剩下%d", e, distribution[e], coins)
		fmt.Println()
	}
	//最优解
	residue := dispatchCoin()
	fmt.Printf("剩余%d金币", residue)

}

func dispatchCoin() (residue int) {
	var (
		coins = 50
		users = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
		distribution = make(map[string]int, len(users))
	)
	for _, v := range users {
		total := 0
		for _, e := range v {
			switch e {
			case 'e', 'E':
				total += 1
			case 'i', 'I':
				total += 2
			case 'o', 'O':
				total += 3
			case 'u', 'U':
				total += 4
			}
		}

		if coins <= total {
			fmt.Printf("金币剩余够分的最后一个用户是%s，它分到%d金币，金币已分完", v, coins)
			distribution[v] = coins
			break
		} else {
			coins -= total
			distribution[v] = total
		}
	}
	fmt.Println(distribution)
	return coins
}

func todoIt(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add11, nil
	case "-":
		return subtract11, nil
	case "*":
		return multiply11, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func add11(x, y int) int {
	fmt.Println("10+10加法！")
	return 10
}
func subtract11(x, y int) int {
	fmt.Println("10-10加法！")
	return 0
}

func multiply11(x, y int) int {
	fmt.Println("10*10加法！")
	return 100
}

func calculate(x, y, z int, a func(int, int) int) int {

	return z + a(x, y)
}
