package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"test1/methodsPackage"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := 0
	inputFile, err := os.Create("input.csv") //create the input.csv file
	if err != nil {
		log.Fatal(err)
	}

	csvwriter := csv.NewWriter(inputFile)
	defer func() {
		csvwriter.Flush()
	}()
	fmt.Println("How many records ?")
	fmt.Scanln(&n)
	fmt.Println("Enter the records")
	var lines []methodsPackage.CSVRecord
	for i := 0; i <= n; i++ {
		scanner.Scan()
		line := scanner.Text()
		data := strings.Split(line, ",")
		if i == 0 {
			continue
		}
		lines = append(lines, methodsPackage.CSVRecord{Fname: strings.TrimSpace(data[0]), Email: strings.TrimSpace(data[1]), Location: strings.TrimSpace(data[2])})

	}

	sort.Slice(lines[:], func(i, j int) bool { //sort the records (because in my head seemed easier to remove duplicates later on)
		return lines[i].Fname < lines[j].Fname
	})

	newLines, err := methodsPackage.RemoveDuplicates(lines) //sorted and without duplicates
	if err != nil {
		log.Fatal(err)
	}

	var linesOutput [][]string
	firstLetter := ""
	for i, v := range newLines {
		if v.Fname[0:1] != firstLetter {
			firstLetter = v.Fname[0:1] //get the first letter of each row of records

			fmt.Println(firstLetter + ":")
			var newLine []string //an empty line is needed before each first letter
			if i != 0 {          //except the first one
				linesOutput = append(linesOutput, newLine)
			}
		}
		line := fmt.Sprintf("%s %s %s", v.Fname, v.Email, v.Location) //convert the records to string, so it can be written in the field

		fmt.Println(line)
	}
}
