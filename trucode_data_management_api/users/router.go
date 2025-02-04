package users

import (
	"github.com/gin-gonic/gin"
	"trucode.com/final-challenge/shared"
)

func AddUserRoutes(r *gin.Engine) {
	group := r.Group("/users")

	group.Use(shared.AuthenticateSession())

	group.GET("", GetAllUsers)
	group.POST("", CreateUser)
	group.GET("/me", GetMe)
	group.GET("/:id", GetUserById)
	group.DELETE("/:id", DeleteUser)
	group.PUT("/:id", UpdateUser)
}
