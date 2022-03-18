package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	user := c.Query("user")
	password := c.Query("password")
	c.JSON(http.StatusOK, gin.H{
		"User":     user,
		"Password": password,
		"Time":     time.Now(),
	})
}

func UserLoginHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Anda berhasil Login!",
	})
}
