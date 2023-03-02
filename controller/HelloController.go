package controller

import (
	"github.com/gin-gonic/gin"
)

func HelloController(c *gin.Context) {
	//panic("这是一个异常")
	num1 := 0
	i := 10 / num1
	c.JSON(200, i)
	//yamlConfig := config.YamlConfig
	//c.JSON(200, util.GetTrueDataResult(yamlConfig))
}
