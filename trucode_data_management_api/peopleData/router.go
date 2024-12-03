package peopleData

import (
	"github.com/gin-gonic/gin"
	"trucode.com/final-challenge/shared"
)

// NewRouter creates a new router for people data
func AddPeopleDataRoutes(r *gin.Engine) {

	group := r.Group("/people-data")
	group.Use(shared.AuthenticateSession())

	group.GET("", GetPeopleData)
	group.PATCH("/updateData", UpdateData)
}
