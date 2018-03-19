const (
    FILE           = "./MX.txt"
    FILE_DELIMITER = '\t'
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
        d.colonia = col[2]
        d.ciudad = col[3]
        d.delegacion = col[5]
        n++
        if n <= TOTAL_ROWS {
            fmt.Printf("%v,%v,%v;\n", d.colonia, d.ciudad, d.delegacion)
        }
    }

}
