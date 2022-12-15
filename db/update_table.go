package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Update() {
	db, err := sql.Open("mysql", "root:Aa12345$@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	n := "Javoh"
	on := ""
	_, err = db.Exec("UPDATE users SET username = ? WHERE username = ?;", n, on)
	if err != nil {
		log.Fatal(err)
	}
}
