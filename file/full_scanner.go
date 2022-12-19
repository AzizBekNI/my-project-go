package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Full_Scan() {
	var fName string
	print("input pass and name ! \n")
	fmt.Scanln(&fName)
	file, err := os.OpenFile(fName, os.O_APPEND, 0666)
	if err != nil {
		panic("file not found")
	}
	defer file.Close()
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for {
		scanner.Scan()
		_, err =file.WriteString(scanner.Text())
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("output:")
	for _, l := range lines {
		fmt.Println(l)
	}
}
