package file

import (
	"fmt"
	"io/ioutil"
	
)
func DirRead(){
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Printf("%s %04o %s\n", file.Mode(), file.Mode().Perm(), file.Name())
	}
}