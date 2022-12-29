package json

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Csv() {
	file, err := os.Open("./file_uchun/data.csv")
	CheckError(err)
	defer file.Close()
	f := csv.NewReader(file)
	record, err := f.ReadAll()
	for _, row := range record{
		fmt.Println(row)
	}
}