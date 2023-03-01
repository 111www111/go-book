package router

import (
	"books/controller"
	err "books/error"
	"books/geteway"
	"github.com/gin-gonic/gin"
)

var Gin *gin.Engine

func init() {
	Gin = gin.Default()
	Router()
}

func Router() {
	//处理跨域
	Gin.Use(geteway.Cors())
	//校验登录
	Gin.Use(geteway.LoginFilter())
	//全局异常捕获
	Gin.Use(err.ExceptionMiddleware())
	//路由
	Gin.GET("/hello", controller.HelloController)
}
