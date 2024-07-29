package routersModule

import (
	"github.com/gin-gonic/gin"
)

func SetupCommonRouter(base *gin.RouterGroup) {
	common := base.Group("/common")

	common.GET("/log", func(c *gin.Context) {
		c.String(200, "common log")
	})
}
