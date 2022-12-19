package file

import (
	"fmt"
	"os"
	//"golang.org/x/tools/go/analysis/passes/printf"
)
// bu func birinchi probilgacha o'qiydi va file ga yozadi
func Write(fName string) {
	file, err := os.OpenFile(fName, os.O_APPEND, 0666)
	if err != nil {
		panic("file not found !")
	}
	defer file.Close()
	bytes := "Bu faylga ma'lumot qo'shildi !"
	fmt.Scan(&bytes)
	_, err = file.WriteString(bytes)
	if err != nil {
		panic(err)
	} else {
		print("write bytes")
	}
}
