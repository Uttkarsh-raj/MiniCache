package main

import (
	"log"

	"github.com/Uttkarsh-raj/redis-go/database"
	"github.com/Uttkarsh-raj/redis-go/queue"
	"github.com/Uttkarsh-raj/redis-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	NewDb := database.CreateDatabase()             // Initialize / Create a new Database
	NewQueue := queue.CreateNewCommandsQueue()     // Initialize a commands queue to store the incoming requests
	server := gin.New()                            // New server
	server.Use(gin.Logger())                       // Use gin's logger
	routes.IncomingRoutes(server, NewDb, NewQueue) // add the routes to the server
	log.Fatal(server.Run(":7000"))                 // serve on port
}
