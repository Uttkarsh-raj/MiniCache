package handler

import (
	"github.com/gin-gonic/gin"
)

func ResponseReformer(errorCode int, success bool, errorMessage string, data any) (int, any) {
	return errorCode, gin.H{"success": success, "message": errorMessage, "data": data}
}
