package file 

import (
	"os"
	"fmt"
)
func FileInfo(){
	path := "./test.txt"
fileInfo, err := os.Stat(path)
if err != nil {
	panic(err)
}
fmt.Printf("File Info %v\n", fileInfo)
}