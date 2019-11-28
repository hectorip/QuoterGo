package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:name", helloHandler)
	r.Run()
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Hola %s", c.Param("name")),
	})
}

func randomQuote(c *gin.Context) {
	quotes := getQuotes
}

type Quote struct {
	Text   string
	Author string
}

func getQuotes() (quotes []Quote) {
	// Open the file
	csvfile, err := os.Open("spanish.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

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
		fmt.Printf("Quote: %s By %s\n", record[0], record[1])
	}
}
