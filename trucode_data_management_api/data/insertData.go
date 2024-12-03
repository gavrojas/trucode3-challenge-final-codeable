package data

import (
	"errors"
	"log"
	"sync"

	"gorm.io/gorm"
	"trucode.com/final-challenge/models"
)

func InsertDataToDB(db *gorm.DB, data []models.PeopleData) {
	var wg sync.WaitGroup
	recordCh := make(chan models.PeopleData)

	// Goroutine para insertar registros en la base de datos
	go func() {
		for record := range recordCh {
			// Verificar si ya existen registros
			var existingRecord models.PeopleData
			if err := db.Where("age = ? AND workclass = ? AND fnlwgt = ? AND education = ? AND education_num = ? AND marital_status = ? AND occupation = ? AND relationship = ? AND race = ? AND sex = ? AND capital_gain = ? AND capital_loss = ? AND hours_per_week = ? AND native_country = ? AND income = ?", record.Age, record.Workclass, record.Fnlwgt, record.Education, record.EducationNum, record.MaritalStatus, record.Occupation, record.Relationship, record.Race, record.Sex, record.CapitalGain, record.CapitalLoss, record.HoursPerWeek, record.NativeCountry, record.Income).First(&existingRecord).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) /* No quiero ver los logs de record not found */ {
				// Log other errors
				log.Printf("Error checking existing record: %v", err)
				continue
			}
			db.Create(&record)
		}
	}()

	for _, record := range data {
		wg.Add(1)
		go func(r models.PeopleData) {
			defer wg.Done()
			recordCh <- r
		}(record)
	}

	wg.Wait()
	close(recordCh)
}
