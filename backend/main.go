package main

import (
	contributionHandler "github.com/AVT-Automation/avt-scenario-demo/internal/handlers" // Importing the functionTest package
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialise Gin router
	router := gin.Default()

	// POST endpoint to set the contribution
	router.POST("/contribution", contributionHandler.SetContribution)

	// GET endpoint to retrieve the contribution
	router.GET("/contribution", contributionHandler.GetContribution)

	router.GET("/test", contributionHandler.FunctionTest) // Test endpoint

	router.Run() // listen and serve on 0.0.0.0:8080
}
