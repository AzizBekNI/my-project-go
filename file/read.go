package file

import(
	"os"
	"fmt"
)

func Read(){
	var fName string
	fmt.Scanln(&fName)
	file, err := os.Open("example.txt")
	cheak(err)
	defer file.Close()
}
func cheak(e error){
	if e != nil {
		panic(e)
	}
}