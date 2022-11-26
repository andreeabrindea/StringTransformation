package methodsPackage

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type CSVRecord struct {
	Fname    string
	Email    string
	Location string
}

func WriteToFile(fileName string, newLines []CSVRecord) {
	inputFile, err := os.Create("input.csv") //create the input.csv file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		inputFile.Close()
	}()

	csvwriter := csv.NewWriter(inputFile)
	//csvwriter.LazyQuotes = true
	defer func() {
		csvwriter.Flush()
	}()

}

func ReadFromFile(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	//reading the file
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}
	err1 := f.Close()
	if err1 != nil {
		return nil, err1
	}

	return data, nil
}

func Parse(data [][]string) []CSVRecord {
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
	copy(list[index:], list[index+1:])
	list[len(list)-1] = CSVRecord{}
	return list[:len(list)-1], nil
}

func RemoveDuplicates(list []CSVRecord) ([]CSVRecord, error) {
	for i := 1; i < len(list); i++ { //removes duplicates
		if list[i-1].Fname == list[i].Fname {
			newList, err := remove(list, i)
			if err != nil {
				return nil, err
			}
			list = newList
		}
	}
	return list, nil
}
