package file

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dst string)(int64, error){
file, err := os.Stat(src)
if err != nil {
	return 0, err
}
if !file.Mode().IsRegular(){
	return 0, fmt.Errorf(" %s file is not regular : ", src)
}
sourc, err := os.Open(src)
if err != nil {
	return 0, err
}
defer sourc.Close()
destination, err := os.Create(dst)
if err != nil {
	return 0, err
}
defer destination.Close()
isByte, err := io.Copy(destination, sourc)
return isByte, err
}
