package controller

import (
	"books/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloController(c *gin.Context) {
	if true {
		c.JSON(http.StatusOK, util.GetCheckErrorResult())
	}
}
