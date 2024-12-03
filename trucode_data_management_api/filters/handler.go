package filters

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"trucode.com/final-challenge/database"
	"trucode.com/final-challenge/models"
)

func SaveFilters(c *gin.Context) {
	var filters models.Filters

	authorizedUser, _ := c.Get("authorizedUser")

	user := authorizedUser.(models.User)
	filters.UserID = user.ID

	filters.Education = c.DefaultQuery("education", "")
	filters.MaritalStatus = c.DefaultQuery("marital_status", "")
	filters.Occupation = c.DefaultQuery("occupation", "")
	filters.Income = c.DefaultQuery("income", "")

	filters.MaxAge, _ = strconv.Atoi(c.DefaultQuery("max_age", "100"))
	filters.MinAge, _ = strconv.Atoi(c.DefaultQuery("min_age", "0"))

	dbConn := database.CreateDbConnection()

	// Intenta encontrar los filtros existentes
	var existingFilters models.Filters
	result := dbConn.Where("user_id = ?", user.ID).First(&existingFilters)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Si no existen, crea un nuevo registro
			if err := dbConn.Create(&filters).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	} else {
		// Si existen, actualiza el registro
		filters.ID = existingFilters.ID // Asegúrate de que el ID esté presente
		if err := dbConn.Save(&filters).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, filters)
}

func LoadFilters(c *gin.Context) {
	authorizedUser, _ := c.Get("authorizedUser")

	user := authorizedUser.(models.User)
	dbConn := database.CreateDbConnection()
	var filters models.Filters
	result := dbConn.Where("user_id = ?", user.ID).First(&filters)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{"data": nil})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": filters})
}

func GetFilterOptions(c *gin.Context) {
	options := gin.H{
		"workclass":      models.GetAllWorkclassOptions(),
		"education":      models.GetAllEducationOptions(),
		"marital_status": models.GetAllMaritalStatusOptions(),
		"occupation":     models.GetAllOccupationOptions(),
		"relationship":   models.GetAllRelationshipOptions(),
		"race":           models.GetAllRaceOptions(),
		"sex":            models.GetAllSexOptions(),
		"income":         models.GetAllIncomeOptions(),
	}

	c.JSON(http.StatusOK, options)
}
