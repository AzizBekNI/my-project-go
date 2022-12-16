package file

import (
	"errors"
	"os"
)

func FileRename() {
	oldPath := "test.txt"
	newPath := "./file_uchun/example.txt"
if _, err := os.Stat(oldPath); errors.Is(err, os.ErrNotExist){
	print("file not exist")
}else{
	err2 := os.Rename(oldPath, newPath)
	if err2 != nil {
		panic(err2)
	}else{
		print("file moved")
	}
}
}
