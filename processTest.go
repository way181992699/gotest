package main

import "fmt"

/*/
程序中的流程控制，如if else for switch goto 等。
*/
type student struct {
	Name string
	Age  int8
}

func main() {
	/**
	if表达式的多种写法,注意if匹配的左括号{必须与if和表达式放在同一行，否者触发编译错误，else同理
	*/
	i := 10
	if i > 10 {
		fmt.Println("A")
	} else if i == 10 {
		fmt.Printf("及格")
	}
	//if特殊写法
	if score := 100.1; score > 100 {
		fmt.Printf("beautiful")
	}

	/**
	for的多种写法
	for 初始语句;条件表达式;结束语句{
	    循环体语句
	}
	*/
	//标准模式
	for i1 := 0; i1 < 100; i1++ {
		if i1 == 50 {
			fmt.Print("标准模式。")
		}

	}
	//特殊模式
	i2 := 10
	for ; i2 < 100; i2++ {
		if i2 == 50 {
			fmt.Print("特殊模式。")
		}

	}
	//类似while模式
	for i2 < 400 {
		fmt.Print(i2)
		i2++
	}
	//无线循环模式，可通过return、goto、break、panic语句强制退出。
	for {
		fmt.Print("无线循环。")
		if i2 > 300 {
			break
		}
	}
	/**
	for range （键值循环）
	Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：
	数组、切片、字符串返回索引和值。
	map返回键和值。
	通道（channel）只返回通道内的值。
	*/
	//数组
	_array01 := [5]int{11, 22, 33, 44, 55}
	for _, e := range _array01 {
		fmt.Println(e)
	}
	//切片
	_array02 := []string{"a", "b", "c"}
	for _, e := range _array02 {
		fmt.Println(e)
	}
	//map
	_map := map[int8]string{1: "赵六", 2: "孙七", 3: "周八"}
	for k, v := range _map {
		//k为map的key值
		fmt.Println(k)
		//v为map的value
		fmt.Println(v)
	}
	/*
		若只想遍历key,写法
		for k:=range _map{
		}
		注意： 遍历map时的元素顺序与添加键值对的顺序无关。
	*/

	/*
		switch
	*/
	//1.标准用法
	expr := 100
	switch expr {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	default:
		fmt.Println("NO MATCH")
	}
	//1.1还可以简写成
	switch expr := 100; expr {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	default:
		fmt.Println("NO MATCH")
	}
	//2.case可以有多个值，用英文逗号分开
	switch expr {
	case 1, 60, 70:
		fmt.Println("A")
	case 10, 90, 100:
		fmt.Println("B")
	default:
		fmt.Println("NO MATCH")
	}
	//3.case还可以使用表达式,这时候switch语句后面不能再跟判断变量
	switch {
	case expr > 60:
		fmt.Println("A")
	case expr <= 60:
		fmt.Println("B")
	default:
		fmt.Println("NO MATCH")
	}
	//4.还可以添加fallthrough语句，可以执行满足条件的case的下一个case(这下一个case无论满不满足条件都会执行)，为了兼容C语言中的case设计的。
	switch expr := 100; expr {
	case 1, 60, 70:
		fmt.Println("A4")
	case 10, 90, 100:
		fmt.Println("B4")
		fallthrough
	case 20, 96, 107:
		fmt.Println("C4")
	default:
		fmt.Println("NO MATCH")
	}

	/**
	goto
	goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用goto语句能简化一些代码的实现过程。
	*/
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i+j == 80 {
				fmt.Printf("i=%v;j=%v", i, j)
				goto breakFun
			}

		}
	}
breakFun:
	fmt.Println("代码结束")
	/*
		break和continue语句个java中的break、continue差不多，不同的是，go中的break和continue可以在后面添加标签，然后跳到标签对应的循环中。
	*/
forLoop1:
	for i := 0; i < 5; i++ {
		// forLoop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				//这是举例的是continue，其实break也是一样的写法
				continue forLoop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

	/**
	  练习题，打印99乘法表
	*/

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%vx%v=%v ", j, i, i*j)
		}
		fmt.Println()
	}

}
