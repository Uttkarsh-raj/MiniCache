package main

import (
	"log"

	"github.com/Uttkarsh-raj/minicache/database"
	"github.com/Uttkarsh-raj/minicache/handler"
	"github.com/Uttkarsh-raj/minicache/model"
	"github.com/Uttkarsh-raj/minicache/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	NewDb := database.CreateDatabase()            // Initialize / Create a new Database
	NewList := model.CreateNewCommandsList()      // Initialize a commands queue to store the incoming requests
	handler.Initialize(NewDb, NewList)            // Read data if already present / Persistent storage if want to use as a normal db
	server := gin.New()                           // New server
	server.Use(gin.Logger())                      // Use gin's logger
	routes.IncomingRoutes(server, NewDb, NewList) // add the routes to the server
	log.Fatal(server.Run(":7000"))                // serve on port
}
