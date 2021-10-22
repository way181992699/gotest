package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

/*
结合net/http和database/sql实现一个使用MySQL存储用户信息的注册及登陆的简易web程序。
*/

//全局
var db1 *sql.DB

//登录类
type login struct {
	id                        int
	userName, password, token string
}

//register

func registerFunc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	query := r.URL.Query()
	userName := query.Get("userName")
	pwd := query.Get("password")
	loginType, _ := strconv.Atoi(query.Get("loginType"))
	var count int

	sql := "Select count(*) from t_login where userName =?"
	err := db1.QueryRow(sql, userName).Scan(&count)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	//若用户存在
	if count > 0 {
		//loginType ==1 执行注册逻辑
		if loginType == 0 {
			warn := "用户名已存在，请重新输入"
			w.Write([]byte(warn))
			return
		}
		//loginType ==1 执行登录逻辑
		if loginType == 1 {
			//获取用户详细信息
			var l login
			sql := "Select userName,password from t_login where userName =?"
			err := db1.QueryRow(sql, userName).Scan(&l.userName, &l.password)
			if err != nil {
				fmt.Printf("query userinfo failed, err:%v\n", err)
				return
			}
			if l.password != pwd {
				warn := "密码不正确，请重新输入"
				w.Write([]byte(warn))
				return
			} else {
				w.Write([]byte("login success !"))
				return
			}
		}
	} else { //用户不存在
		//注册
		if loginType == 0 {
			sql := "Insert into   t_login(userName,password) VALUES(?,?)"
			result, err := db1.Exec(sql, userName, pwd)
			if err != nil {
				fmt.Printf("Insert into failed, err:%v\n", err)
				w.Write([]byte("error"))
				return
			}
			theID, err := result.LastInsertId() // 新插入数据的id
			if err != nil {
				fmt.Printf("get lastinsert ID failed, err:%v\n", err)
				w.Write([]byte("error"))
				return
			}
			sprint := fmt.Sprintf("register success! userId: %d", theID)
			w.Write([]byte(sprint))
			return
		}
		if loginType == 1 {
			warn := "用户名不存在，请先注册"
			w.Write([]byte(warn))
			return
		}
	}

}

// initDB

func initDB1() {
	var err error
	db1, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Printf("mysql open err , err : %v \n", err)
		return
	}
	//尝试连接
	err = db1.Ping()
	if err != nil {
		fmt.Printf("mysql connect failed , err : %v \n", err)
		return
	}

}

func main() {
	////初始化DB
	initDB1()
	defer db1.Close()
	//建立http服务端
	http.HandleFunc("/register", registerFunc)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http listen failed , err %v \n", err)
		return
	}

}
