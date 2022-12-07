package datebase

import (
	"database/sql"
	"fmt"
	//"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// var Num int = 123
func Sql0() {
	db, err := sql.Open("mysql", "root:Aa12345$@/testDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//	db.Exec("CREATE DATABASE testDB")
	/*
		creat table is not exists users(
			id integer auto_increment,
			username varchar(45),
			password varchar(40),
			isActive tinyint(1),
			primary key (id),
			unique key id_unique (id)
		)
	*/
	/*CREATE TABLE Persons (
		ID int NOT NULL,
		LastName varchar(255) NOT NULL,
		FirstName varchar(255),
		Age int,
		UNIQUE (ID)
	);*/
	// query := "CREATE TABLE if not exists users(id int NOT NULL AUTO_INCREMENT,username varchar(45),password varchar(40),isActive tinyint(1),PRIMARY KEY (id),unique (id));"
	// _, err = db.Exec(query)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// var (
	// 	name string
	// 	password string
	// 	isActive int
	// )
	// fmt.Println("Type name : ")
	// fmt.Scan(&name)
	// fmt.Println("Password : ")
	// fmt.Scan(&password)
	// fmt.Println("isActive : ")
	// fmt.Scan(&isActive)

	// res, err := db.Exec("INSERT INTO users(username, password, isActive) VALUES(?,?, ?);", name, password, isActive)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	//fmt.Println(res.LastInsertId())
	//fmt.Println(res.RowsAffected())
	// id, err := res.LastInsertId()
    // if err != nil {
    //     panic (err.Error())
    // }
    // fmt.Println("id: ", id)
	// _, err = db.Exec("UPDATE users SET password = 3456 WHERE username = 'Farrux';")
	// if err != nil{
	// 	log.Panic(err.Error())
	// }
	// _, err = db.Exec("DELETE FROM users WHERE username = 'Alloshukr';")
	// if err != nil{
	// 	log.Panic(err.Error())
	// }
	fmt.Printf("1 User qo'shish \n 2 ")
}
