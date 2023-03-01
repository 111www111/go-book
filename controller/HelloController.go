package controller

import (
	"books/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloController(c *gin.Context) {
	if true {
		c.JSON(http.StatusOK, util.GetErrorMsgResult("你这个参数是空的"))
	}
}
