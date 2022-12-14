package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateTable(name string){
	db, err := sql.Open("mysql","root:Aa12345$@tcp(127.0.0.1:3306)/" + name)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS kutbxona(id int NOT NULL AUTO_INCREMENT, bookName varchar(45), pressYear int, authore varchar(55), PRIMARY KEY(id), UNIQUE(id));")
	if err != nil {
		panic(err.Error())
	}
}