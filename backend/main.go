package main

import (
	router "github.com/AVT-Automation/avt-scenario-demo/internal/router" // Importing the functionTest package
)

func main() {
	// Initialise Gin router
	r := router.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
