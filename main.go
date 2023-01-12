package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/orders", getOrderController)
	return router
}

func main() {
	router := setupRouter()
	router.Run(":3000")
}
