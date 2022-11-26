package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"test1/methodsPackage"
	//"test1/methodsPackage"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := 0
	inputFile, err := os.Create("input.csv") //create the input.csv file
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		inputFile.Close()
	}()

	csvwriter := csv.NewWriter(inputFile)
	defer func() {
		csvwriter.Flush()
	}()
	fmt.Println("How many records ?")
	fmt.Scanln(&n)
	fmt.Println("Enter the records")
	var lines [][]string
	for i := 0; i < n; i++ {
		scanner.Scan()
		text := scanner.Text()
		lines = append(lines, []string{text})

	}
	err = csvwriter.WriteAll(lines)
	if err != nil {
		return
	}

	data, err := methodsPackage.ReadFromFile("input.csv")
	if err != nil {
		log.Fatal("Couldn't open the file")
	}
	fmt.Println(data)
	inputList := methodsPackage.Parse(data) //creating a list of those records from input.csv
	fmt.Println(inputList)
	//sort.Slice(inputList[:], func(i, j int) bool { //sort the records (because in my head seemed easier to remove duplicates later on)
	//	return inputList[i].Fname < inputList[j].Fname
	//})

	//newList, err := methodsPackage.RemoveDuplicates(inputList) //sorted and without duplicates
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//outputFile, err := os.Create("output.csv") //create the output.csv file
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//csvwriterOutput := csv.NewWriter(outputFile)
	//
	//var linesOutput [][]string
	//firstLetter := ""
	//for i, v := range newList {
	//	if v.Fname[0:1] != firstLetter {
	//		firstLetter = v.Fname[0:1]          //get the first letter of each row of records
	//		line := []string{firstLetter + ":"} //convert firstLetter to a []string, so it can be written in the output.csv file
	//		var newLine []string                //an empty line is needed before each first letter
	//		if i != 0 {                         //except the first one
	//			linesOutput = append(linesOutput, newLine)
	//		}
	//
	//		linesOutput = append(linesOutput, line)
	//	}
	//
	//	line := []string{v.Fname, v.Email, v.Location} //convert the records to string, so it can be written in the field
	//	linesOutput = append(linesOutput, line)
	//}
	//
	//err1 := csvwriterOutput.WriteAll(linesOutput)
	//if err1 != nil {
	//	return
	//}
	//
	//csvwriterOutput.Flush()
	//err2 := outputFile.Close()
	//if err2 != nil {
	//	return
	//}

}
