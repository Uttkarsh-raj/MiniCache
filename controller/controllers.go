package controller

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Uttkarsh-raj/minicache/database"
	"github.com/Uttkarsh-raj/minicache/handler"
	"github.com/Uttkarsh-raj/minicache/model"
	"github.com/gin-gonic/gin"
)

func SendHelloMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Hello There !!"})
	}
}

func HandleIncomingCommand(newDb *database.Database, commandList *model.CommandsList) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		var newReq model.CommandRequest
		err := c.BindJSON(&newReq)
		if err != nil {
			c.JSON(handler.ResponseReformer(http.StatusBadRequest, false, "Please follow the specified format <command> <key> <value> :"+err.Error(), ""))
			return
		}

		commands := strings.Split(newReq.Command, " ")
		switch commands[0] {
		case "SET":
			{
				data, err := handler.SetRequestHandler(commands, newDb)
				if err != nil {
					c.JSON(handler.ResponseReformer(http.StatusBadRequest, false, err.Error(), ""))
					return
				}
				// here is the change
				handler.StoreData(newDb)
				c.JSON(handler.ResponseReformer(http.StatusOK, true, "", data))
			}
		case "GET":
			{
				data, err := handler.GetRequestHandler(commands, newDb)
				if err != nil {
					c.JSON(handler.ResponseReformer(http.StatusBadRequest, false, err.Error(), ""))
					return
				}
				c.JSON(handler.ResponseReformer(http.StatusOK, true, "", data))
			}
		case "LPUSH":
			{
				data, err := handler.LpushRequestHandler(commands, commandList)
				if err != nil {
					c.JSON(handler.ResponseReformer(http.StatusBadRequest, false, err.Error(), ""))
					return
				}
				handler.StoreDataList(commandList)
				c.JSON(handler.ResponseReformer(http.StatusOK, true, "", data))
			}
		case "RPUSH":
			{
				data, err := handler.RpushRequestHandler(commands, commandList)
				if err != nil {
					c.JSON(handler.ResponseReformer(http.StatusBadRequest, false, err.Error(), ""))
					return
				}
				handler.StoreDataList(commandList)
				c.JSON(handler.ResponseReformer(http.StatusOK, true, "", data))
			}
		case "LMEMBERS":
			{
				data, err := handler.ListMembers(commands, commandList)
				if err != nil {
					c.JSON(handler.ResponseReformer(http.StatusBadRequest, false, err.Error(), ""))
					return
				}
				c.JSON(handler.ResponseReformer(http.StatusOK, true, "", data))
			}
		default:
			{
				c.JSON(handler.ResponseReformer(http.StatusBadRequest, false, "error: please follow the specified format <command> <key> <value> :"+err.Error(), ""))
				return
			}
		}

	}
}
