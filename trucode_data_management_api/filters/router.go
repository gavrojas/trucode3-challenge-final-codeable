package filters

import (
	"github.com/gin-gonic/gin"
	"trucode.com/final-challenge/shared"
)

func AddFilterRoutes(r *gin.Engine) {
	group := r.Group("/filters")
	group.Use(shared.AuthenticateSession())

	group.POST("/save", SaveFilters)
	group.GET("/load", LoadFilters)
	group.GET("/options", GetFilterOptions)

}
