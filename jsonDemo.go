package main

import (
	"encoding/json"
	"fmt"
)

type StudentOne struct {
	Score int8
	Name  string
}

type Class struct {
	Name       string
	StudentOne []*StudentOne
}

func main() {

	c1 := &Class{
		Name:       "高中247班",
		StudentOne: make([]*StudentOne, 0, 100),
	}
	for i := 0; i < 10; i++ {
		s := &StudentOne{
			Score: int8(60 + i),
			Name:  fmt.Sprintf("July%v", i),
		}
		c1.StudentOne = append(c1.StudentOne, s)
	}
	fmt.Println()
	//fmt.Printf("_c1 == %+v", c1)
	data, err := json.Marshal(c1) //踩坑，结构体成员变量小写则无法解析成json，https://zhuanlan.zhihu.com/p/66941010原因：go用首字母的大小写来确定是共有的还是私有的，也就是一个变量函数等能不能被其他包引用，小写字母开头的（私有）只能包内使用，不能被其他包使用。 因为json.Marshal是另外一个包，json这个包没法给你现在所在的包里的任何私有变量(小写字母开头的)赋值。
	if err != nil {
		fmt.Println("json 解析错误")
		return
	}
	fmt.Println()
	fmt.Printf("json:%s\n", data)

}
