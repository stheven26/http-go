package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BerhasilCheckoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil melakukan checkout pesanan!",
	})
}

func CheckoutHandler(c *gin.Context) {
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"price": price,
		"Time":  time.Now(),
	})
}
