package api

import (
	"github.com/gin-gonic/gin"
)

type Requirement struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func SetupRequirementsRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/requirements", GetRequirements)
	}
}

func GetRequirements(context *gin.Context) {
	context.JSON(200, gin.H{
		"data": Requirement{ID: 1, Title: "req 1", Description: "desc"},
	})
}
