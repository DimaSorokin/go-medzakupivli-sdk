package main

import (
	"github.com/DimaSorokin/go-medzakupivli-sdk/client"
	"log"
	"time"
)

func main() {
	c := client.New()

	hospitalSamples, err := c.GetProtective(client.CovidArgs{
		Edrpou:     "01982591",
		ReportData: time.Now(),
	})

	if err != nil {
		log.Fatalf("error getting hospital list: %v", err)
	}

	log.Println("We got the response:", hospitalSamples)

}
