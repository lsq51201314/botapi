package vocechat

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域设置
func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.Next()
	}
}

// 控制台输出
func console() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Printf("\n* path:%s\n* method:%s\n* status:%d\n********************\n",
			c.Request.URL.Path,
			c.Request.Method,
			c.Writer.Status())
		c.Next()
	}
}

// 处理错误
func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				serverError(c, fmt.Sprintf("%v", err))
				return
			}
		}()
		c.Next()
	}
}
