package csvutil

import (
	"fmt"
	"os"
	"encoding/csv"
)

type record map[string]string

func CSVMapReader(fname string) (document []record) {

	src := open_file(fname)
	
	reader, _ := csv.NewReader(src).ReadAll()

	headers := reader[0]

	for idx, row := range reader {

		if idx > 0 {

			record := make(record)

			for i, head := range headers {
				record[head] = row[i]
			}

			document = append(document, record)
		}
	}

	return
}

func open_file(filename string) (file *os.File) {

	file, err := os.Open(filename)
	
	if err != nil {
		fmt.Println("Error:", err)
	}

	return
}