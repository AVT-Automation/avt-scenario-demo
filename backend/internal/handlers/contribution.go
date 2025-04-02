package contribution

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var contribution float64 // This will store the contribution in memory

// ContributionRequest represents the structure of the JSON payload for contribution requests
type ContributionRequest struct {
	Contribution float64 `json:"contribution,omitempty"` // Fixed value (optional)
	Percentage   float64 `json:"percentage,omitempty"`   // Percentage (optional)
	Salary       float64 `json:"salary,omitempty"`       // Salary for percentage-based calculation (optional)
}

// Function to set the contribution value
func SetContribution(c *gin.Context) {
	var req ContributionRequest

	// Bind the incoming JSON payload to the ContributionRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if both a fixed contribution and percentage-based contribution (with salary) are provided
	if req.Contribution != 0 && (req.Percentage != 0 || req.Salary != 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide either fixed contribution or percentage with salary, not both"})
		return
	}

	// Calculate the monthly contribution based on percentage and salary if both are provided
	if req.Percentage != 0 && req.Salary != 0 {
		monthlySalary := req.Salary / 12                      // Convert yearly salary to monthly salary
		contribution = (monthlySalary * req.Percentage) / 100 // Calculate monthly contribution
		c.JSON(http.StatusOK, gin.H{"contribution": contribution})
		return
	}

	// Store the fixed contribution value if provided and return it in the response
	if req.Contribution != 0 {
		contribution = req.Contribution
		c.JSON(http.StatusOK, gin.H{"contribution": contribution})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contribution data"})
}

// Function to retrieve the current contribution value
func GetContribution(c *gin.Context) {
	// Check if contribution is set
	if contribution == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No contribution set"})
		return
	}

	// Return the stored contribution value
	c.JSON(http.StatusOK, gin.H{"contribution": contribution})
}

// Function to retrieve the current contribution value
func ResetContribution(c *gin.Context) {
	// Reset the contribution value to zero
	contribution = 0
}
