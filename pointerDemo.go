package main

import "fmt"

/*
指针
指针地址、指针类型、指针取值
*/

func main() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	//指针传值经典案例
	x := 1
	fun(x)         //fun()把x改为100
	fmt.Println(x) //x 还是为10
	fun1(&x)       //fun1()通过*&x改为100
	fmt.Println(x) //x =100了， 因为直接通过指针把x的真正值给改了。

	/*
		new：
			函数签名： func new(Type) *Type
			Type表示类型，new函数只接受一个参数，这个参数是一个类型
			*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
			var i * int //只是声明了一个指针变量a，但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。
			*i = 100    //直接运行会报panic异常，因为*int 属于因为引用类型， 所以我们在使用的时候不仅要声音它，还要为它分配内存空间，否则我们的值就没办法储存。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
			应该写为：
			var i *int
			i = new(int)
			*i =  70
			fmt.Println(*i)
	*/
	i2 := new(int)
	fmt.Printf("i2 Type is %T ,value is %d", i2, *i2)
	fmt.Println()
	/*
		make:
			make也是用于分配内存，却别于new，它只用做对slice（切片）、map（集合）、chan（线程通道）的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
			make函数签名： func make(t Type,size ...IntegerType) Type
			make函数是无可替代的，因为我们在使用slice、map、channel的时候，都需要使用make函数进行初始化，然后才可以对它们进行操作。
	*/
	//简单示例：
	var _map map[string]string
	_map = make(map[string]string, 10)
	_map["hi"] = "嗨"
	fmt.Println(_map)
	/*
		new和make的区别
		1.都是用来分配内存空间
		2.make只针对slice、map、chan三个类型的内存分配，而且返回的是这三个类型本身。
		3.new是对类型进行内存分配，返回的是指针类型。
	*/
}

func fun(x int) {
	x = 100
}
func fun1(x *int) {
	*x = 100
}
