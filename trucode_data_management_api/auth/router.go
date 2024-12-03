package auth

import "github.com/gin-gonic/gin"

func AddAuthRoutes(r *gin.Engine) {
	group := r.Group("/auth")
	group.POST("/register", Register)
	group.POST("/login", login)
	group.POST("/logout", logout)
}
