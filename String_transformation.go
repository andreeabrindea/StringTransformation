package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
)

type CSVRecord struct {
	fname    string
	email    string
	location string
}

func createOutput(data [][]string) []CSVRecord {
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

	inputList := createOutput(data)

	sort.Slice(inputList[:], func(i, j int) bool {
		return inputList[i].fname < inputList[j].fname
	})

	for i := 1; i < len(inputList); i++ {
		if inputList[i-1].fname == inputList[i].fname {
			copy(inputList[i:], inputList[i+1:])
			inputList[len(inputList)-1] = CSVRecord{}
			inputList = inputList[:len(inputList)-1]
		}
	}

	outputFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvwriter := csv.NewWriter(outputFile)

	var lines [][]string
	firstLetter := ""
	for i, v := range inputList {
		if v.fname[0:1] != firstLetter {
			firstLetter = v.fname[0:1]
			firstLetterAndDots := firstLetter + ":"
			line := []string{firstLetterAndDots}
			newLine := []string{}
			if i != 0 {
				lines = append(lines, newLine)
			}

			lines = append(lines, line)
		}
		line := []string{v.fname, v.email, v.location}
		lines = append(lines, line)

	}

	csvwriter.WriteAll(lines)

	csvwriter.Flush()
	outputFile.Close()
}
