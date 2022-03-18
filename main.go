package main

import (
	"webapigolang/handler"

	"github.com/gin-gonic/gin"
	// "webapigolang/handler"
	// "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong-pong",
	// 	})
	// })

	//query string
	r.GET("/home", handler.HomeHandler)
	r.GET("/register", handler.RegisterHandler)
	r.GET("/login", handler.LoginHandler)
	r.GET("/logout", handler.LogoutHandler)
	r.GET("/berhasil-checkout", handler.BerhasilCheckoutHandler)

	//query params
	r.GET("/berhasil-login", handler.UserLoginHander)
	r.GET("/product/:id", handler.ProductHandler)
	r.GET("/items", handler.ItemHandler)
	r.GET("/id/:id/item", handler.IdItemHandler)
	r.GET("/checkout", handler.CheckoutHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
