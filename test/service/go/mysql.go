
package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("open the database, sio")

	db, err := sql.Open("mysql", "sio:sio@tcp(127.0.0.1:3306)/sio")
	checkErr(err)

	//insert
	fmt.Println("insert data info into table, student")
	stmt, err := db.Prepare("INSERT student (ID, name, gender) value (?, ?, ?)")
	checkErr(err)
	res, err := stmt.Exec(5, "go", "1")
	checkErr(err)

	fmt.Println(res.LastInsertId())

	//update
	fmt.Println("update data info")
	stmt, err = db.Prepare("update student set name = ? where ID = ?")
	checkErr(err)
	res, err = stmt.Exec("gogogo", 5)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//retrive
	fmt.Println("query data into from table, student")
	rows, err := db.Query("select * from student")
	checkErr(err)
	fmt.Println("iterator data info to show")
	for rows.Next() {
		var ID int
		var name string
		var gender string
		err = rows.Scan(&ID, &name, &gender)
		checkErr(err)
		fmt.Println("ID:", ID, "name:", name, "gender:", gender)
	}

	//delete
	stmt, err = db.Prepare("delete from student where ID = ?")
	checkErr(err)
	res, err = stmt.Exec(5)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()
}
