package gucsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvLine struct {
	col_0 string
	col_1 string
}

func ReadCSVFile(tFilePath string) {

	f, err := os.Open(tFilePath)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	if err == nil {

		for _, line := range lines {
			data := CsvLine{col_0: line[0], col_1: line[1]}
			fmt.Println(data.col_0 + " " + data.col_1)
		}
	}
}
