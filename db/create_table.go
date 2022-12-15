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
	_, err = db.Exec("CREATE TABLE `product` (`id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,`name` varchar(250) NOT NULL,`price` double NOT NULL,`quantity` int(11) NOT NULL,`status` tinyint(1) NOT NULL ) ENGINE=MyISAM DEFAULT CHARSET=latin1;")
	if err != nil {
		panic(err.Error())
	}
}