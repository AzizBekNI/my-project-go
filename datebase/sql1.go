package datebase

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)


func Sql1() {
	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS kutbxona(id int NOT NULL AUTO_INCREMENT, bookName varchar(45), pressYear int, authore varchar(55), PRIMARY KEY(id), UNIQUE(id));")
	// if err != nil {
	// 	panic(err.Error())
	// }
	
	db, err := sql.Open("mysql", "root:Aa12345$@/testDB")
	if err != nil {
		log.Panic(err.Error())
	}
	defer db.Close()
	
	var n int
	fmt.Printf("1 Kitob qo'shish \n2 Kitob tanlash\n")
	fmt.Scan(&n)
	
	switch n {
	case 1:
		var (
			bookName  string
			pressYear int
			authore   string
		)
		fmt.Println("Book Name : ")
		fmt.Scan(&bookName)
		fmt.Println("Press Year : ")
		fmt.Scan(&pressYear)
		fmt.Println("Authore : ")
		fmt.Scan(&authore)
		defer db.Close()

		_, err = db.Exec("INSERT INTO kutbxona(bookName, pressYear, authore) VALUES(?,?,?);", bookName, pressYear, authore)
		if err != nil {
			panic(err.Error())
		}
	case 2:
		var (
			id int
			bookName string
			b_authore string
			pressYear int
			//name string
			//authore string
		)
		fmt.Println("Kitob nomi va authorni nomini kriting")
		//fmt.Scan(&name)
		//fmt.Scan(&authore)
		res, err := db.Query("SELECT * FROM kutbxona;")
		if err != nil {
			log.Fatal(err.Error())
		}
		if res.Next(){
			for res.Next(){
				err = res.Scan(&id, &bookName, &b_authore, &pressYear)
				if err != nil {
					log.Panic(err.Error())
				}
				fmt.Printf("%v, %v, %v",bookName, b_authore, pressYear )
			}
		}else {
			fmt.Println("Kutbxonada bunday kitob mavjud emas")
		}
	}
}
//Muhhammad ibn Abdul Vahob At Tamimiy