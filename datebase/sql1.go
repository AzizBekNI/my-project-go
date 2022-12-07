package datebase

import (
	"database/sql"
	"fmt"
)

func Sql1() {
	db, err := sql.Open("mysql", "root:Aa12345$@/testDB")
	if err != nil {
		panic(err.Error())
	}
	var (
		bookName string
		pressYear int 
		authore string
	)
	fmt.Println("Book Name : ")
	fmt.Scan(&bookName)
	fmt.Println("Press Year : ")
	fmt.Scan(&pressYear)
	fmt.Println("Authore : ")
	fmt.Scan(&authore)
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS kutbxona(id int NOT NULL AUTO_INCREMENT, bookName varchar(45), pressYear int, authore varchar(55), PRIMARY KEY(id), UNIQUE(id));")
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("INSERT INTO kutbxona(bookName, pressYear, authore) VALUES(?,?,?);", bookName, pressYear, authore)
	if err != nil {
		panic(err.Error())
	}
}
