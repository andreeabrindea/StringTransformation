package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type CSVRecords struct {
	fname    string
	email    string
	location string
}

func main() {
	//opening the csv file
	f, error := os.Open("input.csv")

	if error != nil {
		log.Fatal(error)
	}

	//reading the file
	reader := csv.NewReader(f)

	for {
		row, err := reader.Read()
		if err == io.EOF { // Here it stops reading because EOF is the error returned by Read when no more input is available.
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//printing line by line
		fmt.Printf("%+v\n", row)
	}
	defer f.Close()
}
