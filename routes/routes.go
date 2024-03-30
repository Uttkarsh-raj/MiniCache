package routes

import (
	"github.com/Uttkarsh-raj/redis-go/controller"
	"github.com/Uttkarsh-raj/redis-go/database"
	"github.com/Uttkarsh-raj/redis-go/model"
	"github.com/gin-gonic/gin"
)

func IncomingRoutes(incomingServer *gin.Engine, newDb *database.Database, commandList *model.CommandsList) {
	incomingServer.GET("/hello", controller.SendHelloMessage()) // Test
	incomingServer.POST("/", controller.HandleIncomingCommand(newDb, commandList))
}
