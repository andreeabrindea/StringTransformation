package testPackage

import (
	"test1/methodsPackage"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	_, err := methodsPackage.ReadFromFile("inputTest1.csv")
	if err != nil {
		t.Error("Unable to read input file")
	}
	t.Log("Test successful")
}

func TestParse(t *testing.T) {
	//data, err := readFromFile("inputTest1.csv")
	//if err != nil {
	//	fmt.Println("Something went wrong")
	//}
	//for i, _ := range data {
	//	fmt.Print(data[i])
	t.Log("Test successful")
}
