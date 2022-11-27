package test

import (
	"test1/methods"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	list := []methods.CSVRecord{

		{"Anita", "anita@email.com", "California"},
		{"Aron", "aron.bla@email.com", "California"},
		{"Aron", "aron.bla@email.com", "California"},
		{"Crina", "ggl@test.com", "Letcani"},
	}
	got, _ := methods.RemoveDuplicates(list)
	expected := []methods.CSVRecord{

		{"Anita", "anita@email.com", "California"},
		{"Aron", "aron.bla@email.com", "California"},
		{"Crina", "ggl@test.com", "Letcani"},
	}
	result := methods.Equals(got, expected)
	if result == true {
		t.Log("Test successful!")
	} else {
		t.Error("Test failed!")
	}
}
