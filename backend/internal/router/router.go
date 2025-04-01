package router

import (
	contributionHandler "github.com/AVT-Automation/avt-scenario-demo/internal/handlers" // Importing the functionTest package
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// POST endpoint to set the contribution
	router.POST("/contribution", contributionHandler.SetContribution)

	// GET endpoint to retrieve the contribution
	router.GET("/contribution", contributionHandler.GetContribution)

	return router
}
