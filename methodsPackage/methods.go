package methodsPackage

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CSVRecord struct {
	Fname    string
	Email    string
	Location string
}

func Input(n int) []CSVRecord {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the records")
	var lines []CSVRecord
	for i := 0; i <= n; i++ {
		scanner.Scan()
		line := scanner.Text()
		data := strings.Split(line, ",")
		if i == 0 {
			continue
		}
		lines = append(lines, CSVRecord{Fname: strings.TrimSpace(data[0]), Email: strings.TrimSpace(data[1]), Location: strings.TrimSpace(data[2])})

	}
	return lines
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
