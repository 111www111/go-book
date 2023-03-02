package controller

import (
	"books/config"
	"books/util"
	"github.com/gin-gonic/gin"
)

func HelloController(c *gin.Context) {
	yamlConfig := config.YamlConfig
	c.JSON(200, util.GetTrueDataResult(yamlConfig))
}
