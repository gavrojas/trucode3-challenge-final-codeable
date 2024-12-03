package data

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"trucode.com/final-challenge/models"
)

func ExportToCSV(c *gin.Context, data []models.PeopleDataFrontend) {
	// Create csv file
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=people_data.csv")
	c.Writer.Header().Set("Content-Type", "text/csv")
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Write headers
	writer.Write([]string{"ID", "Age", "Education", "Marital status", "Ocupation", "Income"})

	// write data
	for _, record := range data {
		writer.Write([]string{
			strconv.Itoa(int(record.ID)),
			strconv.Itoa(record.Age),
			record.Education,
			record.MaritalStatus,
			record.Occupation,
			record.Income,
		})
	}
}

func ImportCSV(c *gin.Context) ([][]string, error) {
	// get csv file from request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	// open file
	fileReader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	defer fileReader.Close()

	// create csv reader
	reader := csv.NewReader(fileReader)
	reader.FieldsPerRecord = -1

	// read csv file
	records, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	return records[1:], nil

}
