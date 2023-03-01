package err

import (
	"books/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 简单返回友好提示，具体可自定义发生错误后处理逻辑
				context.JSON(http.StatusOK, util.GetErrorResult())
				context.Abort()
			}
		}()
		context.Next()
	}
}
