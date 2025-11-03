package vocechat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 服务器错误
func serverError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  msg,
	})
}

// 操作成功
func sendSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
}
