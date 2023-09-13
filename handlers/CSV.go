package handlers // Package handlers import "CsvManipulator/handlers"

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"
)

// UploadHandler godoc
// @Summary Upload a CSV file
// @Description Upload a CSV file and return a CSV file with all columns in uppercase
// @Accept mpfd
// @Produce mpfd
// @Param csvFile formData file true "CSV file to upload"
// @Success 200 {string} string "CSV file with all columns in uppercase"
// @Failure 400 {string} string "Unable to get file"
// @Router /upload [post]
func (api *API) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	queryParams := r.URL.Query()
	shouldSort := queryParams.Get("sort") == "true"
	columnToSort := queryParams.Get("column")
	columnIndex, err := strconv.Atoi(columnToSort)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse column index", http.StatusBadRequest)
		return
	}

	// Parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
	err = r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("csvFile") // Retrieve file from posted form data, needs to match
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Failed to close file: %v", err)
		}
	}()

	numRows, err := api.CsvService.ProcessCsv(file, shouldSort, columnIndex)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=output.csv")

	writer := csv.NewWriter(w)
	for _, newRow := range numRows {
		if err := writer.Write(newRow); err != nil {
			log.Println(err)
			http.Error(w, "Error writing CSV", http.StatusInternalServerError)
			return
		}
	}
	writer.Flush()
}

func (api *API) SortCSVHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Failed to close file: %v", err)
		}
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to read csv data", http.StatusBadRequest)
		return
	}

	sortColumn, err := strconv.Atoi(r.URL.Query().Get("column"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid sort parameter", http.StatusBadRequest)
		return
	}

	api.CsvService.SortCSV(records, sortColumn)

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=sorted.csv")
	writer := csv.NewWriter(w)
	err = writer.WriteAll(records)
	if err != nil {
		return
	}
	writer.Flush()
}
