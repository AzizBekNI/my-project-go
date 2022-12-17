package file

import (
	"os"
)

var (
	err     error
	newFile *os.File
)

func FileCreate(fName string) {
	newFile, err = os.Create(fName)
	if err != nil {
		panic(err)
	}
}
