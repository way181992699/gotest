package main

import (
	"fmt"
	"time"
)

/*
Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
channel是一种类型，一种引用类型。声明通道类型的格式如下：
var 变量 chan 元素类型
举几个例子：
var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道
*/

func recv(c chan int) {
	ret := <-c
	fmt.Printf("接收成功%v", ret)
	fmt.Println()
}
func send(c chan int) {
	c <- 20
	fmt.Printf("发送成功%v", 20)
	fmt.Println()
}

func main() {
	/*
	   无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。

	   使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。
	*/
	ch := make(chan int)
	go send(ch)
	go recv(ch)
	//ch <- 10
	//fmt.Println("发送成功")

	time.Sleep(time.Second)

	//time.Sleep(time.Second)

	ch1 := make(chan int, 4)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)
		case ch1 <- i:
			fmt.Println(i)
		}
	}

}
