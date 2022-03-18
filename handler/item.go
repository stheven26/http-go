package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ItemHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"message":      "item",
		"product name": title,
		"price":        price,
		"Time":         time.Now(),
	})
}

func IdItemHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"id":           id,
		"message":      "item",
		"product name": title,
		"price":        price,
	})
}
