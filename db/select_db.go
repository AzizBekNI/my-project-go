package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SelectTable() {
	var (
		id       int
		name     string
		password string
		isActive string
	)
	db, err := sql.Open("mysql", "root:Aa12345$@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &password, &isActive)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, password, isActive)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
