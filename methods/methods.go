package methods

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type CSVRecord struct {
	Fname    string
	Email    string
	Location string
}

// Equals verify if two lists of  type CSVRecord are equal
func Equals(c []CSVRecord, k []CSVRecord) bool {
	if len(c) != len(k) {
		return false
	}
	for i := 0; i < len(c); i++ {
		b := c[i] == k[i]
		if b == false {
			return false
			break
		}
	}

	return true
}

// Input handles user input of the records
func Input(n int) []CSVRecord {
	if n <= 0 {
		fmt.Println("invalid number of records")

	}
	scanner := bufio.NewScanner(os.Stdin)
	if n > 0 {
		fmt.Println("The first line would be the header.")
		fmt.Println("Enter the records with a comma between them.")
	}
	var lines []CSVRecord
	for i := 0; i <= n; i++ {
		if n <= 0 {
			break
		}
		scanner.Scan()
		line := scanner.Text()
		data := strings.Split(line, ",")
		if len(data) != 3 { //verify if there are 3 entries on a line: full-name, email and location
			println("invalid number entries per line")
			break
		}
		if i == 0 { //omit the header
			continue
		}
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`) //verify if the second entry is a valid email
		if emailRegex.MatchString(strings.TrimSpace(data[1])) == false {
			fmt.Println("invalid email")
			break

		}
		lines = append(lines, CSVRecord{Fname: strings.TrimSpace(data[0]), Email: strings.TrimSpace(data[1]), Location: strings.TrimSpace(data[2])})

	}

	return lines

}

// removes the element at the given index and returns the new shorted list
func remove(list []CSVRecord, index int) ([]CSVRecord, error) {
	copy(list[index:], list[index+1:])
	list[len(list)-1] = CSVRecord{}
	return list[:len(list)-1], nil
}

// RemoveDuplicates remove duplicates of records with the same full-name and returns the new shorted list
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
