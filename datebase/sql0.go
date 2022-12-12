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

	//fmt.Println(res.LastInsertId())
	//fmt.Println(res.RowsAffected())
	// id, err := res.LastInsertId()
	// if err != nil {
	//     panic (err.Error())
	// }
	// fmt.Println("id: ", id)

	var a int
	c := true
	for c {
		fmt.Printf("1 User qo'shish \n2 User ma'lumotini yangilash \n3 User ning ma'lumotlarini o'chirish \n")
		fmt.Scanf("%d", &a)
		if a > 0 && a < 5 {
			c = false
		}
	}
	switch a {
	case 1:
		var (
			name     string
			password string
			isActive int
		)
		fmt.Println("Type name : ")
		fmt.Scan(&name)
		fmt.Println("Password : ")
		fmt.Scan(&password)
		fmt.Println("isActive : ")
		fmt.Scan(&isActive)

		_, err := db.Exec("INSERT INTO users(username, password, isActive) VALUES(?,?, ?);", name, password, isActive)
		if err != nil {
			log.Fatal(err.Error())
		}
	case 2:
		var (
			username string
			password int
		)
		fmt.Println("Username : ")
		fmt.Scan(&username)
		fmt.Println("Password : ")
		fmt.Scan(&password)
		_, err = db.Exec("UPDATE users SET password = ? WHERE username = ?;", password, username)
		if err != nil {
			log.Panic(err.Error())
		}
	case 3:
		var (
			username string
		)
		fmt.Println("Username : ")
		fmt.Scan(&username)
		_, err = db.Exec("DELETE FROM users WHERE username = ?;", username)
		if err != nil {
			log.Panic(err.Error())
		}
	case 4:
		var (
			id       int
			name     string
			password string
			isActive int
			username string
		)
		fmt.Scan(&username)
		res, err := db.Query("SELECT * FROM users WHERE username = ?;", username)
		if err != nil {
			log.Panic(err.Error())
		}
		//fmt.Println(res.Columns())
		if res.Next() {
			for res.Next() {
				err = res.Scan(&id, &name, &password, &isActive)
				if err != nil {
					log.Fatal(err.Error())
				}
				
				//fmt.Println(username)
				log.Println( id, name, password, isActive)
			}

		} else {
			fmt.Println("Bunday user mavjud emas")
		}
	}
}
