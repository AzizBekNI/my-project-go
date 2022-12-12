package datebase

import (
	"database/sql"
	//"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CanPurchase() {
	db, err := sql.Open("mysql", "root:Aa12345$@/testDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()


	sqlStatement := `SELECT id, email FROM user WHERE id=1;`

var email string

var id int
// Replace 3 with an ID from your database or another random
// value to test the no rows use case.
row := db.QueryRow(sqlStatement)
switch err := row.Scan(&id, &email); err {
case sql.ErrNoRows:
  fmt.Println("No rows were returned!")
case nil:
  fmt.Println(id, email)
default:
  panic(err)
}
}
