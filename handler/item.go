package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ItemHandler(c *gin.Context) {
	judul := c.Query("judul")
	harga := c.Query("harga")

	c.JSON(http.StatusOK, gin.H{
		"message":   "ini item",
		"item name": judul,
		"price":     harga,
		"Time":      time.Now(),
	})
}

func IdItemHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"id":           id,
		"message":      "item name",
		"product name": title,
		"price":        price,
	})
}
