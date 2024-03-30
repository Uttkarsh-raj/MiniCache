package routes

import (
	"github.com/Uttkarsh-raj/minicache/controller"
	"github.com/Uttkarsh-raj/minicache/database"
	"github.com/Uttkarsh-raj/minicache/model"
	"github.com/gin-gonic/gin"
)

func IncomingRoutes(incomingServer *gin.Engine, newDb *database.Database, commandList *model.CommandsList) {
	incomingServer.GET("/hello", controller.SendHelloMessage()) // Test
	incomingServer.POST("/", controller.HandleIncomingCommand(newDb, commandList))
}
