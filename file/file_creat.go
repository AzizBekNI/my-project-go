package file

import(
	"os"
)

var (
	err error
	newFile *os.File
)
func FileCreate(){
newFile, err = os.Create("test.txt")
if err != nil {
	panic(err)
}
}