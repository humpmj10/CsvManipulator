package services

import (
	"encoding/csv"
	"strings"
	"testing"
)

func TestCsvService_SortCsv(t *testing.T) {
	content := `name,age,email
Bob,30,john@example.com
Jane,25,jane@example.com
Cindy,55,cindy@example.com`
	r := csv.NewReader(strings.NewReader(content))
	records, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	CsvService := CsvService{}
	CsvService.SortCSV(records, 0)

	if records[1][0] != "Bob" {
		t.Errorf("Expected Bob, got %s", records[1][0])
	}
	if records[2][0] != "Cindy" {
		t.Errorf("Expected Cindy, got %s", records[2][0])
	}
	if records[3][0] != "Jane" {
		t.Errorf("Expected Jane, got %s", records[3][0])
	}

	CsvService.SortCSV(records, 1)

	if records[1][1] != "25" {
		t.Errorf("Expected 25, got %s", records[1][1])
	}
	if records[2][1] != "30" {
		t.Errorf("Expected 30, got %s", records[2][1])
	}
	if records[3][1] != "55" {
		t.Errorf("Expected 55, got %s", records[3][1])
	}

	CsvService.SortCSV(records, 2)

	if records[1][2] != "cindy@example.com" {
		t.Errorf("Expected cindy@example.com, got %s", records[1][2])
	}
	if records[2][2] != "jane@example.com" {
		t.Errorf("Expected jane@example.com, got %s", records[2][2])
	}

}
