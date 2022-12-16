package file

import (
	"os"
	"fmt"
)

func OpenFile(){
	path := "test.txt"
	file, err :=os.Open(path)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File Info : %v\n ", fileInfo)
}