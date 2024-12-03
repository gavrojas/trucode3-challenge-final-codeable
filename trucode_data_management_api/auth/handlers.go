package auth

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
	"trucode.com/final-challenge/database"
	"trucode.com/final-challenge/models"
	"trucode.com/final-challenge/shared"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var userInput models.UserInput
	c.BindJSON(&userInput)

	user := models.User{
		Username: userInput.Username,
		Password: userInput.Password,
	}

	dbconn := database.CreateDbConnection()

	if tx := dbconn.Create(&user); tx.Error != nil {
		// Error si el user ya existe por error postgress 23505... el error contiene "duplicate key value violates unique constraint"
		if strings.Contains(tx.Error.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"username": user.Username,
		"id":       user.ID,
	})
}

func login(c *gin.Context) {
	var input models.UserInput
	var user models.User
	c.BindJSON(&input)

	dbconn := database.CreateDbConnection()
	dbconn.Where("username = ?", input.Username).Find(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	sessionToken := uuid.NewV5(uuid.UUID{}, "session").String()

	session := shared.Session{
		Uid: user.ID,
		// 24 horas
		ExpiryTime: time.Now().Add(time.Hour * 24),
	}

	shared.Sessions[sessionToken] = session

	claims := shared.Payload{
		MapClaims: jwt.MapClaims{
			"iat": jwt.NewNumericDate(time.Now()),                       // issued at,
			"eat": jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), // expired at - tiempo que va a durar el token
		},
		Session: sessionToken,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}

func logout(c *gin.Context) {
	tokenStr := shared.GetTokenFromRequest(c)

	if _, ok := shared.Sessions[tokenStr]; !ok {
		delete(shared.Sessions, tokenStr)
		c.JSON(http.StatusOK, gin.H{"error": "Logged out successfully"})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
}
