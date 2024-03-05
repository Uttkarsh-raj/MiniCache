package routes

import (
	"github.com/Uttkarsh-raj/redis-go/controller"
	"github.com/Uttkarsh-raj/redis-go/database"
	"github.com/Uttkarsh-raj/redis-go/list"
	"github.com/gin-gonic/gin"
)

func IncomingRoutes(incomingServer *gin.Engine, newDb *database.Database, commandList *list.CommandsList) {
	incomingServer.GET("/hello", controller.SendHelloMessage()) // Test
	incomingServer.POST("/", controller.HandleIncomingCommand(newDb, commandList))
}
