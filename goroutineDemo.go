package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
在java/c++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量的心智。那么能不能有一种机制，
程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？
Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，
就是因为它在语言层面已经内置了调度和上下文切换的机制。
在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。
*/
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine！", i)
}

func _a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%v", i)
		fmt.Println()
	}

}

func _b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%v", i)
		fmt.Println()
	}

}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束

	//nano := time.Now().UnixNano() / 1e6

	runtime.GOMAXPROCS(2) //当逻辑核心数设置为2时，此时两个任务并行执行。
	go _a()               //在函数前面加一个"go"关键字， 就会开启一个goroutine 那就可以进行并发进行了
	go _b()
	//_a()//常规的串行
	//_b()//常规的串行

	//fmt.Println()
	//unixNano := time.Now().UnixNano() / 1e6
	//fmt.Printf("time=%v", unixNano-nano)
	//fmt.Println()
	//1376
	time.Sleep(time.Second)

}
