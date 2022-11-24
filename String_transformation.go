package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

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
		if err == io.EOF {
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
