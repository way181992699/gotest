package main

import (
	"fmt"
)

func main() {

	var i int32
	var f32 float32 = 132
	var f64 float64 = 132
	var b bool = true
	var str string = "hello七月小王子"

	fmt.Printf("i=%d ", i)
	fmt.Printf("i=%f ", f64)
	fmt.Printf("i=%.3f ", f32)
	fmt.Printf("i=%v ", b)
	fmt.Printf("i=%s ", str)
	fmt.Println()
	//求str包含几个汉字
	for _, v := range str {
		bytes := []byte(string(v))
		if len(bytes) > 1 {
			fmt.Println(string(bytes))
			i++
		}
	}
	fmt.Printf("汉字的个数为:%d个", i)

}
