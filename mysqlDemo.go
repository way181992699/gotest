package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
操作Mysql
*/

var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

// 定义一个初始化数据库的函数
func initDB(dsn string) (err error) {

	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

/*
查询单条数据
*/
func queryOne(sqlStr string) {
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

/*
查询多条条数据
*/
func queryMore(sqlStr string) {

	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

/***************插入、更新和删除操作都使用Exec方法。**************************/

/*
insert into
*/
func insertInto(sqlStr string, args ...interface{}) {
	ret, err := db.Exec(sqlStr, args[0], args[1])
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

/*
update
*/
// 更新数据
func updateRowDemo(sqlStr string, args ...interface{}) {
	ret, err := db.Exec(sqlStr, args[0], args[1], args[2])
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

/*
delete
*/
func deleteRow(sqlStr string, args interface{}) {

	ret, err := db.Exec(sqlStr, args)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func main() {
	//defer db.Close()
	// DSN:Data Source Name
	err := initDB("root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	defer db.Close()
	//Insert func
	insertInto("insert into user(name,age) values (?,?)", "张三", 24)

	//Update func
	updateRowDemo("update user set name=?,age=? where id = ?", "cindy", 17, 2)

	//Delete func
	deleteRow("delete from user where id = ?", 8)

	//因为插入、更新和删除操作都使用Exec方法，所以如果封装得好的话，可以用updateRowDemo函数也能达到一样的效果
	//Query func
	//queryOne("select * from user where id =?")
	//queryMore("select * from user")
}
