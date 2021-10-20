package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//类似java的字节输入流
func readFile() ([]byte, error) {
	file, err := os.Open("D:/July/temporary/windows.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return nil, err
	}
	defer file.Close()
	var content []byte
	var tmp = make([]byte, 256)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return nil, err
		}
		content = append(content, tmp[:n]...)
	}
	return content, nil
}

//类似java的字符输入流
func bufIoTest() {
	file, err := os.Open("D:/July/temporary/windows.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//var content string
	for {

		character, err := reader.ReadString('\n') //'\n'读取字符的固定写法
		if err == io.EOF {
			if len(character) != 0 {
				fmt.Println(character)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(character)
	}
	//fmt.Println(content)
}

//文件io工具类
func ioUtilTest() ([]byte, error) {
	bytes, e := ioutil.ReadFile("D:/July/temporary/windows.txt")
	if e != nil {
		return nil, e
	}
	return bytes, nil
}

//字节输出流
func writerTest(b []byte, filePath string) error {
	file, e := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	if e != nil {
		return e
	}
	defer file.Close()
	file.Write(b)

	//测试直接插入字符串看是否能插入且是否直接在后面插入
	str := "\n" + "帅哥 July"
	file.WriteString(str)
	return nil
}

//ioUtilsWriter
func ioUtilsWriter(b []byte, filePath string) error {
	return ioutil.WriteFile(filePath, b, 0666)
}

//main方法
func main() {
	//读取字节流
	//fmt.Println("字节流读取文件")
	//readFile()
	//fmt.Println("字符流读取文件")
	//bufIoTest()
	//ioutis工具读取文件
	bytes, e := ioUtilTest()
	if e != nil {
		fmt.Println(e)
	}

	err := writerTest(bytes, "D:/July/temporary/windowsCopy.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("写入成功！")
	}

	ioUtilsWriter(bytes, "D:/July/temporary/windowsIoUtilsWriter.txt")

}
