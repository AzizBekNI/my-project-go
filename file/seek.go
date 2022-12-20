 package file

import (
	"fmt"
	"log"
	"os"
)

func Seek_() {

	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for {
		o2, err := file.Seek(0, 1)
		if err != nil {
			log.Fatal(err)
		}
		b1 := make([]byte, 30)
		n1, err := file.Read(b1)
		if err != nil {
			log.Fatal("Failed here ", err)
		}
		fmt.Printf("start:%d %d bytes: %s\n", o2, n1, string(b1))
	}
} 


 