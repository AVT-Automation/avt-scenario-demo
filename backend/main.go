package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var contribution float64 // This will store the contribution in memory

// Request struct for storing a percentage-based contribution
type ContributionRequest struct {
	Contribution float64 `json:"contribution,omitempty"` // Fixed value (optional)
	Percentage   float64 `json:"percentage,omitempty"`   // Percentage (optional)
	Salary       float64 `json:"salary,omitempty"`       // Salary for percentage-based calculation (optional)
}

func main() {
	// Initialise Gin router
	router := gin.Default()

	// POST endpoint to set the contribution
	router.POST("/contribution", setContribution)

	// GET endpoint to retrieve the contribution
	router.GET("/contribution", getContribution)

	router.Run() // listen and serve on 0.0.0.0:8080
}

// Function to set the contribution value
func setContribution(c *gin.Context) {
	var req ContributionRequest

	// Bind the JSON payload to the request struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if both contribution amount and percentage with salary are provided
	if req.Contribution != 0 && (req.Percentage != 0 || req.Salary != 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide either fixed contribution or percentage with salary, not both"})
		return
	}

	// Case 2: Percentage with Salary (Calculate Monthly Contribution)
	if req.Percentage != 0 && req.Salary != 0 {
		monthlySalary := req.Salary / 12                       // Convert yearly salary to monthly salary
		contribution := (monthlySalary * req.Percentage) / 100 // Calculate monthly contribution
		c.JSON(http.StatusOK, gin.H{"contribution": contribution})
		return
	}

	// If fixed contribution is provided, store that value
	if req.Contribution > 0 {
		contribution = req.Contribution
		c.JSON(http.StatusOK, gin.H{"contribution": contribution})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contribution data"})
}

// Function to retrieve the current contribution value
func getContribution(c *gin.Context) {
	// Check if contribution is set
	if contribution == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No contribution set"})
		return
	}

	// Return the stored contribution value
	c.JSON(http.StatusOK, gin.H{"contribution": contribution})
}
