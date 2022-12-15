package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDb(name string){
	db, err := sql.Open("mysql", "root:Aa12345$@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
_, err = db.Exec("CREATE DATABASE " + name)
if err != nil {
	log.Fatal(err)
}
}