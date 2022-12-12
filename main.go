package main

import (
	//"fmt"
	"fmt"
	db "my-project/datebase"
)

func main() {
	//fmt.Println(M)
	fmt.Println(db.SayHi())
	var id int 
	fmt.Scan(&id)
	fmt.Println(db.CanPurchase1(id))
	//db.SayHi()
	//sql0()
}