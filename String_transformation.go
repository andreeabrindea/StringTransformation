package main

import (
	"fmt"
	"log"
	"sort"
	"test1/methods"
)

func main() {
	n := 0
	fmt.Println("How many records?")
	_, err := fmt.Scanln(&n)

	input := methods.Input(n)

	sort.Slice(input[:], func(i, j int) bool { //sort the records (because in my head seemed easier to remove duplicates later on)
		return input[i].Fname < input[j].Fname
	})

	output, err := methods.RemoveDuplicates(input) //sorted and without duplicates
	if err != nil {
		log.Fatal(err)
	}
	if len(input) != 0 {
		fmt.Println("----------------------")
		fmt.Println("output:")
	}
	firstLetter := ""
	for _, v := range output {
		if v.Fname[0:1] != firstLetter {
			firstLetter = v.Fname[0:1] //get the first letter of each row of records

			fmt.Println(" ")
			fmt.Println(firstLetter + ":")
		}
		lineOutput := fmt.Sprintf("%s, %s, %s", v.Fname, v.Email, v.Location) //convert the records to string, so it can be written in the field

		fmt.Println(lineOutput)
	}
}
