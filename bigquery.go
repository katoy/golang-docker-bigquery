// Query: "SELECT corpus FROM publicdata:samples.shakespeare GROUP BY corpus;",
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	bigquery "google.golang.org/api/bigquery/v2"
)

const PROJECTID = "sample-1385"

func main() {
	data, err := ioutil.ReadFile("client.json")
	if err != nil {
		log.Fatal(err)
	}

	conf, err := google.JWTConfigFromJSON(data, bigquery.BigqueryScope)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext)

	bq, err := bigquery.New(client)
	if err != nil {
		log.Fatal(err)
	}

	// table 一覧の表示
	call := bq.Tables.List("publicdata", "samples")
	call.MaxResults(10)

	list, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	buf, err := json.Marshal(list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))

	// rquery の例
	conn, err := bigquery.New(client)
	query := "select * from publicdata:samples.shakespeare limit 20;"
	result, err := conn.Jobs.Query(PROJECTID, &bigquery.QueryRequest{
		Query: query,
	}).Do()

	for _, row := range result.Rows {
		for _, cell := range row.F {
			fmt.Print(cell.V)
			fmt.Print(",")
		}
		fmt.Print("\n")
	}
}
