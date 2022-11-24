package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strings"
)

type CSVRecord struct {
	fname    string
	email    string
	location string
}

func createAListOfRecords(data [][]string) []CSVRecord {
	var inputList []CSVRecord

	for i, line := range data {
		if i > 0 { // omit header line
			var rec CSVRecord
			for j, field := range line {
				if j == 0 {
					rec.fname = field
				}
				if j == 1 {
					rec.email = field
				}
				if j == 2 {
					rec.location = field
				}
			}
			inputList = append(inputList, rec)

		}
	}
	return inputList
}

func main() {
	//opening the csv file
	f, error := os.Open("input.csv")

	if error != nil {
		log.Fatal(error)
	}

	//reading the file
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	inputList := createAListOfRecords(data) //creating a list of those records from input.csv

	sort.Slice(inputList[:], func(i, j int) bool { //sort the records (because in my head seemed easier to remove duplicates later on)
		return inputList[i].fname < inputList[j].fname
	})

	for i := 1; i < len(inputList); i++ { //eliminates duplicates
		if inputList[i-1].fname == inputList[i].fname {
			copy(inputList[i:], inputList[i+1:])
			inputList[len(inputList)-1] = CSVRecord{}
			inputList = inputList[:len(inputList)-1]
		}
	}

	outputFile, err := os.Create("output.csv") //create the output.csv file
	if err != nil {
		log.Fatal(err)
	}

	csvwriter := csv.NewWriter(outputFile)

	var lines [][]string
	firstLetter := ""
	for i, v := range inputList {
		if v.fname[0:1] != firstLetter {
			firstLetter = v.fname[0:1]              //get the first letter of each row of records
			firstLetterAndDots := firstLetter + ":" //add that :
			line := []string{firstLetterAndDots}    //convert firstLetter to a string so it can be written in the output.csv file
			newLine := []string{}                   //an empty line is needed before each first letter
			if i != 0 {                             //except the first one
				lines = append(lines, newLine)
			}

			lines = append(lines, line)
		}

		line := []string{v.fname, v.email, v.location} //convert the records to string so it can be written in the field

		for i := range line {
			line[i] = strings.ReplaceAll(line[i], " ", "") // needed to do that so the email and location won't have a space before
		}
		lines = append(lines, line)

	}

	csvwriter.WriteAll(lines)

	csvwriter.Flush()
	outputFile.Close()
}
