package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func success(c *gin.Context, msg string, data interface{}) {
	d := gin.H{
		"code": 0,
		"msg":  msg,
		"data": data,
	}

	c.JSON(http.StatusOK, d)
}

func failed(c *gin.Context, err error) {
	d := gin.H{
		"code": 0,
		"msg":  err.Error(),
	}

	c.JSON(http.StatusOK, d)
}
