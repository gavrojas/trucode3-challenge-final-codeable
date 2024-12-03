package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"trucode.com/final-challenge/auth"
	"trucode.com/final-challenge/data"
	"trucode.com/final-challenge/database"
	"trucode.com/final-challenge/filters"
	"trucode.com/final-challenge/models"
	"trucode.com/final-challenge/peopleData"
	"trucode.com/final-challenge/shared"
	"trucode.com/final-challenge/users"
)

// Configurar router y el uso de cors
func setupRouter(dbConn *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(shared.Cors())
	users.AddUserRoutes(router)
	auth.AddAuthRoutes(router)
	peopleData.AddPeopleDataRoutes(router)
	filters.AddFilterRoutes(router)

	router.GET("/", func(c *gin.Context) {
		tx := dbConn.Exec("SELECT 1")
		fmt.Printf("Error: %v\n", tx.Error)
		c.JSON(http.StatusOK, gin.H{"Success": true})
	})

	return router
}

// Cargar variables de entorno
func loadEnvVars() {
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Unable to load env vars")
		}
	}
}

func main() {
	loadEnvVars()

	dbConn := database.CreateDbConnection()
	dbConn.AutoMigrate(&models.User{}, &models.PeopleData{}, &models.Filters{})

	filePath := "./data/source.data"
	// Revisar si los datos ya fueron cargados de antes
	dataAlreadyLoaded, err := data.IsDataAlreadyLoaded(dbConn, filePath)
	if err != nil {
		log.Fatal("Error checking if data is already loaded:", err)
	}

	if !dataAlreadyLoaded {
		log.Println("Data not loaded")
		dataLoaded, err := data.LoadDataFromFile("./data/source.data")
		if err != nil {
			log.Fatal("Error loading file:", err)
		}

		data.InsertDataToDB(dbConn, dataLoaded)
	} else {
		log.Println("Data already loaded")
	}

	router := setupRouter(dbConn)

	fmt.Printf("Server running on port %s\n", os.Getenv("PORT"))
	if err := router.RunTLS(os.Getenv("PORT"), "/etc/ssl/certs/fullchain.pem", "/etc/ssl/private/privkey.pem"); err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}
}
