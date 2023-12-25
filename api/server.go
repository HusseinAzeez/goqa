package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	SetupRequirementsRoutes(router)

	router.Run()

	fmt.Println("GoQA server is up and running...")
}
