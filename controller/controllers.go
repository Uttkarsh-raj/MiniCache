package controller

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Uttkarsh-raj/redis-go/database"
	"github.com/Uttkarsh-raj/redis-go/handler"
	"github.com/Uttkarsh-raj/redis-go/model"
	"github.com/gin-gonic/gin"
)

func SendHelloMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Hello There !!"})
	}
}

func HandleIncomingCommand(newDb *database.Database) gin.HandlerFunc {
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
		default:
			{
				c.JSON(handler.ResponseReformer(http.StatusBadRequest, false, "error: please follow the specified format <command> <key> <value> :"+err.Error(), ""))
				return
			}
		}

	}
}
