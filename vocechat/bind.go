package vocechat

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func bind(c *gin.Context, params any, data any) bool {
	//绑定参数
	if params != nil {
		if err := c.ShouldBindQuery(params); err != nil {
			if errs, ok := err.(validator.ValidationErrors); !ok {
				serverError(c, err.Error())
			} else {
				serverError(c, errs.Error())
			}
			return false
		}
	}
	//绑定数据
	if data != nil {
		if err := c.ShouldBindJSON(data); err != nil {
			if err == io.EOF {
				serverError(c, "数据为空。")
			} else if errs, ok := err.(validator.ValidationErrors); !ok {
				serverError(c, err.Error())
			} else {
				serverError(c, errs.Error())
			}
			return false
		}
	}
	return true
}
