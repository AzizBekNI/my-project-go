package datebase

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CanPurchase1(id int) (bool, error) {
	db, err := sql.Open("mysql", "root:Aa12345$@/testDB")
	if err != nil {
		log.Panic(err.Error())
	}
	defer db.Close()
	var  enough bool
	// Query for a value based on a single row.
	err = db.QueryRow("SELECT * from users where id = ?", id).Scan(&enough)
	if  err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("canPurchase %d: unknown album", id)
		}
		return false, fmt.Errorf("canPurchase %d: ", id)
	}
	return enough, nil
}
