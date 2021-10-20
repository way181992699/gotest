package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
在 Go http包的Server中，每一个请求在都有一个对应的 goroutine 去处理。请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和RPC服务。用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，
比如终端用户的身份认证信息、验证相关的token、请求的截止时间。 当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。
 */

var wg01 sync.WaitGroup //https://studygolang.com/articles/12972?fr=sidebar（WaitGroup的具体用法详见链接）
var i  int
var exit bool
func worker() {
	for {
		fmt.Println("worker"+strconv.Itoa(i))
		i++
		time.Sleep(time.Second)
		if exit {
			break
		}
	}

	wg01.Done()
}


func main() {

	wg01.Add(1)
	go worker()
	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exit = true                 // 修改全局变量实现子goroutine的退出
	// 如何优雅的实现结束子goroutine
	wg01.Wait()
	fmt.Println("over")

}


