package file
 
import (
	"os"
	"fmt"
)

func FileInfo2(){
	file, err := os.Stat("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File Name : " ,file.Name())
	fmt.Println("File Size in Byte : ",file.Size())
	fmt.Println("File Filepermision  : ",file.Mode())
	fmt.Println("File Last Modifaiyed : ",file.ModTime())
	fmt.Println("File isDirector : " ,file.IsDir())

}