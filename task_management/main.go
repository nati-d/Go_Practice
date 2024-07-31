package main

import (
	"task_management/router"
)

func main() {
	// Set up the router using the SetupRouter function from the router package
	r := router.SetupRouter()

	// Run the Gin router on the default port 8080
	r.Run(":8080")
}
