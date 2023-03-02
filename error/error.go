package err

import (
	"books/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

//ExceptionMiddleware 全局异常处理
func ExceptionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err := err.(type) {
				case runtime.Error:
					// 运行时异常
					fmt.Println("Recovering from runtime error:", err)
					// 简单返回友好提示，具体可自定义发生错误后处理逻辑
					context.JSON(http.StatusOK, util.GetErrorResult())
					context.Abort()
				default:
					// 其他异常
					if msg, ok := err.(string); ok {
						fmt.Println("Recovering from panic with msg:", msg)
						// 简单返回友好提示，具体可自定义发生错误后处理逻辑
						context.JSON(http.StatusOK, util.GetErrorMsgResult(msg))
						context.Abort()
					} else {
						fmt.Println("Recovering from unknown panic:", err)
						context.JSON(http.StatusOK, util.GetErrorResult())
						context.Abort()
					}
				}
			}
		}()
		context.Next()
	}
}
