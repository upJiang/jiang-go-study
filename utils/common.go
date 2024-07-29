package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fail(c *gin.Context, httpCode, bizCode int, format string, a ...interface{}) {
	c.AbortWithStatusJSON(httpCode, gin.H{
		"code":    bizCode,
		"message": fmt.Sprintf(format, a...),
		"result":  gin.H{},
	})
}

func FailAs400(c *gin.Context, format string, a ...interface{}) {
	Fail(c, http.StatusBadRequest, http.StatusBadRequest, format, a...)
}

func FailAs500(c *gin.Context, format string, a ...interface{}) {
	Fail(c, http.StatusInternalServerError, http.StatusInternalServerError, format, a...)
}
