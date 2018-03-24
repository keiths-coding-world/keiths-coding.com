package main

import (
	"errors"
	"fmt"
	"net/http"
	"gopkg.in/olivere/elastic.v3"
	"reflect"
	"time"
)

const (
	indexName    = "applications"
	docType      = "log"
	appName      = "myApp"
	indexMapping = `{
				"mappings" : {
						"log" : {
					"properties" : {
					"app" : { "type" : "string", "index" : "not_analyzed" },
					"message" : { "type" : "string", "index" : "not_analyzed" },
					"time" : { "type" : "date" }
								}
							}
						}
					}`
)

type Log struct {
	App     string    `json:"app"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/GOElasticSearch", handler1)
	
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Baseball player search:\n")

func handler1(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Search Results:\n")

client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}

	err = createIndexWithLogsIfDoesNotExist(client)
	if err != nil {
		panic(err)
	}

	err = findAndPrintAppLogs(client)
	if err != nil {
		panic(err)
	}
func createIndexWithLogsIfDoesNotExist(client *elastic.Client) error {
	exists, err := client.IndexExists(indexName).Do()
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	res, err := client.CreateIndex(indexName).
		Body(indexMapping).
		Do()

	if err != nil {
		return err
	}
	if !res.Acknowledged {
		return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
	}

	return addLogsToIndex(client)
}

func addLogsToIndex(client *elastic.Client) error {
	for i := 0; i < 10; i++ {
		l := Log{
			App:     "myApp",
			Message: fmt.Sprintf("message %d", i),
			Time:    time.Now(),
		}

		_, err := client.Index().
			Index(indexName).
			Type(docType).
			BodyJson(l).
			Do()

		if err != nil {
			return err
		}
	}

	return nil
}

func findAndPrintAppLogs(client *elastic.Client) error {
	termQuery := elastic.NewTermQuery("app", appName)

	res, err := client.Search(indexName).
		Index(indexName).
		Query(termQuery).
		Sort("time", true).
		Do()

	if err != nil {
		return err
	}

	fmt.Println("Logs found:")
	var l Log
	for _, item := range res.Each(reflect.TypeOf(l)) {
		l := item.(Log)
		fmt.Printf("time: %s message: %s\n", l.Time, l.Message)
	}

	return nil
}
