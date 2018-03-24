package main

import (
	"encoding/csv"
	"fmt"
	"os"
)


type Document struct {
	colonia, ciudad, delegacion string
}

const (
	FILE           = "./Salaries.csv"
	FILE_DELIMITER = ','
	TOTAL_ROWS     = 5
)

func main() {

	file, err := os.Open(FILE)
	printError(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = FILE_DELIMITER

	rows, err := reader.ReadAll()
	printError(err)

	fmt.Printf("\nrows type: %T\n\n", rows)

	d := new(Document)

	for n, col := range rows {
		d.colonia = col[0]
		d.ciudad = col[1]
		d.delegacion = col[2]
		n++
		if n <= TOTAL_ROWS {
		    fmt.Printf("%v,%v,%v;\n", d.colonia, d.ciudad, d.delegacion)
		}
	}

}

func printError(err error) {
	if err != nil {
		fmt.Printf("\nError: %v \n ", err.Error())
		os.Exit(1)
	}
}
