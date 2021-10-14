package main

import (
	"fmt"
	"sort"
)

func main() {

	/**
	数组声明
	*/
	//1.1 标准声明
	var testArray01 [5]int
	var testArray02 = [2]int{1, 54}
	var testArray03 = [2]string{"A", "B"}
	fmt.Println(testArray01)
	fmt.Println(testArray02)
	fmt.Println(testArray03)
	//1.2 自动判断
	var testArray05 = [...]int{1, 54}
	var testArray06 = [...]string{"A", "B"}
	fmt.Println(testArray05)
	fmt.Println(testArray06)
	//1.3 制定索引值模式
	var testArray07 = [...]int{1: 1, 4: 99} //[0 1 0 0 99]
	fmt.Println(testArray07)
	//注意数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

	//指针
	//var i int = 20
	//var ip *int
	////fmt.Println(&i)
	////fmt.Println(*&i)
	//ip = &i
	////fmt.Println(ip)
	////fmt.Println(*ip)
	//
	////切片（可变数组）
	//var list []int            //定义
	//list1 := []int{1, 4, 7}   //定义
	//ints := append(list, 1)   //追加
	//list2 := append(list1, 1) //追加
	//fmt.Println(ints)
	//fmt.Println(list2)
	//追加
	//x := []int{1, 2, 3}
	//y := []int{4, 5, 6}
	//x = append(x, y...)
	//fmt.Println(x)
	//复制
	//x1 := []int{1, 2, 3}
	//y1 := append([]int(nil), x1...)
	//fmt.Println(y1)
	//x1 = x1[1:2]
	//fmt.Println(x1)
	//fmt.Println(y1)
	//删除某一个
	x3 := []int{1, 2, 3, 7, 8}
	x3 = append(x3[:2], x3[2+1:]...)
	//	fmt.Println(x3)
	//删除部分
	x4 := []int{1, 2, 3, 7, 8}
	x4 = append(x4[:2], x4[4:]...)
	//fmt.Println(x4)
	//实现栈和队列
	i5 := 9
	x5 := []int{1, 2, 3, 7, 8}
	x5 = append(x5, i5)
	fmt.Println(x5)
	//栈
	x5 = x5[:len(x5)-1]
	fmt.Println(x5)
	//队列
	x5 = x5[1:]
	fmt.Println(x5)

	var list []int

	for i := 0; i < 10; i++ {
		//类似java的list.add(),但原理完全不一样
		list = append(list, i)
	}
	fmt.Println(list)

	//切片有一点需要注意的是， 因为切片属于引用类型， 所以当A和B都指向同一个地址时， 修改B的值也会修改到A的值。 所以衍生出了一个copy()的概念。
	a := []int{1, 2, 3, 4, 5}
	b := a
	b[1] = 6
	fmt.Println(a) //[1 6 3 4 5]
	//copy()函数:copy(destSlice, srcSlice []T) -> srcSlice: 数据来源切片; destSlice: 目标切片 ; 函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：
	c := make([]int, 5, 10)
	copy(c, a)
	fmt.Println(c)
	c[2] = 7
	fmt.Println(a) //[1 6 3 4 5]
	fmt.Println(c) //[1 6 7 4 5]
	//切片中的删除功能， 因为切片并没有删除的专用方法， 所以我们可以利用切片的特性（切片简单表达式 a[:]），去删除元素,以上面的c切片为例，想要删除C切片的索引为2的元素,就是把7删掉。
	c = append(c[:2], c[3:]...) //之所以用...是因为append方法的参数决定的 func append(slice []Type, elems ...Type) []Type
	fmt.Println(c)              //[1 6 4 5] 总结来说就是，当想删除掉切片中索引为index的元素，操作方法是:a=append(a[:index],a[index+1:]...)

	/**
	练习
	*/
	//练习1给出打印结果
	aa := make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		aa = append(aa, fmt.Sprintf("%v", i))
	}
	fmt.Println(aa)
	fmt.Println(len(aa))
	fmt.Println(cap(aa))
	//练习2，对某个已知数组进行排序
	//Ints()方法会对数值类型进行升序排序
	sort.Ints(c)
	//如果想要自定义排序，则需要去实现sort的三个接口 Len()--返回长度 ,Less()--具体排序比较 ,Swap()--相邻元素调换位置
	fmt.Println(c)
	as := ArraySort{}
	as = []int{1, 56, 7, 46, 788, 9}
	sort.Sort(as)
	fmt.Println(as)
	sort.Ints(as)
	fmt.Println(as)
}

type ArraySort []int

func (as ArraySort) Len() int {
	return len(as)
}

func (as ArraySort) Less(i, j int) bool {
	return as[i] > as[j]
}

func (as ArraySort) Swap(i, j int) {
	as[i], as[j] = as[j], as[i]
}
