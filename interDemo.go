package main

import (
	"fmt"
	"os"
)

type People1 interface {
	Speak(string) string
	run()
	jump()
}

type People2 interface {
	walk()
	plLog(path, logMsg string)
}

type studentOne struct{}

func (stu *studentOne) run() {

	panic("implement me")
}

func (stu *studentOne) jump() {
	panic("implement me")
}

type staff struct{}

func (s *staff) plLog(path, log string) {
	//    O_RDONLY  文件以只读模式打开
	//    O_WRONLY  文件以只写模式打开
	//    O_RDWR   文件以读写模式打开
	//    O_APPEND 追加写入
	//    O_CREATE 文件不存在时创建
	//    O_EXCL   和 O_CREATE 配合使用,创建的文件必须不存在
	//    O_SYNC   开启同步 I/O
	//    O_TRUNC  打开时截断常规可写文件
	file, e := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	defer file.Close()
	if e != nil {
		create, e := os.Create(path)
		if e != nil {
			fmt.Printf("err=%v", e)
			panic("无法创建日志文件")
		} else {
			create.Write([]byte(log))
			fmt.Printf("log=%v", log)
		}
	} else {
		file.Write([]byte(log))
		fmt.Printf("log=%v", log)
	}

}

func (s *staff) walk() {

	fmt.Println("散步散步")
}

func (s *staff) run() {

}

func (s *staff) jump() {

}

func (stu *studentOne) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func (s *staff) Speak(t string) (talk string) {
	return "打工人，打工魂！"
}

// 使用空接口实现可以接收任意类型的函数参数。
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
	//类型断言，也可以用于接口类型强转
	t, ok := a.(People1)
	if ok {
		fmt.Println(t.Speak("sb"))
	}
}

func main() {
	var peo People1 = &studentOne{}
	var st People2 = &staff{}
	think := "sb"
	fmt.Println(peo.Speak(think))
	st.walk()
	show("1")
	show(1)
	show(st)
	show(peo)
	log := "error info "
	path := "D:/Code/Go/Test/static/log.txt"
	st.plLog(path, log)
}
