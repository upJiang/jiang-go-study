package routers

import (
	consts "jiang-go-study/contants"

	utils "jiang-go-study/utils"

	"github.com/gin-gonic/gin"
)

// 在 HTTP 响应头里添加自定义属性
func AddCustomizedResponseHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Powered-By", consts.SystemName)
		logID := ""
		if v := c.Request.Context().Value(consts.LogGroupIDKey); v != nil {
			logID = v.(string)
		}
		c.Writer.Header().Set(consts.RequestIDKey, logID)
		c.Next()
	}
}

func CustomizedRecovery(c *gin.Context, recovered interface{}) {
	switch t := recovered.(type) {
	case string:
		utils.FailAs500(c, recovered.(string))
	default:
		utils.FailAs500(c, "[%T]%v", t, recovered)
	}
}
