package main

import (
	"fmt"
	"unsafe"
)

/*
结构体 struct
Go语言中没有“类”的概念，也不支持“类“的继承等面向对象的概念。 Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的拓展性和灵活性
*/
/*
	自定义类型
		自定义类型是定义一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。
*/
type MyInt int //例如，定义自己的int类型,通过type关键字的定义，MyInt就是一种全新的类型，并且它具有int的特性。
/*
	go 1.9版本新功能，类型别名
		type TypeAlias = Type
		type byte = uint8
		type rune = int32
*/
//类型别名与类型定义的区别
type NewInt int
type AliasInt = int

/*
结构体定义
	通过type和struct定义
	type 类型名 struct {
		字段名 字段类型
		字段名 字段类型
		字段名 字段类型
		。。。
	}
*/
//定义了一个person的类型，语言内置的基础数据类型是用来描述一个值得，而结构体struct是用来描述一组值的。比如下面定义的person，表示了一个人有名字、年龄和居住城市等，本质上是一种聚合性的数据类型罢了，
type person struct {
	name, city string
	age        int8
}

type student1 struct {
	name string
	age  int
}

func main() {
	var n NewInt
	var a AliasInt
	fmt.Printf("NweInt type is %T", n) //main.NewInt
	fmt.Println()
	fmt.Printf("AliasInt type is %T", a) //还是 int，压根没变
	//结构体struct实例化
	//1.
	var p person
	p.name = "July"
	p.city = "ShenZhen"
	p.age = 18
	//2.
	p1 := person{"Jully", "NanNing", 18}
	//3.
	var p2 = person{"cindy", "ZhanJiang", 17}
	//4.
	var p3 person = struct {
		name, city string
		age        int8
	}{name: string("Jack"), city: string("Paris"), age: int8(16)}
	fmt.Println()
	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	//匿名结构体struct
	var user struct {
		name string
		age  int
	} //直接定义，一般在临时场景下使用
	user.name = "user"
	user.age = 18
	fmt.Println(user)
	//new 来实例化
	p5 := new(person)               //p5是一个指针类型
	fmt.Printf("p5 type is %T", p5) //p5 type is *main.person
	fmt.Println()
	fmt.Printf("p5 value is %#v", p5) //p5 value is &main.person{name:"", city:"", age:0}  %#v:相应值的Go语法表示
	fmt.Println()
	fmt.Printf("p5 value is %+v", p5) //p5 value is &{name: city: age:0} %+v:相应值的类型的Go语法表示
	//取结构体的地址实例化
	fmt.Println()
	p6 := &person{}       //使用&对结构体取地址操作相当于对该结构体struct类型进行了一次new实例化操作。
	fmt.Printf("%+v", p6) // 没有初始化的结构体，其成员变量都是对应其类型的零值：&{name: city: age:0}
	p6.name = "王一二"       //p6.name="王一二" 其实在底层是(*p6).name="王一二",这是Go语言帮我们实现的语法糖
	fmt.Println()
	fmt.Printf("%+v", p6)
	fmt.Println()
	m := make(map[string]*student1)
	stus := []student1{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
		fmt.Println(m)
	}

	for k, v := range m {

		fmt.Println(k, "=>", v.name)
		// 小王子 => 大王八；娜扎 => 大王八；大王八 => 大王八；值永远是大王八，其实可以猜测映射的值都是同一个地址，遍历到切片的最后一个元素“大王八”时，将“大王八”写入了该地址，所以导致映射所有值都相同。其实真实原因也是如此，因为for range创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，就会导致错误，在迭代时，返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以值的地址总是相同的，导致结果不如预期。
		//详见：https://studygolang.com/articles/9701
	}

	p7 := NewPerson("zhangsan", "jh", 18)
	fmt.Printf("%+v", p7)
	p7.setName("LiSi")
	fmt.Println()
	fmt.Println(p7) //&{LiSi jh 18}

	pt := PersonTest{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	pt.SetDreams(data)
	fmt.Println(pt.dreams) //[吃饭 睡觉 打豆豆]
	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(pt.dreams) // ? 如果SetDreams()中的赋值是直接赋值则结果为： [吃饭 不睡觉 打豆豆] ,用copy()赋值则不会受到影响

	fmt.Println()
	fmt.Println("/**************进阶篇*********/")
	//进阶篇(内存对齐)
	var t1 test1
	var t2 test2
	tt1 := &test1{}
	ttt1 := new(test1)
	tt2 := &test2{}
	fmt.Println(unsafe.Sizeof(t1))
	fmt.Println(unsafe.Sizeof(t2))
	fmt.Println(unsafe.Sizeof(tt1))
	fmt.Println(unsafe.Sizeof(tt2))
	fmt.Println(unsafe.Sizeof(ttt1))
}

/*
Go语言中没有构造函数，所以一般都是自己去写构造方法，返回值用指针类型接收（因为值返回的开销很大，所以用指针更合理）
*/
func NewPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

/**
方法和接受者(我个人更愿意称之为"主动绑定")
	定义格式如下:
	func (接受者变量 接受者类型) 方法名(参数)(返回参数){
		....
	}
	例子:
		首先要有一个结构体 type  User struct {}
		然后写一个函数func 然后把这个func绑定到这个结构体user中
		func (u User) doSomething(param1 string, param2 int) *User { ...逻辑代码 }  有点类似与Java中的User是一个class类，  然后doSomething 是这个User类中的一个方法，所以我更愿意理解为把doSomething()方法主动绑定到User结构体中，后续便可以通过User.doSomething()直接调用
	在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。
	注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
*/

func (p *person) setName(name string) {
	p.name = name
}

//Person 结构体Person类型，结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
type Person struct {
	string
	int
}

type PersonTest struct {
	name   string
	age    int8
	dreams []string
}

func (p *PersonTest) SetDreams(dreams []string) {
	//p.dreams = dreams  //这样赋值会有隐患， 当外部的dreams 值改变时，p.dreams也会跟着改变， 所以正确的做法是在方法中使用传入的slice的拷贝进行结构体的赋值
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

/*
进阶篇：内存对齐

*/
type test1 struct {
	a bool   // 1
	b int32  // 4
	c string // 16
}

type test2 struct {
	b int32  // 4
	c string // 16
	a bool   // 1

}
