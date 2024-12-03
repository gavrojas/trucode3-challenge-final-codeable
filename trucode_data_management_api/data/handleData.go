package data

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"trucode.com/final-challenge/models"
)

func LoadDataFromFile(path string) ([]models.PeopleData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []models.PeopleData
	// Reader porque mi archivo es grande.
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n') //<-- con Read String leermos cada línea del archivo
		if err != nil {
			// EOF = End of file si alcanzamos el final del archivo, es decir que no hay más datos por leer.
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		line = strings.TrimSpace(line)
		fields := strings.Split(line, ", ")

		if len(fields) < 15 {
			continue // Si no hay suficientes campos, saltamos esta línea
		}

		age, _ := strconv.Atoi(fields[0])
		fnlwgt, _ := strconv.Atoi(fields[2])
		educationNum, _ := strconv.Atoi(fields[4])
		capitalGain, _ := strconv.Atoi(fields[10])
		capitalLoss, _ := strconv.Atoi(fields[11])
		hoursPerWeek, _ := strconv.Atoi(fields[12])

		record := models.PeopleData{
			Age:           age,
			Workclass:     models.Workclass(fields[1]),
			Fnlwgt:        fnlwgt,
			Education:     models.Education(fields[3]),
			EducationNum:  educationNum,
			MaritalStatus: models.MaritalStatus(fields[5]),
			Occupation:    models.Occupation(fields[6]),
			Relationship:  models.Relationship(fields[7]),
			Race:          models.Race(fields[8]),
			Sex:           models.Sex(fields[9]),
			CapitalGain:   capitalGain,
			CapitalLoss:   capitalLoss,
			HoursPerWeek:  hoursPerWeek,
			NativeCountry: fields[13],
			Income:        models.Income(fields[14]),
		}
		data = append(data, record)
	}

	return data, nil
}

func CountLinesInFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lines := 0

	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return 0, err
		}
		lines++
	}

	return lines, nil
}

func IsDataAlreadyLoaded(db *gorm.DB, filePath string) (bool, error) {
	fileLineCount, err := CountLinesInFile(filePath)
	if err != nil {
		return false, err
	}

	var dbCount int64
	db.Model(&models.PeopleData{}).Count(&dbCount)
	fmt.Printf("dbcount %d, filelines %d\n", dbCount, fileLineCount)

	// -1 por la fila de encabezados de la tabla
	return (int64(fileLineCount-1) == dbCount || int64(fileLineCount-1) <= dbCount), nil
}

func ApplyFilters(query *gorm.DB, education, maritalStatus, occupation, minAge, maxAge, income string) *gorm.DB {
	if education != "" {
		query = query.Where("education = ?", education)
	}
	if maritalStatus != "" {
		query = query.Where("marital_status = ?", maritalStatus)
	}
	if occupation != "" {
		query = query.Where("occupation = ?", occupation)
	}
	if minAge != "" {
		query = query.Where("age >= ?", minAge)
	}
	if maxAge != "" {
		query = query.Where("age <= ?", maxAge)
	}
	if income != "" {
		query = query.Where("income = ?", income)
	}

	return query
}
