package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Anda berhasil logout!",
	})
}
