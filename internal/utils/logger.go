package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/microgolang/logs"
)

//AbortWithLog will abort the gin Context request with error, and print the error message before
func AbortWithLog(c *gin.Context, status int, err error, message interface{}) {
	if err != nil {
		logs.Error(err.Error())
	} else {
		logs.Error(message)
	}
	c.AbortWithStatusJSON(status, message)
}
