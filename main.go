package main

import (
	"fmt"
	file "my-project/file"
)

func main() {
	var (
		src string
		dst string
	)
	fmt.Scan(&src, &dst)
	fmt.Println(file.CopyFile(src, dst))
}
