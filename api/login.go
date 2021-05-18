package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "weishi",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "weishi",
	})
}
