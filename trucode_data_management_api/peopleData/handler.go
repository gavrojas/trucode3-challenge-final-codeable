package peopleData

import (
	"net/http"
	"strconv"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"trucode.com/final-challenge/data"
	"trucode.com/final-challenge/database"
	"trucode.com/final-challenge/models"
)

// Obtener los datos de la tabla de personas en formato JSON PeopleDataFrontend
func GetPeopleData(c *gin.Context) {
	dbConn := database.CreateDbConnection()
	var dataFrontend []models.PeopleDataFrontend
	var totalCount int64
	var filters models.Filters

	authorizedUser, _ := c.Get("authorizedUser")
	user := authorizedUser.(models.User)
	if err := dbConn.Where("user_id = ?", user.ID).First(&filters).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			filters = models.Filters{}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user filters", "details": err.Error()})
			return
		}
	}

	// Filtros
	education := c.DefaultQuery("education", filters.Education)
	maritalStatus := c.DefaultQuery("marital_status", filters.MaritalStatus)
	occupation := c.DefaultQuery("occupation", filters.Occupation)
	minAge := c.DefaultQuery("min_age", strconv.Itoa(filters.MinAge))
	maxAge := c.DefaultQuery("max_age", strconv.Itoa(filters.MaxAge))
	income := c.DefaultQuery("income", filters.Income)

	sortBy := c.Query("sort_by")
	order := c.Query("order")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	export := c.Query("export")

	// Construir la consulta
	query := dbConn.Model(&models.PeopleData{}).
		Select("ID, age, education, marital_status, occupation, income")

	query = data.ApplyFilters(query, education, maritalStatus, occupation, minAge, maxAge, income)

	// conteo de registros
	if err := query.Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if sortBy != "" && order != "" {
		query = query.Order(sortBy + " " + order)
	}
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	// Ejecutar la consulta
	if err := query.Find(&dataFrontend).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(dataFrontend) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data found"})
		return
	}

	if export == "true" {
		data.ExportToCSV(c, dataFrontend)
		return
	}

	// Devolver los datos
	c.JSON(http.StatusOK, gin.H{
		"data":  dataFrontend,
		"total": totalCount,
	})
}

// Importar datos de una CSV a la tabla de personas
func UpdateData(c *gin.Context) {
	dbConn := database.CreateDbConnection()

	records, err := data.ImportCSV(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update db
	for _, record := range records {

		if len(record) < 6 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		age, err := strconv.Atoi(record[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data := models.PeopleData{
			Model: gorm.Model{
				ID: uint(id),
			},
			Age:           age,
			Education:     models.Education(record[2]),
			MaritalStatus: models.MaritalStatus(record[3]),
			Occupation:    models.Occupation(record[4]),
			Income:        models.Income(record[5]),
		}

		if err := dbConn.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"age", "education", "marital_status", "occupation", "income"}),
		}).Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}
