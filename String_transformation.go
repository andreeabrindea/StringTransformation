package main

import (
	"fmt"
	"log"
	"sort"
	"test1/methodsPackage"
)

func main() {
	n := 0
	fmt.Println("How many records ?")
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	lines := methodsPackage.Input(n)

	sort.Slice(lines[:], func(i, j int) bool { //sort the records (because in my head seemed easier to remove duplicates later on)
		return lines[i].Fname < lines[j].Fname
	})

	newLines, err := methodsPackage.RemoveDuplicates(lines) //sorted and without duplicates
	if err != nil {
		log.Fatal(err)
	}

	firstLetter := ""
	for _, v := range newLines {
		if v.Fname[0:1] != firstLetter {
			firstLetter = v.Fname[0:1] //get the first letter of each row of records

			fmt.Println(" ")
			fmt.Println(firstLetter + ":")
		}
		line := fmt.Sprintf("%s %s %s", v.Fname, v.Email, v.Location) //convert the records to string, so it can be written in the field

		fmt.Println(line)
	}
}
