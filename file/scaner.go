package file

import (
	"bufio"

	"fmt"
	"os"
)

// bu func bitta qator tugaguncha o'qidi va file ga yozadi
func Scanner() {
	var fName string
	fmt.Scanln(&fName)
	file, err := os.OpenFile(fName, os.O_APPEND, 0666)
	if err != nil {
		panic("file not found !")
	}
	defer file.Close()
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	err = scan.Err()
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(scan.Text())
	if err != nil {
		panic(err)
	} else {
		print("write bytes")
	}
}
