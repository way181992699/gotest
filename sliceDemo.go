package main

import "fmt"

/**
关于切片的一些进阶技巧
*/
var sliceA = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {

	//1.copy 切片,把sliceA复制到b中，可以用go内置的copy（）函数
	b := make([]int, len(sliceA))
	copy(b, sliceA)
	fmt.Println(b)
	//2.将sliceA的3~6位的元素减掉
	b = append(b[:3], b[6:]...)
	fmt.Println(b)
	fmt.Println(sliceA)

}
