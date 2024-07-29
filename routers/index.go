package routers

import (
	"github.com/gin-gonic/gin"
	"jiang-go-study/routers/module"
)

func CreateHttpHandler() *gin.Engine {
	root := gin.New()
	root.Use(
		AddCustomizedResponseHeaders(), // 添加 X-Powered-By
		gin.CustomRecovery(CustomizedRecovery),
	)

	base := root.Group("/jiang")

	routersModule.SetupCommonRouter(base)

	return root
}
