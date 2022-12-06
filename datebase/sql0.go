package datebase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// var Num int = 123
func Sql1() {
	db, err := sql.Open("mysql", "root:Aa12345$@/testDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//	db.Exec("CREATE DATABASE testDB")
	/*
		creat table is not exists 'users'(
			'id' int(100) not null auto_increment,
			'username' varchar(45) not null,
			'password' varchar(40) not null,
			'isActive' tinyint(1) default null,
			primary key ('id'),
			unique key 'id_unique' ('id')
		)
	*/
	query := "CREATE TABLE example ( id integer, data varchar(32) )"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("blah blah ")
}
