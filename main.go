package main

import (
"fmt"
db "my-project/db"
)

func main() {
	var DbName string
	fmt.Scan(&DbName) 
	// db.CreateDb(Db_Name)
	//db.DropDb(DbName)
	db.SelectTable()
	db.Update()
	db.SelectTable()
}