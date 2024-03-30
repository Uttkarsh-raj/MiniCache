package handler

import (
	"errors"
	"fmt"

	"github.com/Uttkarsh-raj/redis-go/database"
	"github.com/Uttkarsh-raj/redis-go/model"
	"github.com/gin-gonic/gin"
)

func SetRequestHandler(request []string, newDb *database.Database) (any, error) {
	if len(request) < 3 {
		return nil, errors.New("error: please follow the specified format <command> <key> <value>")
	}
	for i := 3; i < len(request); i++ {
		request[2] = request[2] + " " + request[i]
	}
	newDb.SetValue(request[1], request[2])
	resp := gin.H{
		"key":   request[1],
		"value": request[2],
	}
	return resp, nil
}

func GetRequestHandler(request []string, newDb *database.Database) (any, error) {
	if len(request) < 2 {
		return nil, errors.New("error: please follow the specified format <command> <key>")
	}
	var result []string
	for i := 1; i < len(request); i++ {
		val, err := newDb.GetValue(request[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	resp := gin.H{
		"key":   request[1:],
		"value": result,
	}
	return resp, nil
}

func LpushRequestHandler(request []string, newList *model.CommandsList) (any, error) {
	if len(request) < 3 {
		return nil, errors.New("error: please follow the specified format <command> <key> <value>")
	}
	name := request[1]
	isPresent := newList.ListExists(name)
	if isPresent {
		newList.LeftPushInList(name, request[2:])
	} else {
		newList.CreateList(name)
		newList.LeftPushInList(name, request[2:])
	}
	resp := gin.H{
		"key":   request[1],
		"value": newList.Store[name],
	}
	return resp, nil
}

func RpushRequestHandler(request []string, newList *model.CommandsList) (any, error) {
	if len(request) < 3 {
		return nil, errors.New("error: please follow the specified format <command> <key> <value>")
	}
	name := request[1]
	isPresent := newList.ListExists(name)
	if isPresent {
		newList.RightPushInList(name, request[2:])
	} else {
		newList.CreateList(name)
		newList.RightPushInList(name, request[2:])
	}
	resp := gin.H{
		"key":   request[1],
		"value": newList.Store[name],
	}
	return resp, nil
}

func ListMembers(request []string, newList *model.CommandsList) (any, error) {
	if len(request) < 2 {
		return nil, errors.New("error: please follow the specified format <command> <key>")
	}
	name := request[1]
	isPresent := newList.ListExists(name)
	if !isPresent && name != "*" {
		return nil, fmt.Errorf("error: Create a list of this key before trying to access it")
	}
	resp := gin.H{
		"key":   request[1],
		"value": newList.Store[name],
	}
	var response map[string]*model.List
	if name == "*" {
		response = newList.Store
		resp = gin.H{
			"key":   request[1],
			"value": response,
		}
	}
	return resp, nil
}
