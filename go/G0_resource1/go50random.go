package main

import (
   "fmt"
   "encoding/csv"
   "net/http"
   "io"
   "os"
   "math/rand"
   "time"
)

func main() {

        http.HandleFunc("/", handler)
		
	http.HandleFunc("/elasticSearch", handler1)

        http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello World\n")
}

func handler1(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Salaries Elastic Search\n")
	//start_time := time.Now()

  // Loading csv file
  rFile, err := os.Open("./results.csv") //3 columns
  if err != nil {
    fmt.Println("Error:", err)
    return
   }
  defer rFile.Close()

  // Creating csv reader
  reader := csv.NewReader(rFile)

  lines, err := reader.ReadAll()
  if err == io.EOF {
      fmt.Println("Error:", err)
      return
  }

  // Creating csv writer
  wFile, err := os.Create("./result.csv")
  if err != nil {
      fmt.Println("Error:",err)
      return
  }
  defer wFile.Close()
  //writer := csv.NewWriter(wFile)

  // Read data, randomize columns and write new lines to results.csv
  rand.Seed(int64(time.Now().Nanosecond()))
  var col_index []int
  //var playerID string
  //var salary int
  //var team string
  //for i,line :=range lines{
  for i,line :=range lines {
    if i == 0 {
        //randomize column index based on the number of columns recorded in the 1st line
    ///   col_index = rand.Perm(len(line))
   // }
	
	if i == 50 {
		return
	}

    playerID := line[col_index[0]]
    salary := line[col_index[1]]
    team := line[col_index[2]]
    fmt.Fprintf(team)
   
    //fmt.Fprintf([]string{line[col_index[0]], line[col_index[1]], line[col_index[2]]})
    //writer.Write([]string{line[col_index[0]], line[col_index[1]], line[col_index[2]]}) //3 columns
    //writer.Flush
        i++
	}

}
}
