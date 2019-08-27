package main

import (
	"database/sql" // 这是一个抽象层包，比如区分mysql、orcal等数据库，只有这个包是连接不上mysql的，还需要搭配下面的mysql包
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //导入mysql驱动包
)

type Student struct {
	Id   int    `db:id`
	Name string `db:name`
}

func HandleErr(when string, err error) {
	if err != nil {
		fmt.Println(when, ":", err)
		os.Exit(1)
	}
}

var s []Student

func main() {
	db, err := sql.Open("mysql", "root:1234qwer@tcp(127.0.0.1:3306)/gouse?charset=utf8")
	HandleErr("打开数据库", err)
	defer db.Close()

	err = db.Ping()
	HandleErr("ping数据库", err)

	// _, err = db.Exec("INSERT INTO student(id,name) values(4,?)", "Dave")
	// HandleErr("INSERT err:", err)

	// rows, err := db.Query("SELECT * FROM student")
	// HandleErr("Query err:", err)
	// defer rows.Close()

	// count := 0
	// for rows.Next() {
	// 	err := rows.Scan(&s[count].Id, &s[count].Name)
	// 	HandleErr("rows scan", err)
	// 	fmt.Println(s[count])
	// 	count++
	// }
	
	// rows, err := db.Query("SELECT * FROM student")
	// HandleErr("Query err:", err)
	// defer rows.Close()

}
