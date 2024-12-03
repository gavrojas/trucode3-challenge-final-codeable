package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"trucode.com/final-challenge/database"
	"trucode.com/final-challenge/models"
)

func GetAllUsers(c *gin.Context) {
	dbConn := database.CreateDbConnection()
	var users []models.User
	dbConn.Find(&users)
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	dbConn := database.CreateDbConnection()

	var foundUser models.User
	tx := dbConn.Where("id = ?", c.Param("id")).First(&foundUser)
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, foundUser)
}

func GetMe(c *gin.Context) {
	sessionUser, _ := c.Get("authorizedUser")
	c.JSON(http.StatusOK, sessionUser)
	fmt.Println(sessionUser.(models.User).ID)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	dbConn := database.CreateDbConnection()
	dbConn.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	dbConn := database.CreateDbConnection()
	dbConn.Delete(&models.User{}, id)
	c.JSON(http.StatusNoContent, gin.H{"id": id})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	dbConn := database.CreateDbConnection()
	dbConn.First(&user, id)
	c.BindJSON(&user)
	dbConn.Updates(&user)
	c.JSON(http.StatusOK, user)
}
