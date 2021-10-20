package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Go语言内置的net/http包十分的优秀，提供了HTTP客户端和服务端的实现。
*/

type cindy struct {
	Name string
	Age  int
}

func httpDemo() {
	resp, err := http.Get("https://www.liwenzhou.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

func getHttpByte(addr string) ([]byte, error) {
	resp, err := http.Get(addr)
	if err != nil {
		fmt.Printf("http access failed , err: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("io read failed , err: %v\n", err)
		return nil, err
	}
	return bytes, nil
}

//服务端处理Get请求HandlerFun函数
func getParamsFun(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	query := r.URL.Query()
	name := query.Get("name")
	fmt.Printf("name is %T ;name == %v \n", name, name)
	age := query.Get("age")
	fmt.Printf("age is %T ;age = %v \n", age, age)
	s := `{"status":"ok"}`
	w.Write([]byte(s))
}
func postParamsFun(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	//1.如果请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	//2.请求类型是application.json时,直接从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	c := &cindy{}
	json.Unmarshal(b, c)
	fmt.Println(string(b))
	s := `{"status":"data":${c}}`
	w.Write([]byte(s))
	fmt.Printf("%+v\n", c)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v", c.Age+1)

}

func main() {
	//httpDemo()
	bytes, _ := getHttpByte("https://www.liwenzhou.com")

	normalize, _ := getHttpByte("https://www.liwenzhou.com/css/normalize.css")

	style, _ := getHttpByte("https://www.liwenzhou.com/css/style.css")

	sidebar, _ := getHttpByte("https://www.liwenzhou.com/css/sidebar_menu.css")

	ioutil.WriteFile("D://liwenzhou.html", bytes, 0666)
	ioutil.WriteFile("D:/css/normalize.css", normalize, 0666)
	ioutil.WriteFile("D:/css/style.css", style, 0666)
	ioutil.WriteFile("D:/css/sidebar_menu.css", sidebar, 0666)

	http.HandleFunc("/get", getParamsFun)
	http.HandleFunc("/post", postParamsFun)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http access failed , err:%v", err)
		return
	}
}
