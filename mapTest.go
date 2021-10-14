package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
map集合
*/
func main() {

	/*
		map[KeyType]ValueType
		KeyType:表示键的类型。
		ValueType:表示键对应的值的类型。
		map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
		make(map[KeyType]ValueType, [cap])
		其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
	*/

	//1.先声明后赋值
	_map := make(map[int8]string, 8)
	_map[1] = "张三"
	_map[2] = "李四"
	_map[3] = "王五"
	fmt.Println(_map)
	//2.声明直接赋值,因为直接赋值不是nil，所以不用make()
	_map02 := map[int8]string{1: "赵六", 2: "孙七", 3: "周八"}
	fmt.Println(_map02)
	//3.判断某个键是否存在的表达式value, ok := map[key]
	v, ok := _map[5]
	fmt.Println(v)
	fmt.Println(ok)
	//4.遍历Map
	for k, v := range _map {
		fmt.Println(k)
		fmt.Println(v)
	}
	//5.删除map中的键值对delete(map,key)函数
	fmt.Printf("原先的:%v", _map)
	delete(_map, 2)
	fmt.Println()
	fmt.Printf("删除键值为2后的:%v", _map)
	//6.注意在遍历map时，元素的顺序与添加的顺序无关，所以经常会做排序处理
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)
	/*
		练习
	*/
	//1.写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1
	str := "how do you do how you do how do you how do do you do"
	split := strings.Split(str, " ")
	var _mapStr = make(map[string]int8)
	for _, v := range split {

		if _, ok := _mapStr[v]; !ok {
			_mapStr[v] = 1

		} else {
			_mapStr[v]++
		}
	}
	fmt.Println(_mapStr)
	//2.观察下面代码，写出最终的打印结果。
	funny()

}

func funny() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	fmt.Printf("%+v\n", m)
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
