package main

import (
	"encoding/csv"
	"errors"
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

func parse(data [][]string) []CSVRecord {
	var inputList []CSVRecord

	for i, line := range data {
		// omit header line
		if i == 0 {
			continue
		}
		rec := CSVRecord{strings.TrimSpace(line[0]), strings.TrimSpace(line[1]), strings.TrimSpace(line[2])}
		inputList = append(inputList, rec)
	}
	return inputList
}

func remove(list []CSVRecord, index int) ([]CSVRecord, error) {
	if len(list) == index+1 {
		return nil, errors.New("index out of bounds")
	}
	copy(list[index:], list[index+1:])
	list[len(list)-1] = CSVRecord{}
	return list[:len(list)-1], nil
}

func removeDuplicates(list []CSVRecord) ([]CSVRecord, error) {
	for i := 1; i < len(list); i++ { //removes duplicates
		if list[i-1].fname == list[i].fname {
			newList, err := remove(list, i)
			if err != nil {
				return nil, err
			}
			list = newList
		}
	}
	return list, nil
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

	inputList := parse(data) //creating a list of those records from input.csv

	sort.Slice(inputList[:], func(i, j int) bool { //sort the records (because in my head seemed easier to remove duplicates later on)
		return inputList[i].fname < inputList[j].fname
	})

	newList, err := removeDuplicates(inputList) //sorted and without duplicates
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create("output.csv") //create the output.csv file
	if err != nil {
		log.Fatal(err)
	}

	csvwriter := csv.NewWriter(outputFile)

	var lines [][]string
	firstLetter := ""
	for i, v := range newList {
		if v.fname[0:1] != firstLetter {
			firstLetter = v.fname[0:1]          //get the first letter of each row of records
			line := []string{firstLetter + ":"} //convert firstLetter to a string so it can be written in the output.csv file
			newLine := []string{}               //an empty line is needed before each first letter
			if i != 0 {                         //except the first one
				lines = append(lines, newLine)
			}

			lines = append(lines, line)
		}

		line := []string{v.fname, v.email, v.location} //convert the records to string so it can be written in the field
		lines = append(lines, line)
	}

	csvwriter.WriteAll(lines)

	csvwriter.Flush()
	outputFile.Close()
}
