package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"

	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "zouhl:z@/test?charset=utf8")
	checkErr(err)

	//insert sql
	fmt.Println("insert sql")
	stmt, err := db.Prepare("insert userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("zouhl", "研发部门", "2018-09-25")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//update data
	fmt.Println("update sql")
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("zouhl_update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("affect:", affect)

	//query data
	fmt.Println("query sql")
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
		fmt.Println("###########")
	}

	// delete data
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("affect", affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
