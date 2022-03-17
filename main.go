package main

import (
	"webapigolang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong-pong",
	// 	})
	// })
	r.GET("/home", handler.HomeHandler)
	r.GET("/product/:id", handler.ProductHandler)
	r.GET("/item", handler.ItemHandler)
	r.GET("/id/:id/product", handler.IdItemHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
