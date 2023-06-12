package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id         int64
	username   string
	departName string
	createAt   string
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	checkErr(err)

	// 插入1万条数据
	insertManyUsers(db, 1000)

	// 查询数据
	queryUserOne(db)

	db.Close()
}

// 插入多个用户数据
func insertManyUsers(db *sql.DB, count int) {
	stmt, err := db.Prepare("INSERT user SET user_name=?, depart_name=?, create_at=?")
	checkErr(err)

	for i := 0; i < count; i++ {
		username := fmt.Sprintf("用户%d", i+1)
		departName := fmt.Sprintf("部门%d", i+1)
		createAt := "2000-10-10" // 假设都使用相同的创建日期

		_, err := stmt.Exec(username, departName, createAt)
		checkErr(err)
	}

	fmt.Printf("成功插入%d条数据\n", count)
}

// 查询用户
func queryUserOne(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.id, &user.username, &user.departName, &user.createAt)
		checkErr(err)
		fmt.Println("查询到的用户数据", user.id, user.username, user.departName, user.createAt)
	}
}

// 检查错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
