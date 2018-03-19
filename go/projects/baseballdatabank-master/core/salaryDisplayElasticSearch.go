package main

import (
	"encoding/csv"
	"fmt"
	"os"
)


type Document struct {
	yearID, teamID, playerID, salary string
}

const (
    FILE = "./Salaries.txt"
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

   // d := new(Document)

   // for n, col := range rows {
   //     d.yearID = col[0]
   //     d.teamID = col[1]
   //     d.playerID = col[3]
   //	d.salary = col[4]
   //     n++
   //     if n <= TOTAL_ROWS {
   //         fmt.Printf("%v,%v,%v;\n", d.yearID, d.teamID, d.playerID, d.salary)
   //     }
   // }

}
