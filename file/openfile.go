package file 

import (
	"os"
)

func OpenFile2(fName string){
	file, err := os.OpenFile(fName, os.O_RDWR,0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}