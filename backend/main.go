package main

import (
	router "github.com/AVT-Automation/avt-scenario-demo/internal/router" // Importing the functionTest package
)

func main() {
	// Initialise Gin router
	r := router.SetupRouter()

	r.Run() // listen and serve on api:8080
}
