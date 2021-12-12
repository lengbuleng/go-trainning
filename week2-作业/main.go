package main

import (
	"database/sql"
	"fmt"

	xerrors "github.com/pkg/errors"
)

// 用户模型, 表结构, 需要一个结构来接收查询结果集
type User struct {
	Id   int32
	Name string
	Age  int8
}

func DaoWrapErr(db *sql.DB, user *User) error {

	// Query结果集需要调用Next()方法进行逐条遍历
	rows, err := db.Query(`
        SELECT id,name,age FROM user
    `)
	if err != nil {
		return xerrors.Wrap(err, "Query failed.")
	}
	defer rows.Close()

	for rows.Next() {
		// 对于遍历, 只需要判断每次是否有错误产生即可
		// 参数绑定需要数量和位置一一对应
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return xerrors.Wrap(err, "Get rows failed.")
		}
	}
	return nil
}

func main() {
	// mysql扩展 https://github.com/go-sql-driver/mysql
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// 保存用户信息列表
	var user User

	err = DaoWrapErr(db, &user)
	//判断是否是sql.ErrNoRows
	if xerrors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("user:%v", user)
}
