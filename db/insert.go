<<<<<<< HEAD
package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Insert(name string) {
	db, err := sql.Open("mysql", "root:Aa12345$@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `product` (`id`, `name`, `price`, `quantity`, `status`) VALUES(1, 'Tivi 1', 200, 4, 1),(2, 'Tivi 2', 500, 6, 0),(3, 'Laptop 1', 250, 11, 1),(4, 'Laptop 2', 241, 6, 0),(5, 'Laptop 3', 600, 11, 0),(6, 'Mobile 1', 120, 20, 1),(7, 'Mobile 2', 210, 25, 0);")
	if err != nil {
		log.Fatal(err)
	}
}
=======
package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Insert(name string) {
	db, err := sql.Open("mysql", "root:Aa12345$@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `product` (`id`, `name`, `price`, `quantity`, `status`) VALUES(1, 'Tivi 1', 200, 4, 1),(2, 'Tivi 2', 500, 6, 0),(3, 'Laptop 1', 250, 11, 1),(4, 'Laptop 2', 241, 6, 0),(5, 'Laptop 3', 600, 11, 0),(6, 'Mobile 1', 120, 20, 1),(7, 'Mobile 2', 210, 25, 0);")
	if err != nil {
		log.Fatal(err)
	}
}
>>>>>>> c799814a28fe0b6c5b0331ef28c89a611725ddac
