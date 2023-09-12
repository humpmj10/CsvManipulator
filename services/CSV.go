package services

import (
	"encoding/csv"
	"log"
	"mime/multipart"
	"sort"
	"strconv"
	"strings"
)

type CsvProcessor interface {
	ProcessCsv(file multipart.File, shouldSort bool, columnToSort int) ([][]string, error)
}

type CsvService struct{}

func ProcessCsv(file multipart.File, shouldSort bool, column int) ([][]string, error) {
	log.Println("reading csv file")

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var newRows [][]string
	for _, row := range rows {
		var newRow []string
		for _, column := range row {
			newRow = append(newRow, strings.ToUpper(column))
		}
		newRows = append(newRows, newRow)
	}

	if shouldSort {
		sortCSV(newRows, column)
	}

	return newRows, nil
}

func (c *CsvService) ProcessCsv(file multipart.File, shouldSort bool, columnToSort int) ([][]string, error) {
	return ProcessCsv(file, shouldSort, columnToSort)
}

func sortCSV(records [][]string, column int) {

	_, err := strconv.Atoi(records[1][column])
	isNumeric := err == nil

	sort.Slice(records[1:], func(i, j int) bool {
		iVal := records[i+1][column]
		jVal := records[j+1][column]
		if isNumeric {
			jInt, _ := strconv.Atoi(jVal)
			iInt, _ := strconv.Atoi(iVal)
			return iInt < jInt
		}
		return iVal < jVal
	})
}
