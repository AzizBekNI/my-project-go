package file 

import (
	"os"
	"fmt"
)

func OpenFile2(fName string){
	file, err := os.OpenFile(fName, os.O_RDWR,0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprint(file)
}