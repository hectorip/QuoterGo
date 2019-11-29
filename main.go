package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:name", helloHandler)
	r.GET("/random", randomQuote)
	r.Run()
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Hola %s", c.Param("name")),
	})
}

func randomQuote(c *gin.Context) {
	quotes := getQuotes()
	q := pickOne(quotes)
	c.JSON(200, gin.H{
		"quote": fmt.Sprintf("%s Por %s", q.Text, q.Author),
	})
}

type Quote struct {
	Text   string
	Author string
}

func getQuotes() (quotes []Quote) {
	// Open the file
	csvfile, err := os.Open("sources/spanish.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	r.LazyQuotes = true

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quotes = append(quotes, Quote{record[0], record[1]})
		fmt.Printf("Quote: %s By %s\n", record[0], record[1])
	}
	return
}

func pickOne(choices []Quote) Quote {
	rand.Seed(time.Now().Unix())
	index := rand.Int() % len(choices)
	return choices[index]
}
