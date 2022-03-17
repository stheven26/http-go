package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "ini product",
		"id":      id,
	})
}
