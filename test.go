package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	fmt.Println("emmmm....")

	total := 0

	for i := 0; i <= 10; i++ {
		total += i
	}

	fmt.Println(total)

	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	for {
		sum++

		if sum > 1000 {
			break
		}
	}

	fmt.Println(sum)
	nums := []int{1, 2, 5, 4, 6}

	//当不需要index时，可以用空白符"_"代替
	for _, e := range nums {
		fmt.Printf("%d ", e)

	}
	//同理
	for index, _ := range nums {
		fmt.Print(index)
	}
	fmt.Println()
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, e := range kvs {
		fmt.Printf("%s -> %s\n", k, e)
	}

	var mapTest map[string]string
	mapTest = make(map[string]string)
	mapTest["a"] = "A"
	mapTest["b"] = "B"
	mapTest["c"] = "C"

	for k, v := range mapTest {
		fmt.Printf("%s == %s", k, v)
	}

	capital, s := mapTest["c"]
	if s {
		fmt.Println("1", capital)
	} else {
		fmt.Println("1")
	}
	//或者毫秒级别的时间戳
	fmt.Println(time.Now().UnixNano() / 1e6)
	versionStr := "2.0.2"

	versionArray := versionStr[:5]
	fmt.Println(versionArray)

	e := "E "
	fmt.Println(strings.EqualFold(e, "e"))

	//尝试直接遍历一个字符串,发现居然可以遍历字符串，而且遍历结果是字符，而且类型是int32，
	str := "Matthew12中"
	for _, v := range str {
		fmt.Printf("数据类型为%T,值为%d", v, v)
		fmt.Println()
		if v == 'a' {
			fmt.Printf("true")
		}
		if v == '中' {

			fmt.Printf("汉字")
		}
		//switch v {
		//case 'a':
		//	fmt.Printf("----")
		//}
	}

}
